package models

// RolePermissionProxy ... This the User model
type RolePermissionProxy struct {
	ID           int64        `json:"id"`
	RoleID       string       `gorm:"type:BIGINT(20); default:0; column:role_id;" json:"role_id"`
	PermissionID string       `gorm:"type:BIGINT(20); default:0; column:permission_id;" json:"permission_id"`
	Role         []Role       `gorm:"foreignkey:role_id" json:"role"`
	Permission   []Permission `gorm:"foreignkey:permission_id" json:"permission"`
}
