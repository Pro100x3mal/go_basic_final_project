package handlers

import (
	"net/http"
	"time"
)

func (th *TaskHandler) handleGetNextDate(w http.ResponseWriter, r *http.Request) {
	now := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	if now == "" || date == "" || repeat == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	nowDate, err := time.Parse("20060102", now)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	nextDate, err := th.reader.NextDate(nowDate, date, repeat)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(nextDate))
}
