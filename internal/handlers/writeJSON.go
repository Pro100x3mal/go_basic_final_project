package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
)

func writeJson(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, models.ErrInternalServerError.Error(), http.StatusInternalServerError)
	}
}
