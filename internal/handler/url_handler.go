package handler

import (
	"url-shortener/internal/service"
)

type URLHandler struct {
	service *service.URLService
}

func NewURLHandler(s *service.URLService) *URLHandler {
	return &URLHandler{
		service: s,
	}
}
