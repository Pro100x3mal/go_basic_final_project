package config

import (
	"flag"
	"os"
)

type Config struct {
	ServerPort string
}

func NewConfig() *Config {
	var cfg Config
	flag.StringVar(&cfg.ServerPort, "p", "7540", "port of HTTP server")
	flag.Parse()

	if envPort := os.Getenv("TODO_PORT"); envPort != "" {
		cfg.ServerPort = envPort
	}
	return &cfg
}
