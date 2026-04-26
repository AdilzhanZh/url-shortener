package handler

import (
	"net/http"
	"url-shortener/internal/model"

	"github.com/gin-gonic/gin"
)

func (h *URLHandler) Shorten(c *gin.Context) {
	var req model.CreateURL

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.Shorten(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	baseURL := c.Request.Host
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	shortURL := scheme + "://" + baseURL + "/" + result.ShortCode
	c.JSON(http.StatusOK, gin.H{
		"short_url": shortURL,
	})
}
