package services

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
)

const limit = 50

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

func (ts *TaskService) CreateTask(task *models.Task) (int64, error) {
	err := ts.checkDate(task)
	if err != nil {
		return 0, err
	}

	return ts.writer.AddTask(task)
}

func (ts *TaskService) ChangeTask(task *models.Task) error {
	err := ts.checkDate(task)
	if err != nil {
		return err
	}

	return ts.writer.UpdateTask(task)
}

func (ts *TaskService) ChangeTaskDate(task *models.Task) error {
	return ts.writer.UpdateTaskDate(task)
}

func (ts *TaskService) RemoveTask(id string) error {
	return ts.writer.DeleteTask(id)
}

func (ts *TaskService) GetTaskByID(id string) (*models.Task, error) {
	return ts.reader.GetTask(id)
}

func (ts *TaskService) GetAllTasks() ([]*models.Task, error) {
	return ts.reader.GetTasks(limit)
}

func (ts *TaskService) SearchTasks(search string) ([]*models.Task, error) {
	if date, err := time.Parse("02.01.2006", search); err == nil {
		return ts.reader.GetTasksByDate(date.Format("20060102"), limit)
	}
	return ts.reader.GetTasksByKeyword(search, limit)
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

func (ts *TaskService) NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", errors.New("repeat is required")
	}

	curDate, err := time.Parse("20060102", dstart)
	if err != nil {
		return "", err
	}

	parts := strings.Split(repeat, " ")

	var nextDate time.Time
	switch {
	case parts[0] == "y" && len(parts) == 1:
		for {
			curDate = curDate.AddDate(1, 0, 0)
			if afterNow(curDate, now) {
				nextDate = curDate
				break
			}
		}

	case parts[0] == "d" && len(parts) == 2:
		days, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", err
		}

		if days < 0 || days > 400 {
			return "", errors.New("days too big")
		}

		for {
			curDate = curDate.AddDate(0, 0, days)
			if afterNow(curDate, now) {
				nextDate = curDate
				break
			}
		}

	case parts[0] == "w" && len(parts) == 2:
		weekdays := strings.Split(parts[1], ",")
		if len(weekdays) > 7 {
			return "", errors.New("too many weekdays")
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
					return "", err
				}
				if day > 8 || day < 1 {
					return "", errors.New("weekday too big")
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
				nextDate = curDate
				break
			}
		}

	default:
		return "", errors.New("invalid repeat format")
	}

	return nextDate.Format("20060102"), nil
}

func (ts *TaskService) checkDate(task *models.Task) error {
	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format("20060102")
	}

	t, err := time.Parse("20060102", task.Date)
	if err != nil {
		return err
	}

	var next string
	if task.Repeat != "" {
		next, err = ts.NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}

	}

	if afterNow(now, t) {
		if task.Repeat == "" {
			task.Date = now.Format("20060102")
		} else {
			task.Date = next
		}
	}

	return nil
}
