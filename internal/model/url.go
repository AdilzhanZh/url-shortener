package model

import "time"

type URL struct {
	ID          int64  `db:"id" json:"id"`
	OriginalURL string `db:"original_url" json:"original_url"`
	ShortCode   string `db:"short_code" json:"short_cocreated_at"`
	Clicks      int64  `db:"clicks" json:"clicks"`

	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	ExpiresAt *time.Time `db:"expires_at" json:"expires_at,omitempty"`
}
