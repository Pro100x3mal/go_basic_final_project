package handlers

import (
	"errors"
	"net/http"

	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
)

func (th *TaskHandler) handleRemoveTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		writeJson(w, errors.New("task id is required"))
		return
	}

	err := th.writer.RemoveTask(id)
	if err != nil {
		writeJson(w, err)
		return
	}

	writeJson(w, &models.RespOk{})
}
