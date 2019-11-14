package models

import (
	"time"
)

// Permission ... This the User model
type Permission struct {
	ID                  int64      `json:"id"`
	Name                string     `gorm:"type:varchar(40); not null; column:name;" json:"name"`
	Key                 string     `gorm:"type:varchar(50);unique_index; not null; column:key;" json:"key"`
	CreatedBy           int64      `gorm:"type:BIGINT(20);default:0" json:"created_by"`
	UpdatedBy           int64      `gorm:"type:BIGINT(20);default:0" json:"updated_by"`
	DeletedBy           int64      `gorm:"type:BIGINT(20);default:0" json:"deleted_by"`
	CreatedAt           time.Time  `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt           time.Time  `gorm:"column:updated_at;" json:"updated_at"`
	DeletedAt           *time.Time `gorm:"column:deleted_at;" json:"deleted_at"`
	RolePermissionProxy []RolePermissionProxy
}
