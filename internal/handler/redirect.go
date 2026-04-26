package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *URLHandler) Redirect(c *gin.Context) {
	code := c.Param("code")

	url, err := h.service.Resolve(c.Request.Context(), code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, url) // 302
}
