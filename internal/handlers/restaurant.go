package handlers

import (
	"net/http"

	"github.com/ahmed-wassim/wassimo-catalog/internal/repos"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) List(c *gin.Context) {
	out, err := repos.List(c.Request.Context(), h.DB, c.Query("status"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
		return
	}
	c.JSON(http.StatusOK, out)
}
