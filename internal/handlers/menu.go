package handlers

import (
	"net/http"
	"strconv"

	"github.com/ahmed-wassim/wassimo-catalog/internal/models"
	"github.com/ahmed-wassim/wassimo-catalog/internal/repos"
	"github.com/gin-gonic/gin"
)

func (h *Handler) requireExists(c *gin.Context, model any, id int64, notFound string) bool {
	var count int64
	if err := h.DB.WithContext(c.Request.Context()).
		Model(model).Where("id = ?", id).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
		return false
	}
	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": notFound})
		return false
	}
	return true
}

type createCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func (h *Handler) CreateCategory(c *gin.Context) {
	branchID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || branchID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid branch id"})
		return
	}

	if !h.requireExists(c, &models.Branch{}, branchID, "branch not found") {
		return
	}
	var in createCategoryInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	cat := &models.Category{BranchID: branchID, Name: in.Name}
	if err := repos.CreateCategory(c.Request.Context(), h.DB, cat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
		return
	}

	c.JSON(http.StatusCreated, cat)
}

type createItemInput struct {
	Name        string `json:"name" binding:"required"`
	PriceCents  int    `json:"price_cents" binding:"gte=0"`
	Currency    string `json:"currency"`
	IsAvailable *bool  `json:"is_available"`
}

func (h *Handler) CreateItem(c *gin.Context) {
	catID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || catID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category id"})
		return
	}

	if !h.requireExists(c, &models.Category{}, catID, "category not found") {
		return
	}

	var in createItemInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}

	available := true
	if in.IsAvailable != nil {
		available = *in.IsAvailable
	}

	it := &models.Item{CategoryID: catID, Name: in.Name, PriceCents: in.PriceCents,
		Currency: in.Currency, IsAvailable: available}
	if err := repos.CreateItem(c.Request.Context(), h.DB, it); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
		return
	}
	c.JSON(http.StatusCreated, it)
}

func (h *Handler) Menu(c *gin.Context) {
	branchID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || branchID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid branch id"})
		return
	}

	if !h.requireExists(c, &models.Branch{}, branchID, "branch not found") {
		return
	}

	out, err := repos.MenuByBranch(c.Request.Context(), h.DB, branchID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
		return
	}
	c.JSON(http.StatusOK, out)
}

type availabilityInput struct {
	IsAvailable *bool `json:"is_available" binding:"required"`
}

func (h *Handler) SetAvailability(c *gin.Context) {
	itemID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || itemID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item id"})
		return
	}
	var in availabilityInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "is_available is required"})
		return
	}

	res := h.DB.WithContext(c.Request.Context()).Model(&models.Item{}).
		Where("id = ?", itemID).Updates(map[string]any{"is_available": *in.IsAvailable})
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
		return
	}
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
		return
	}
	var it models.Item
	h.DB.First(&it, itemID)
	c.JSON(http.StatusOK, it)
}
