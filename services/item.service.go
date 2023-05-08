package services

import "aesthetic/models"

type ItemService interface {
	GetAllItem(ItemCode string) ([]*models.ItemResponse, error)
	CreateItem(*models.Item) (*models.ItemResponse, error)
}
