package error_pkg

import "errors"

var (
	ErrInvalidURL = errors.New("invalid url")
	ErrNotFound   = errors.New("url not found")
	ErrExpired    = errors.New("url expired")
)

var (
	ErrDuplicateShortCode = errors.New("short code already exists")
)
