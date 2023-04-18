package services

import (
	"errors"
	// "log"

	"aesthetic/models"

	"gorm.io/gorm"
)

type UserServiceImpl struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &UserServiceImpl{db}
}

func (us *UserServiceImpl) FindUserById(id string) (*models.DBResponse, error) {
	// var user *models.DBResponse
	user := &models.DBResponse{}
		// Find user by id
		query := "SELECT * FROM users WHERE user_id = ?"
		if err := us.db.Raw(query, id).Scan(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &models.DBResponse{}, err
		}
		return nil, err
	}
	return user, nil
}

func (us *UserServiceImpl) FindUserByEmail(email string) (*models.DBResponse, error) {
	// var user *models.DBResponse
	user := &models.DBResponse{}
		// Find user by email
		query := "SELECT * FROM users WHERE email = ?"
			if err := us.db.Raw(query, email).Scan(user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return &models.DBResponse{}, err
			}
			return nil, err
		}
	return user, nil
}
