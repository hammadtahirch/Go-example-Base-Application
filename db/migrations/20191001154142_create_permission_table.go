package main

import (
	"git-lab.boldapps.net/nifty-logix/mvc/app/models"
	"github.com/jinzhu/gorm"
)

// Up is executed when this migration is applied
func Up_20191001154142(txn *gorm.DB) {
	txn.CreateTable(&models.Permission{}).AddIndex("idx_id_name_key", "id", "name", "key")
}

// Down is executed when this migration is rolled back
func Down_20191001154142(txn *gorm.DB) {
	txn.DropTable(&models.Permission{})
}
