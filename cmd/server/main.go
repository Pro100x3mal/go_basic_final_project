package main

import (
	"log"

	"github.com/Pro100x3mal/go_basic_final_project/internal/config"
	"github.com/Pro100x3mal/go_basic_final_project/internal/handlers"
	"github.com/Pro100x3mal/go_basic_final_project/internal/repositories"
	"github.com/Pro100x3mal/go_basic_final_project/internal/services"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg := config.NewConfig()
	repo, err := repositories.NewRepository()
	if err != nil {
		return err
	}
	defer repo.Close()

	taskService := services.NewTaskService(repo)
	return handlers.Serve(cfg, taskService)
}
