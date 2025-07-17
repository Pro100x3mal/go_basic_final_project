package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
)

func (th *TaskHandler) handleChangeTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJson(w, &models.RespError{Error: "invalid request body"}, http.StatusBadRequest)
		return
	}

	err := th.writer.ChangeTask(&task)
	if err != nil {
		if errors.Is(err, models.ErrInternalServerError) {
			log.Println(err)
			writeJson(w, &models.RespError{Error: "internal server error"}, http.StatusInternalServerError)
			return
		}
		if errors.Is(err, models.ErrTaskNotFound) {
			writeJson(w, &models.RespError{Error: err.Error()}, http.StatusNotFound)
			return
		}
		writeJson(w, &models.RespError{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	writeJson(w, &models.RespOk{}, http.StatusOK)
}
