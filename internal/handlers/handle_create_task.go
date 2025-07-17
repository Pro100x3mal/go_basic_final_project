package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
)

func (th *TaskHandler) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJson(w, &models.RespError{Error: "invalid request body"}, http.StatusBadRequest)
		return
	}

	id, err := th.writer.CreateTask(&task)
	if err != nil {
		if errors.Is(err, models.ErrInternalServerError) {
			log.Println(err)
			writeJson(w, &models.RespError{Error: "internal server error"}, http.StatusInternalServerError)
			return
		}
		writeJson(w, &models.RespError{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	writeJson(w, &models.RespID{ID: id}, http.StatusCreated)
}
