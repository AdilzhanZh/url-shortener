package service

import (
	"context"
	"crypto/rand"
	"database/sql"
	"errors"
	"math/big"
	"net/url"
	"time"
	"url-shortener/internal/pkg/error_pkg"

	"url-shortener/internal/model"
	"url-shortener/internal/repository"
)

type URLService struct {
	repo repository.UrlRepo
}

func NewURLService(repo repository.UrlRepo) *URLService {
	return &URLService{
		repo: repo,
	}
}

func (s *URLService) Shorten(ctx context.Context, req model.URL) (*model.URL, error) {
	if !isValidURL(req.OriginalURL) {
		return nil, error_pkg.ErrInvalidURL
	}

	shortCode := req.ShortCode
	if shortCode == "" {
		code, err := generateShortCode(6)
		if err != nil {
			return nil, err
		}

		shortCode = code
	}

	item := &model.URL{
		OriginalURL: req.OriginalURL,
		ShortCode:   shortCode,
	}

	if err := s.repo.Create(ctx, item); err != nil {
		return nil, err
	}

	if err2 := s.repo.SetExpirationDate(ctx, item.ShortCode); err2 != nil {
		return nil, err2
	}

	return item, nil
}

func (s *URLService) Resolve(ctx context.Context, code string) (string, error) {
	item, err := s.repo.GetByShortCode(ctx, code)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", error_pkg.ErrNotFound
		}

		return "", err
	}

	if item.ExpiresAt != nil && item.ExpiresAt.Before(time.Now()) {
		return "", error_pkg.ErrExpired
	}

	if err = s.repo.IncrementClicks(ctx, code); err != nil {
		return "", err
	}

	return item.OriginalURL, nil
}

func isValidURL(rawURL string) bool {
	parsed, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return false
	}

	return parsed.Scheme == "http" || parsed.Scheme == "https"
}

func generateShortCode(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := make([]byte, length)

	for i := range result {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}

		result[i] = charset[n.Int64()]
	}

	return string(result), nil
}
