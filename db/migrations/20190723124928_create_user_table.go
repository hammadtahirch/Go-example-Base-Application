package migration

import (
	"github.com/hammadtahirch/nifty_logix/app/models"
	"github.com/jinzhu/gorm"
)

// Up_20190723124928 ... is executed when this migration is applied
func Up_20190723124928(txn *gorm.DB) {
	txn.CreateTable(&models.User{}).AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT").AddIndex("idx_id_name_email", "id", "name", "email")
}

// Down_20190723124928 ... is executed when this migration is rolled back
func Down_20190723124928(txn *gorm.DB) {
	txn.DropTable(&models.User{})
}
