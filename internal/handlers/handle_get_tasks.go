package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
)

func (th *TaskHandler) handleGetTasks(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")

	tasks, err := th.reader.GetTasks(search)
	if err != nil {
		if errors.Is(err, models.ErrInternalServerError) {
			log.Println(err)
			writeJson(w, &models.RespError{Error: "internal server error"}, http.StatusInternalServerError)
			return
		}
		writeJson(w, &models.RespError{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	if tasks == nil || len(tasks) == 0 {
		tasks = []*models.Task{}
	}

	writeJson(w, &models.RespTasks{Tasks: tasks}, http.StatusOK)
	return
}
