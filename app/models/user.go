package models

import (
	"time"
)

// User ... This the User model
type User struct {
	ID        int        `gorm:"primary_key; column:id;" json:"id"`
	Name      string     `gorm:"not null; column:name;" json:"name" valid:"required~Whoops! Name is required."`
	Email     string     `gorm:"unique_index; not null; column:email;" json:"email" valid:"email~Whoops! Email is not valid.,optional"`
	Password  string     `gorm:"not null; column:password;" json:"password"`
	RoleID    int        `gorm:"not null; column:role_id;" json:"role_id" valid:"required~oops! Role is required."`
	CreatedAt time.Time  `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;" json:"deleted_at"`
	Role      Role       `gorm:"foreignkey:role_id" json:"role"`
}
