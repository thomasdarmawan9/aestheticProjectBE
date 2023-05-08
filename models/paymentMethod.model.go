package models

import (
	"gorm.io/gorm"
)

type PaymentMethodGroup struct {
	gorm.Model
	Id uint `gorm:"primaryKey" json:"id"`
	NameGroup     string            `json:"nameGroup"`
	PaymentMethodItems []PaymentMethodItem `gorm:"foreignKey:PaymentMethodGroupID"`
}

type PaymentMethodItem struct {
	gorm.Model
	Id uint `gorm:"primaryKey" json:"id"`
	NameItem       string          `json:"nameItem"`
	PaymentMethodGroupID uint   `json:"-"`
}


