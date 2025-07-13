package handlers

import (
	"net/http"
)

func (r *router) initRoutes(th *TaskHandler) {
	r.Get("/api/nextdate", th.handleGetNextDate)
	r.Get("/api/tasks", th.handleGetTasks)
	r.Post("/api/task", th.handleCreateTask)
	r.Handle("/*", http.FileServer(http.Dir("./web")))
}
