package handlers

import (
	"net/http"

	"github.com/Pro100x3mal/go_basic_final_project/internal/config"
	"github.com/Pro100x3mal/go_basic_final_project/internal/middlewares"
	"github.com/go-chi/chi/v5"
)

func (r *router) initRoutes(cfg *config.Config, th *TaskHandler, ah *AuthHandler) {
	authMiddleware := middlewares.NewAuthMiddleware(ah, cfg)

	r.Route("/api", func(r chi.Router) {
		r.Get("/nextdate", th.handleGetNextDate)
		r.With(authMiddleware).Get("/tasks", th.handleGetTasks)
		r.Post("/signin", ah.handleAuth)
		r.Route("/task", func(r chi.Router) {
			r.Use(authMiddleware)
			r.Post("/done", th.handleCompleteTask)
			r.Post("/", th.handleCreateTask)
			r.Get("/", th.handleGetTask)
			r.Put("/", th.handleChangeTask)
			r.Delete("/", th.handleRemoveTask)
		})
	})
	r.Handle("/*", http.FileServer(http.Dir("./web")))
}
