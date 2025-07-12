package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
)

func writeJson(w http.ResponseWriter, data any) {
	var taskResp models.TaskResp

	switch v := data.(type) {
	case string:
		taskResp.ID = v
		log.Println(taskResp)
	case error:
		taskResp.Error = v.Error()
	case []*models.Task:
		taskResp.Tasks = v
		if taskResp.Tasks == nil || len(taskResp.Tasks) == 0 {
			taskResp.Tasks = []*models.Task{}
		}
	default:
		log.Printf("writeJson: unsupported data type %T\n", data)
		taskResp.Error = "unsupported data type"
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(taskResp); err != nil {
		log.Println(err)
	}
}
