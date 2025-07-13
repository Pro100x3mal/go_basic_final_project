package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
)

func (th *TaskHandler) handleUpdateTask(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var task models.Task

	err = json.Unmarshal(data, &task)
	if err != nil {
		writeJson(w, err)
		return
	}

	if task.ID == "" || task.Date == "" || task.Title == "" {
		writeJson(w, errors.New("task id or title or date is empty"))
		return
	}

	err = th.writer.UpdateTask(&task)
	if err != nil {
		writeJson(w, err)
		return
	}

	writeJson(w, &models.RespOk{})
}
