package migration

import (
	"github.com/hammadtahirch/nifty_logix/app/models"
	"github.com/jinzhu/gorm"
)

// Up is executed when this migration is applied
func Up_20191001154200(txn *gorm.DB) {
	txn.CreateTable(&models.RolePermissionProxy{})
}

// Down is executed when this migration is rolled back
func Down_20191001154200(txn *gorm.DB) {
	txn.DropTable(&models.RolePermissionProxy{})
}
