package main

import (
	"git-lab.boldapps.net/nifty-logix/mvc/app/models"
	"github.com/jinzhu/gorm"
)

// Up is executed when this migration is applied
func Up_20190919115924(txn *gorm.DB) {
	txn.CreateTable(&models.Log{})
}

// Down is executed when this migration is rolled back
func Down_20190919115924(txn *gorm.DB) {
	txn.DropTable(&models.Log{})
}