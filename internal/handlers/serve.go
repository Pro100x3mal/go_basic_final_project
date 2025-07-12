package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/Pro100x3mal/go_basic_final_project/internal/config"
	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
	"github.com/go-chi/chi/v5"
)

type TaskService interface {
	CreateTask(task *models.Task) (int64, error)
	NextDate(now time.Time, dstart string, repeat string) (string, error)
	GetAllTasks(limit int) ([]*models.Task, error)
}

type TaskHandler struct {
	taskService TaskService
}

func NewTaskHandler(taskService TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func newRouter(th *TaskHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/api/nextdate", th.nextDayHandler)
	r.Get("/api/tasks", th.getAllTasksHandler)
	r.Post("/api/task", th.addTaskHandler)
	r.Handle("/*", http.FileServer(http.Dir("./web")))
	return r
}

func Serve(cfg *config.Config, taskService TaskService) error {
	h := NewTaskHandler(taskService)
	r := newRouter(h)

	srv := &http.Server{
		Addr:    "localhost:" + cfg.ServerPort,
		Handler: r,
	}

	log.Printf("Starting server on port %s", cfg.ServerPort)

	return srv.ListenAndServe()
}
