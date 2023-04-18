package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"aesthetic/models"
	"aesthetic/utils"

	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	db  *gorm.DB
	ctx context.Context
}

func NewAuthService(db *gorm.DB, ctx context.Context) AuthService {
	return &AuthServiceImpl{db, ctx}
}

func (uc *AuthServiceImpl) SignUpUser(user *models.SignUpInput) (*models.DBResponse, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	user.Email = strings.ToLower(user.Email)
	user.PasswordConfirm = ""
	user.Verified = false
	user.Role = "user"

	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword

	// Check if user with email already exists
	var existingUser models.User
	if err := uc.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("user with that email already exists")
	}

	// Create user using raw SQL query
	create := "INSERT INTO users (created_at, updated_at, name, email, phone_number, address, date_of_birth, preferences, password, verified, role) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result := uc.db.Exec(create, user.CreatedAt, user.UpdatedAt, user.Name, user.Email, user.PhoneNumber, user.Address, user.DateOfBirth, user.Preferences, user.Password, user.Verified, user.Role)

	if result.Error != nil {
		return nil, errors.New("error creating data")
	}

	// Get Data User after create using raw SQL query
	var newUser models.DBResponse
	query := "SELECT * FROM users WHERE email = ?"
	if err := uc.db.Raw(query, user.Email).Scan(&newUser).Error; err != nil {
		return nil, errors.New("error querying the database")
	}

	return &newUser, nil
}

func (uc *AuthServiceImpl) SignInUser(input *models.SignInInput) (*models.DBResponse, error) {
	var user models.User

	// Find user by email
	query := "SELECT * FROM users WHERE email = ?"
	if err := uc.db.Raw(query, user.Email).Scan(&user).Error; err != nil {
		return nil, errors.New("error querying the database")
	}

	// Compare password
	passwordMatch := utils.CheckPassword(input.Password, user.Password)
	if passwordMatch != nil {
		return nil, errors.New("incorrect password")
	}

	// Create JWT token
	token, err := utils.GenerateToken(user.UserID)
	if err != nil {
		return nil, err
	}

	// Create DBResponse
	dbResponse := &models.DBResponse{
		UserID: user.UserID,
		Token: token,
	}

	return dbResponse, nil
}

// func (uc *AuthServiceImpl) SignInUser(*models.SignInInput) (*models.DBResponse, error) {
// 	return nil, nil
// }