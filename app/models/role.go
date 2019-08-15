package models

import (
	"time"
)

// Role ... This is Role model
type Role struct {
	ID        int        `gorm:"primary_key;" json:"id"`
	Title     string     `gorm:"not null;" json:"title"`
	Value     string     `gorm:"unique_index;not null;" json:"value"`
	CreatedAT time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
