package repository

import (
	"context"
	"database/sql"
	"url-shortener/internal/model"

	"github.com/jmoiron/sqlx"
)

type UrlRepo interface {
	Create(ctx context.Context, url *model.URL) error
	GetByShortCode(ctx context.Context, code string) (*model.URL, error)
	IncrementClicks(ctx context.Context, code string) error
}

type PsgURLRepo struct {
	db *sqlx.DB
}

func NewPsgURLRepo(db *sqlx.DB) *PsgURLRepo {
	return &PsgURLRepo{
		db: db,
	}
}

func (r *PsgURLRepo) Create(ctx context.Context, url *model.URL) error {
	query := `
		INSERT INTO urls (original_url, short_code, expires_at)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at, clicks
	`

	return r.db.QueryRowxContext(
		ctx,
		query,
		url.OriginalURL,
		url.ShortCode,
		url.ExpiresAt,
	).Scan(
		&url.ID,
		&url.CreatedAt,
		&url.UpdatedAt,
		&url.Clicks,
	)
}

func (r *PsgURLRepo) GetByShortCode(ctx context.Context, code string) (*model.URL, error) {
	query := `
		SELECT id, original_url, short_code, created_at, updated_at, expires_at, clicks
		FROM urls
		WHERE short_code = $1
	`

	var url model.URL

	if err := r.db.GetContext(ctx, &url, query, code); err != nil {
		return nil, err
	}

	return &url, nil
}

func (r *PsgURLRepo) IncrementClicks(ctx context.Context, code string) error {
	query := `
		UPDATE urls
		SET clicks = clicks + 1,
		    updated_at = NOW()
		WHERE short_code = $1
	`

	result, err := r.db.ExecContext(ctx, query, code)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
