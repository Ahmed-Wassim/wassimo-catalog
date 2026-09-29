package models

import "time"

type Category struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	BranchID  int64     `json:"branch_id"`
	Name      string    `gorm:"not null" json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
