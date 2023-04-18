package models

import (
	"time"
)

type Item struct {
	ItemID uint `gorm:"primaryKey" json:"item_id"`
	Name string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
	ImageURL string `gorm:"not null" json:"image_url"`
	Price float64 `gorm:"not null" json:"price"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}