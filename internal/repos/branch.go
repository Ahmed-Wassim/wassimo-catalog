package repos

import (
	"context"
	"fmt"

	"github.com/ahmed-wassim/wassimo-catalog/internal/models"
	"gorm.io/gorm"
)

func ListByRestaurant(ctx context.Context, db *gorm.DB, restaurantID int64) ([]models.Branch, error) {
	out := []models.Branch{}
	err := db.WithContext(ctx).Where("restaurant_id = ?", restaurantID).Order("id ASC").Find(&out).Error
	if err != nil {
		return nil, fmt.Errorf("list branches: %w", err)
	}
	return out, nil
}

func Create(ctx context.Context, db *gorm.DB, b *models.Branch) error {
	if err := db.WithContext(ctx).Create(b).Error; err != nil {
		return fmt.Errorf("create branch: %w", err)
	}
	return nil
}
