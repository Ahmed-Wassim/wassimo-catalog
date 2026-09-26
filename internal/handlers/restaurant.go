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

var allowedStatus = map[string]bool{"active": true, "closed": true}

func (h *Handler) List(c *gin.Context) {
	status := c.Query("status")
	if status != "" && !allowedStatus[status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid status, want active|closed"})
		return
	}
	out, err := repos.List(c.Request.Context(), h.DB, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
		return
	}
	c.JSON(http.StatusOK, out)
}
