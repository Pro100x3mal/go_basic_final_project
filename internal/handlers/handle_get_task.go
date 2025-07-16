package handlers

import (
	"errors"
	"net/http"

	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
)

func (th *TaskHandler) handleGetTask(w http.ResponseWriter, r *http.Request) {
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

	writeJson(w, task)
}
