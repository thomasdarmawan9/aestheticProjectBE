package services

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
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

func (uc *AuthServiceImpl) SignUpUser(user *models.SignUpInput) (*models.SignUpInput, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	user.Email = strings.ToLower(user.Email)
	user.Verified = false

	// Check if password and password confirmation match
	if user.Password != user.PasswordConfirm {
		return nil, errors.New("password and password confirmation do not match")
	}

	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword

	// Check if user with email already exists
	var existingUser models.Tb_Customers
	if err := uc.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("user with that email already exists")
	}

	// Generate customer serial
	user.CustomerSerial = generateCustomerSerial()

	// Create user using raw SQL query
	create := "INSERT INTO tb_customers (created_at, updated_at, customer_serial, name, email, phone_number, address, date_of_birth, password, verified) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result := uc.db.Exec(create, user.CreatedAt, user.UpdatedAt, user.CustomerSerial, user.Name, user.Email, user.PhoneNumber, user.Address, user.DateOfBirth, user.Password, user.Verified)

	if result.Error != nil {
		return nil, errors.New("error creating data")
	}

	// Get Data User after create using raw SQL query
	var newUser models.SignUpInput
	hashedPassConfirm, _ := utils.HashPassword(user.PasswordConfirm)
	newUser.PasswordConfirm = hashedPassConfirm
	query := "SELECT * FROM tb_customers WHERE email = ?"
	if err := uc.db.Raw(query, user.Email).Scan(&newUser).Error; err != nil {
		return nil, errors.New("error querying the database")
	}

	return &newUser, nil
}


// generateCustomerSerial generates a unique customer serial
func generateCustomerSerial() string {
	// You can implement your own logic to generate the customer serial here
	// For example, you can use a combination of current timestamp and a random number

	// Here's a simple example using a timestamp and a random number:
	timestamp := time.Now().Unix()
	randomNumber := rand.Intn(10000) // Change the range as needed
	customerSerial := fmt.Sprintf("%d-%d", timestamp, randomNumber)

	return customerSerial
}

func (uc *AuthServiceImpl) SignInUser(input *models.SignInInput) (*models.DBResponse, error) {
	var user models.Tb_Customers

	// Find user by email
	query := "SELECT * FROM tb_customers WHERE email = ?"
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