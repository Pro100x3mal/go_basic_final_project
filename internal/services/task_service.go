package services

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
)

const (
	limit   = 50
	tFormat = "20060102"
)

type TaskRepoWriter interface {
	AddTask(task *models.Task) (int64, error)
	UpdateTask(task *models.Task) error
	UpdateTaskDate(task *models.Task) error
	DeleteTask(id string) error
}

type TaskRepoReader interface {
	GetTask(id string) (*models.Task, error)
	GetTasks(limit int) ([]*models.Task, error)
	GetTasksByDate(date string, limit int) ([]*models.Task, error)
	GetTasksByKeyword(search string, limit int) ([]*models.Task, error)
}

type TaskRepoInterface interface {
	TaskRepoWriter
	TaskRepoReader
}

type TaskService struct {
	writer TaskRepoWriter
	reader TaskRepoReader
}

func NewTaskService(repo TaskRepoInterface) *TaskService {
	return &TaskService{
		writer: repo,
		reader: repo,
	}
}

func (ts *TaskService) CreateTask(task *models.Task) (string, error) {
	if task.Title == "" {
		return "", errors.New("title is required")
	}

	if err := checkDate(task); err != nil {
		return "", err
	}

	id, err := ts.writer.AddTask(task)
	if err != nil {
		return "", models.ErrInternalServerError
	}

	return strconv.Itoa(int(id)), nil
}

func (ts *TaskService) ChangeTask(task *models.Task) error {
	if task.ID == "" || task.Date == "" || task.Title == "" {
		return errors.New("task id or title or date is empty")
	}

	err := checkDate(task)
	if err != nil {
		return err
	}

	err = ts.writer.UpdateTask(task)
	if err != nil {
		if errors.Is(err, models.ErrTaskNotFound) {
			return err
		}
		return models.ErrInternalServerError
	}
	return nil
}

func (ts *TaskService) CompleteTask(id string) error {
	task, err := ts.GetTaskByID(id)
	if err != nil {
		return err
	}

	if task.Repeat == "" {
		err = ts.RemoveTask(id)
		if err != nil {
			return err
		}
		return nil
	}

	task.Date, err = nextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		return err
	}

	err = ts.writer.UpdateTaskDate(task)
	if err != nil {
		if errors.Is(err, models.ErrTaskNotFound) {
			return err
		}
		return models.ErrInternalServerError
	}

	return nil
}

func (ts *TaskService) RemoveTask(id string) error {
	if id == "" {
		return errors.New("task id is required")
	}

	err := ts.writer.DeleteTask(id)
	if err != nil {
		if errors.Is(err, models.ErrTaskNotFound) {
			return err
		}
		return models.ErrInternalServerError
	}
	return nil
}

func (ts *TaskService) GetTaskByID(id string) (*models.Task, error) {
	if id == "" {
		return nil, errors.New("task ID must not be empty")
	}

	task, err := ts.reader.GetTask(id)
	if err != nil {
		if errors.Is(err, models.ErrTaskNotFound) {
			return nil, err
		}
		return nil, models.ErrInternalServerError
	}

	return task, nil
}

func (ts *TaskService) GetTasks(search string) ([]*models.Task, error) {
	var (
		tasks []*models.Task
		err   error
	)

	if search == "" {
		tasks, err = ts.reader.GetTasks(limit)
		if err != nil {
			return nil, models.ErrInternalServerError
		}
		return tasks, nil
	}

	var date time.Time

	if date, err = time.Parse("02.01.2006", search); err == nil {
		tasks, err = ts.reader.GetTasksByDate(date.Format(tFormat), limit)
		if err != nil {
			return nil, models.ErrInternalServerError
		}
		return tasks, nil
	}

	tasks, err = ts.reader.GetTasksByKeyword(search, limit)
	if err != nil {
		return nil, models.ErrInternalServerError
	}

	return tasks, nil
}

func (ts *TaskService) GetNextDate(now, date, repeat string) (string, error) {
	if now == "" || date == "" || repeat == "" {
		return "", errors.New("time or date or repeat is required")
	}

	nowDate, err := time.Parse(tFormat, now)
	if err != nil {
		return "", fmt.Errorf("invalid date format: %s", now)
	}

	repeatDate, err := nextDate(nowDate, date, repeat)
	if err != nil {
		return "", err
	}

	return repeatDate, nil
}

func afterNow(date, now time.Time) bool {
	dy, dm, dd := date.Date()
	ny, nm, nd := now.Date()

	if dy != ny {
		return dy > ny
	}

	if dm != nm {
		return dm > nm
	}

	return dd > nd
}

func nextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", errors.New("repeat rule is required")
	}

	curDate, err := time.Parse(tFormat, dstart)
	if err != nil {
		return "", fmt.Errorf("invalid date format: %s", dstart)
	}

	parts := strings.Split(repeat, " ")

	var next time.Time
	switch {
	case parts[0] == "y" && len(parts) == 1:
		for {
			curDate = curDate.AddDate(1, 0, 0)
			if afterNow(curDate, now) {
				next = curDate
				break
			}
		}

	case parts[0] == "d" && len(parts) == 2:
		days, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", fmt.Errorf("invalid repeat rule format: %s", parts[0])
		}

		if days < 0 || days > 400 {
			return "", errors.New("invalid repeat rule: number of days must be between 1 and 400")
		}

		for {
			curDate = curDate.AddDate(0, 0, days)
			if afterNow(curDate, now) {
				next = curDate
				break
			}
		}

	case parts[0] == "w" && len(parts) == 2:
		weekdays := strings.Split(parts[1], ",")
		if len(weekdays) > 7 {
			return "", errors.New("invalid repeat rule format: too many weekdays")
		}

		for {
			curWeekday := int(curDate.Weekday())
			if curWeekday == 0 {
				curWeekday = 7
			}
			isGreat := false
			minTrue, minFalse := 7, 7
			for _, weekday := range weekdays {
				day, err := strconv.Atoi(weekday)
				if err != nil {
					return "", fmt.Errorf("invalid repeat rule format: %s", parts[0])
				}
				if day > 8 || day < 1 {
					return "", errors.New("invalid repeat rule format: weekday must be between 1 and 7")
				}
				if day > curWeekday {
					isGreat = true
					if day < minTrue {
						minTrue = day
					}
				} else {
					if day < minFalse {
						minFalse = day
					}
				}
			}
			if isGreat {
				curDate = curDate.AddDate(0, 0, minTrue-curWeekday)
			} else {
				curDate = curDate.AddDate(0, 0, 7-curWeekday+minFalse)
			}
			if afterNow(curDate, now) {
				next = curDate
				break
			}
		}

	default:
		return "", fmt.Errorf("invalid repeat rule format: %s", parts[0])
	}

	return next.Format(tFormat), nil
}

func checkDate(task *models.Task) error {
	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format(tFormat)
	}

	t, err := time.Parse(tFormat, task.Date)
	if err != nil {
		return errors.New("invalid date format")
	}

	var next string
	if task.Repeat != "" {
		next, err = nextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}

	}

	if afterNow(now, t) {
		if task.Repeat == "" {
			task.Date = now.Format(tFormat)
		} else {
			task.Date = next
		}
	}

	return nil
}
