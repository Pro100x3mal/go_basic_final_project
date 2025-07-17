package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
)

func (th *TaskHandler) handleGetTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	task, err := th.reader.GetTaskByID(id)
	if err != nil {
		if errors.Is(err, models.ErrInternalServerError) {
			log.Println(err)
			writeJson(w, &models.RespError{Error: "internal server error"}, http.StatusInternalServerError)
			return
		}
		if errors.Is(err, models.ErrTaskNotFound) {
			writeJson(w, &models.RespError{Error: "task not found"}, http.StatusNotFound)
			return
		}
		writeJson(w, &models.RespError{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	writeJson(w, &task, http.StatusOK)
}
