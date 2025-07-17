package handlers

import (
	"io"
	"net/http"

	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
)

func (th *TaskHandler) handleGetNextDate(w http.ResponseWriter, r *http.Request) {
	now := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	nextDate, err := th.reader.GetNextDate(now, date, repeat)
	if err != nil {
		writeJson(w, &models.RespError{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, nextDate)
}
