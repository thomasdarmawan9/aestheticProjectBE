package models

import (
	"time"
)

type PaymentMethodGroup struct {
	NameGroup							 string									`json:"nameGroup" bson:"nameGroup" binding:"required"`
	CreateAt  						 time.Time 							`json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt 						 time.Time 							`json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type PaymentMethodItem struct {
	NameItem  							string    							`json:"nameItem" bson:"nameItem" binding:"required"`
	CreateAt  						 	time.Time 							`json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt 						 	time.Time 							`json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

