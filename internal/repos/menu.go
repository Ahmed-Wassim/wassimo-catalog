package repos

import (
	"context"
	"fmt"

	"github.com/ahmed-wassim/wassimo-catalog/internal/models"
	"gorm.io/gorm"
)

func MenuByBranch(ctx context.Context, db *gorm.DB, branchID int64) ([]models.Item, error) {
	out := []models.Item{}
	err := db.WithContext(ctx).
		Joins("JOIN categories ON categories.id = items.category_id").
		Where("categories.branch_id = ?", branchID).
		Order("categories.id, items.id").
		Find(&out).Error
	if err != nil {
		return nil, fmt.Errorf("menu by branch: %w", err)
	}
	return out, nil
}

func CreateCategory(ctx context.Context, db *gorm.DB, c *models.Category) error {
	if err := db.WithContext(ctx).Create(c).Error; err != nil {
		return fmt.Errorf("create category: %w", err)
	}
	return nil
}

func CreateItem(ctx context.Context, db *gorm.DB, i *models.Item) error {
	if err := db.WithContext(ctx).Create(i).Error; err != nil {
		return fmt.Errorf("create item: %w", err)
	}
	return nil
}

type ItemLookup struct {
	ID         int64  `json:"id"`
	CategoryID int64  `json:"category_id"`
	BranchID   int64  `json:"branch_id"`
	Name       string `json:"name"`
	PriceCents int    `json:"price_cents"`
	Currency   string `json:"currency"`
	Available  bool   `json:"is_available"`
}

func ItemsByIDs(ctx context.Context, db *gorm.DB, ids []int64) ([]ItemLookup, error) {
	out := []ItemLookup{}
	if len(ids) == 0 {
		return out, nil
	}
	err := db.WithContext(ctx).
		Table("items").
		Select("items.id, items.category_id, categories.branch_id, items.name, items.price_cents, items.currency, items.is_available").
		Joins("JOIN categories ON categories.id = items.category_id").
		Where("items.id IN ?", ids).
		Order("items.id").
		Find(&out).Error
	if err != nil {
		return nil, fmt.Errorf("items by ids: %w", err)
	}
	return out, nil
}
