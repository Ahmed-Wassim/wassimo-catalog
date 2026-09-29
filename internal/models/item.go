package models

import "time"

type Item struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	CategoryID  int64     `json:"category_id"`
	BranchID    int64     `json:"branch_id"`
	Name        string    `gorm:"not null" json:"name"`
	PriceCents  int       `json:"price_cents"`
	Currency    string    `gorm:"not null" json:"currency"`
	IsAvailable bool      `json:"is_available"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
