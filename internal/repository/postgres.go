package repository

import (
	"context"
	"fmt"
	"log/slog"
	"time"
	"url-shortener/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	connectionPath string
}

func NewPostgresDB(cfg *config.Config) (*sqlx.DB, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := sqlx.ConnectContext(ctx, "pgx", cfg.DatabaseURL())
	if err != nil {
		return nil, fmt.Errorf("connect db: %w", err)
	}

	slog.Info("PostgresSQL connected successfully")

	return db, nil
}
