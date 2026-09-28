package handlers

import (
	"net/http"
	"strconv"

	"github.com/ahmed-wassim/wassimo-catalog/internal/models"
	"github.com/ahmed-wassim/wassimo-catalog/internal/repos"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ListBranches(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant id"})
		return
	}

	var count int64
	if err := h.DB.WithContext(c.Request.Context()).Model(&models.Restaurant{}).Where("id = ?", id).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
		return
	}

	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "restaurant not found"})
		return
	}

	out, err := repos.ListByRestaurant(c.Request.Context(), h.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
		return
	}
	c.JSON(http.StatusOK, out)
}

type createBranchInput struct {
	Name   string `json:"name" binding:"required"`
	Status string `json:"status"`
}

func (h *Handler) CreateBranch(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant id"})
		return
	}

	var count int64
	if err := h.DB.WithContext(c.Request.Context()).Model(&models.Restaurant{}).Where("id = ?", id).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
		return
	}

	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "restaurant not found"})
		return
	}

	var in createBranchInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	if in.Status != "" && !allowedStatus[in.Status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid status, want active|closed"})
		return
	}

	b := &models.Branch{RestaurantID: id, Name: in.Name, Status: in.Status}
	if err := repos.Create(c.Request.Context(), h.DB, b); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
		return
	}
	c.JSON(http.StatusCreated, b)
}
