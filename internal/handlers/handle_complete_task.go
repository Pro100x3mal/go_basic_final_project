package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
)

func (th *TaskHandler) handleCompleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var task *models.Task
	var err error

	if id == "" {
		writeJson(w, errors.New("task id is required"))
		return
	}

	task, err = th.reader.GetTaskByID(id)
	if err != nil {
		writeJson(w, err)
		return
	}

	if task.Repeat == "" {
		err = th.writer.RemoveTask(id)
		if err != nil {
			writeJson(w, err)
		}
		writeJson(w, &models.RespOk{})
		return
	}

	task.Date, err = th.reader.NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		writeJson(w, err)
		return
	}

	err = th.writer.ChangeTaskDate(task)
	if err != nil {
		writeJson(w, err)
		return
	}

	writeJson(w, &models.RespOk{})
}
