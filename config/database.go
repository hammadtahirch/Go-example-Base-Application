package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DBConfig ... DBConfig model
type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

//DBConnection ... This helps to establish the db connection
func DBConnection() *gorm.DB {
	DB := DBConfig{
		Dialect:  os.Getenv("DB_CONNECTION"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_DATABASE"),
		Charset:  "utf8mb4",
	}

	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		DB.Username,
		DB.Password,
		DB.Name,
		DB.Charset)

	db, err := gorm.Open(DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Database not connected")
	}
	return db
}
