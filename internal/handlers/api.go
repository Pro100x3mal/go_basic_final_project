package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (r *router) initRoutes(th *TaskHandler) {
	r.Route("/api", func(r chi.Router) {
		r.Get("/nextdate", th.handleGetNextDate)
		r.Get("/tasks", th.handleGetTasks)
		r.Route("/task", func(r chi.Router) {
			r.Post("/", th.handleCreateTask)
			r.Get("/", th.handleGetTask)
			r.Put("/", th.handleUpdateTask)
		})
	})
	r.Handle("/*", http.FileServer(http.Dir("./web")))
}
