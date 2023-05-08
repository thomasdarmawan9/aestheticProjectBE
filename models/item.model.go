package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ItemID uint `gorm:"primaryKey" json:"item_id"`
	ItemCode	string `gorm:"not null" json:"item_code"`
	Name string `gorm:"not null" json:"item_name"`
	Description string `gorm:"not null" json:"description"`
	ImageURL string `gorm:"not null" json:"image_url"`
	Price float64 `gorm:"not null" json:"price"`
	CreatedAt    						time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    						time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type ItemResponse struct {
	ItemID uint `gorm:"primaryKey" json:"item_id"`
	ItemCode	string `gorm:"not null" json:"item_code"`
	Name string `gorm:"not null" json:"item_name"`
	Description string `gorm:"not null" json:"description"`
	ImageURL string `gorm:"not null" json:"image_url"`
	Price float64 `gorm:"not null" json:"price"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}