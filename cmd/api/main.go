package main

import (
	"log"
	"log/slog"
	"url-shortener/internal/config"
	"url-shortener/internal/pkg/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %s", err.Error())
	}

	slogger := logger.New(cfg.LogLevel)
	slog.SetDefault(slogger)

	slog.Info("starting url-shortener")
}
