package models

import (
	"time"
)

// Role ... This is Role model
type Role struct {
	ID                  int64      `json:"id"`
	Title               string     `gorm:"type:varchar(20); not null;  column:title;" json:"title" json:"title"`
	Value               string     `gorm:"type:varchar(20); unique_index; not null;  column:value;" json:"value"`
	CreatedBy           int64      `gorm:"type:BIGINT(20);default:0" json:"created_by"`
	UpdatedBy           int64      `gorm:"type:BIGINT(20);default:0" json:"updated_by"`
	DeletedBy           int64      `gorm:"type:BIGINT(20);default:0" json:"deleted_by"`
	CreatedAT           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	DeletedAt           *time.Time `json:"deleted_at"`
	RolePermissionProxy []RolePermissionProxy
}
