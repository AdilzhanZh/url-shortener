package main

import (
	"log"
	"log/slog"
	"os"
	"url-shortener/internal/config"
	"url-shortener/internal/pkg/logger"
	"url-shortener/internal/server"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %s", err.Error())
	}

	slogger := logger.New(cfg.LogLevel)
	slog.SetDefault(slogger)

	slog.Info("starting url-shortener")

	r := gin.New()
	srv := server.New(r, cfg.Port)

	err = srv.Run()
	if err != nil {
		slog.Error("failed to start server", "error", err.Error())
		os.Exit(1)
	}

}
