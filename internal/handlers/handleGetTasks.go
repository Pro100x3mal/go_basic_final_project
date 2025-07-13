package handlers

import (
	"net/http"

	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
)

func (th *TaskHandler) handleGetTasks(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	var tasks []*models.Task
	var err error

	if search != "" {
		tasks, err = th.reader.SearchTasks(search)
		if err != nil {
			writeJson(w, err)
			return
		}
	} else {
		tasks, err = th.reader.GetAllTasks()
		if err != nil {
			writeJson(w, err)
			return
		}
	}

	writeJson(w, tasks)
}
