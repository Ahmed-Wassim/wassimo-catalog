package repos

import (
	"context"
	"fmt"

	"github.com/ahmed-wassim/wassimo-catalog/internal/models"
	"gorm.io/gorm"
)

func List(ctx context.Context, db *gorm.DB, status string) ([]models.Restaurant, error) {
	out := []models.Restaurant{}
	q := db.WithContext(ctx).Model(&models.Restaurant{})
	if status != "" {
		q = q.Where("status = ?", status)
	}
	if err := q.Order("id ASC").Find(&out).Error; err != nil {
		return nil, fmt.Errorf("list restaurants: %w", err)
	}
	return out, nil
}
