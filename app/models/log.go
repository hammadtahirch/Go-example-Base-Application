package models

import "time"

// Log .. This helps to add log in db
type Log struct {
	ID            int64     `josn:"id"`
	Type          string    `gorm:"type:varchar(10);column:type; " json:"type"`
	Message       string    `gorm:"type:text;column:message;" json:"message"`
	SyatemMessage string    `gorm:"type:text; column:system_message" josn:"system_message"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

//Error This will help to save error stuff
type Error struct {
	SystemError string
	Code        int
	Message     string
}
