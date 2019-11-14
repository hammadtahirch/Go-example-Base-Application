package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// User ... This the User model
type User struct {
	ID          int64      `json:"id"`
	Name        string     `gorm:"type:varchar(40); not null; column:name;" json:"name"`
	Email       string     `gorm:"type:varchar(50);unique_index; not null; column:email;" json:"email"`
	Password    string     `gorm:"type:text;not null; column:password;" json:"password"`
	PhoneNumber string     `gorm:"type:varchar(20);not null; column:phone_number;" json:"phone_number"`
	SecureCode  int        `gorm:"type:int(5); not null; unsigned; column:secure_code;" json:"secure_code"`
	RoleID      int64      `gorm:"column:role_id;" json:"role_id"`
	CreatedBy   int64      `gorm:"type:BIGINT(20);default:0" json:"created_by"`
	UpdatedBy   int64      `gorm:"type:BIGINT(20);default:0" json:"updated_by"`
	DeletedBy   int64      `gorm:"type:BIGINT(20);default:0" json:"deleted_by"`
	CreatedAt   time.Time  `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at;" json:"deleted_at"`
	Role        Role       `gorm:"foreignkey:role_id" json:"role"`
}

// UserCredentials ... User Credentails models.
type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims ..  this will help to make jwt payload
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// TokenPayload ... helps to set the token
type TokenPayload struct {
	Token string `json:"token"`
}
