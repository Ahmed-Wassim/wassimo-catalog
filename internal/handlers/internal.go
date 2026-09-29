package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/ahmed-wassim/wassimo-catalog/internal/repos"
	"github.com/gin-gonic/gin"
)

const maxLookupIDs = 100

func (h *Handler) ItemsLookup(c *gin.Context) {
	raw := c.Query("ids")
	if raw == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ids query parameter is required"})
		return
	}

	ids := strings.Split(raw, ",")
	if len(ids) > maxLookupIDs {
		c.JSON(http.StatusBadRequest, gin.H{"error": "too many ids, max 100"})
		return
	}

	parsedIds := make([]int64, 0, len(ids))
	for _, id := range ids {
		parsedId, err := strconv.ParseInt(strings.TrimSpace(id), 10, 64)
		if err != nil || parsedId <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ids must be positive integers"})
			return
		}
		parsedIds = append(parsedIds, parsedId)
	}

	out, err := repos.ItemsByIDs(c.Request.Context(), h.DB, parsedIds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
		return
	}

	c.JSON(http.StatusOK, out)
}
