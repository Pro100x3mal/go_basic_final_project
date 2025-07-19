package handlers

import (
	"log"
	"net/http"

	"github.com/Pro100x3mal/go_basic_final_project/internal/config"
	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
	"github.com/go-chi/chi/v5"
)

type TaskServiceWriter interface {
	CreateTask(task *models.Task) (string, error)
	ChangeTask(task *models.Task) error
	RemoveTask(id string) error
	CompleteTask(id string) error
}

type TaskServiceReader interface {
	GetTaskByID(id string) (*models.Task, error)
	GetTasks(search string) ([]*models.Task, error)
	GetNextDate(now, date, repeat string) (string, error)
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
			Addr: cfg.ServerAddr,
		},
	}
}

func Serve(cfg *config.Config, th *TaskHandler, ah *AuthHandler) error {
	r := newRouter()
	r.initRoutes(cfg, th, ah)

	srv := newServer(cfg)
	srv.Handler = r

	log.Printf("Starting server on port %s", cfg.ServerAddr)
	return srv.ListenAndServe()
}
