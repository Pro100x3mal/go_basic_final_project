package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
)

func writeJson(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	switch v := data.(type) {
	case struct{}:
		if err := json.NewEncoder(w).Encode(v); err != nil {
			log.Println(err)
		}
	case string:
		if err := json.NewEncoder(w).Encode(models.RespID{ID: v}); err != nil {
			log.Println(err)
		}
	case error:
		if err := json.NewEncoder(w).Encode(models.RespError{Error: v.Error()}); err != nil {
			log.Println(err)
		}
	case []*models.Task:
		if v == nil || len(v) == 0 {
			v = []*models.Task{}
		}
		if err := json.NewEncoder(w).Encode(models.RespTasks{Tasks: v}); err != nil {
			log.Println(err)
		}
	case *models.Task:
		if err := json.NewEncoder(w).Encode(v); err != nil {
			log.Println(err)
		}
	default:
		if err := json.NewEncoder(w).Encode(models.RespError{Error: "unsupported data type"}); err != nil {
			log.Println(err)
		}
	}
}
