package models

import (
	"time"
)

type Transaction struct {
    TransactionID uint `gorm:"primaryKey" json:"transaction_id"`
    UserID uint `gorm:"not null" json:"user_id"`
    DateTime time.Time `gorm:"not null" json:"date_time"`
    TotalPrice float64 `gorm:"not null" json:"total_price"`
    PaymentMethod string `gorm:"not null" json:"payment_method"`
    Status string `gorm:"not null" json:"status"`
    CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}