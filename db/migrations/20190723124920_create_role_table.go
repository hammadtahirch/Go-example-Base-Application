package migration

import (
	"github.com/hammadtahirch/golang_basic_app/app/models"
	"github.com/jinzhu/gorm"
)

// Up_20190723124920 ... is executed when this migration is applied
func Up_20190723124920(txn *gorm.DB) {
	txn.CreateTable(&models.Role{}).AddIndex("idx_id_title_value", "id", "title", "value")
}

// Down_20190723124920 ... is executed when this migration is rolled back
func Down_20190723124920(txn *gorm.DB) {
	txn.DropTable(&models.Role{})
}
