package main

import (
	"log"
	"log/slog"
	"os"

	"url-shortener/internal/config"
	"url-shortener/internal/handler"
	"url-shortener/internal/pkg/logger"
	"url-shortener/internal/repository"
	"url-shortener/internal/server"
	"url-shortener/internal/service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %s", err.Error())
	}

	slogger := logger.New(cfg.LogLevel)
	slog.SetDefault(slogger)

	db, err2 := repository.NewPostgresDB(cfg)
	if err2 != nil {
		slog.Error("failed to connect postgres", "error", err2.Error())
		os.Exit(1)
	}
	defer func() {
		if err := db.Close(); err != nil {
			slog.Error("failed to close db", "error", err)
		}
	}()

	urlRepo := repository.NewPsgURLRepo(db)
	urlService := service.NewURLService(urlRepo)
	urlHandler := handler.NewURLHandler(urlService)

	r, err3 := urlHandler.InitRouts()
	if err3 != nil {
		slog.Error("failed on InitRouts() function")
	}

	srv := server.New(r, cfg.Port)

	slog.Info("starting url-shortener", "port", cfg.Port)

	if err := srv.Run(); err != nil {
		slog.Error("failed to start server", "error", err.Error())
		os.Exit(1)
	}
}
