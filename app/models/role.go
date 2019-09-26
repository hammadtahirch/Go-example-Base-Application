package models

import (
	"time"
)

// Role ... This is Role model
type Role struct {
	ID        int64      `json:"id"`
	Title     string     `gorm:"type:varchar(20); not null;  column:title;" json:"title" json:"title"`
	Value     string     `gorm:"type:varchar(20); unique_index; not null;  column:value;" json:"value"`
	CreatedAT time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
