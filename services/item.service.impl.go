package services

import (
	"aesthetic/models"

	"gorm.io/gorm"
)

type ItemServiceImpl struct {
	db *gorm.DB
}

func NewItemService(db *gorm.DB) ItemService {
	return &ItemServiceImpl{db}
}

func (s *ItemServiceImpl) GetAllItem(itemCode string) ([]*models.ItemResponse, error) {
	var items []*models.Item
	result := s.db.Where("item_code LIKE ?", "%"+itemCode+"%").Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}

	itemResponses := make([]*models.ItemResponse, len(items))
	for i, item := range items {
		itemResponses[i] = &models.ItemResponse{
			ItemID:      item.ItemID,
			ItemCode:    item.ItemCode,
			Name:        item.Name,
			Description: item.Description,
			ImageURL:    item.ImageURL,
			Price:       item.Price,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		}
	}

	return itemResponses, nil
}

func (s *ItemServiceImpl) CreateItem(item *models.Item) (*models.ItemResponse, error) {
	result := s.db.Create(item)
	if result.Error != nil {
		return nil, result.Error
	}

	// Return the created item in the response format
	itemResponse := &models.ItemResponse{
		ItemID:      item.ItemID,
		ItemCode:    item.ItemCode,
		Name:        item.Name,
		Description: item.Description,
		ImageURL:    item.ImageURL,
		Price:       item.Price,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}
	return itemResponse, nil
}
