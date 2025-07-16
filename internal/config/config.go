package config

import (
	"flag"
	"os"
)

type Config struct {
	ServerPort string
	Password   string
	DBFile     string
	JWTSecret  string
}

func NewConfig() *Config {
	var cfg Config
	flag.StringVar(&cfg.ServerPort, "p", "7540", "port of HTTP server")
	flag.Parse()

	if envPort := os.Getenv("TODO_PORT"); envPort != "" {
		cfg.ServerPort = envPort
	}

	envPassword := os.Getenv("TODO_PASSWORD")
	if envPassword == "" {
		envPassword = "admin"
	}
	cfg.Password = envPassword

	jwtSecret := os.Getenv("TODO_JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "very-secret-key"
	}
	cfg.JWTSecret = jwtSecret

	envDBFile := os.Getenv("TODO_DBFILE")
	if envDBFile == "" {
		envDBFile = "scheduler.db"
	}
	cfg.DBFile = envDBFile

	return &cfg
}
