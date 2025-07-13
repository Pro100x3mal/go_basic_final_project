package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/Pro100x3mal/go_basic_final_project/internal/config"
	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
	"github.com/go-chi/chi/v5"
)

type TaskServiceWriter interface {
	CreateTask(task *models.Task) (int64, error)
	UpdateTask(task *models.Task) error
}

type TaskServiceReader interface {
	NextDate(now time.Time, dstart string, repeat string) (string, error)
	GetTaskByID(id string) (*models.Task, error)
	GetAllTasks() ([]*models.Task, error)
	SearchTasks(search string) ([]*models.Task, error)
}

type TaskServiceInterface interface {
	TaskServiceWriter
	TaskServiceReader
}

type TaskHandler struct {
	writer TaskServiceWriter
	reader TaskServiceReader
}

func NewTaskHandler(service TaskServiceInterface) *TaskHandler {
	return &TaskHandler{
		writer: service,
		reader: service,
	}
}

type router struct {
	*chi.Mux
}

func newRouter() *router {
	return &router{
		chi.NewRouter(),
	}
}

type server struct {
	*http.Server
}

func newServer(cfg *config.Config) *server {
	return &server{
		&http.Server{
			Addr: "localhost:" + cfg.ServerPort,
		},
	}
}

func Serve(cfg *config.Config, th *TaskHandler) error {
	r := newRouter()
	r.initRoutes(th)

	srv := newServer(cfg)
	srv.Handler = r

	log.Printf("Starting server on port %s", cfg.ServerPort)
	return srv.ListenAndServe()
}
