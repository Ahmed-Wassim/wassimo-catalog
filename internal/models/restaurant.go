package models

import "time"

type Restaurant struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Status    string    `gorm:"not null;default:active" json:"status"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
