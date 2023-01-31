package models

import "github.com/jinzhu/gorm"

func InitTableCustomer(db *gorm.DB) {
	db.Debug().AutoMigrate(&Customer{})
}
