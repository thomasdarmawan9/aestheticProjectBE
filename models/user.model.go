package models

import (
	"time"

	"gorm.io/gorm"
)

// 👈 User struct
type User struct {
	UserID       						uint   `gorm:"primaryKey" json:"user_id"`
	Role      							string    `json:"role,omitempty"`
	Name     								string `gorm:"not null" json:"full_name"`
	Email        						string `gorm:"unique;not null" json:"email"`
	Password 								string `gorm:"not null" json:"password" binding:"required"`
	PhoneNumber 						string `gorm:"not null" json:"phone_number"`
	Preferences  						string `json:"preferences"`
	Address      						string `gorm:"not null" json:"address"`
	DateOfBirth  						string `gorm:"not null" json:"date_of_birth"`
	VerificationCode   			string    `json:"verificationCode,omitempty"`
	ResetPasswordToken 			string    `json:"resetPasswordToken,omitempty"`
	ResetPasswordAt    			time.Time `json:"resetPasswordAt,omitempty"`
	Verified           			bool      `json:"verified"`
	CreatedAt    						time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    						time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// 👈 SignInInput struct
type SignInInput struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex;not null" json:"email" binding:"required"`
	Password string `gorm:"not null" json:"password" binding:"required"`
}

// 👈 SignUpInput struct
type SignUpInput struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement"`
	Name               string    `json:"name" gorm:"not null"`
	Email              string    `json:"email" gorm:"not null;unique"`
	Password           string    `json:"password" gorm:"not null"`
	PasswordConfirm    string    `json:"passwordConfirm" gorm:"-"`
	Role               string    `json:"role" gorm:"default:user"`
	PhoneNumber 			 string `gorm:"not null" json:"phone_number"`
	Address      			 string `gorm:"not null" json:"address"`
	DateOfBirth  			 string `gorm:"not null" json:"date_of_birth"`
	Preferences  			 string ` json:"preferences"`
	VerificationCode   string    `json:"verificationCode,omitempty" gorm:"-"`
	ResetPasswordToken string    `json:"resetPasswordToken,omitempty" gorm:"-"`
	ResetPasswordAt    time.Time `json:"resetPasswordAt,omitempty" gorm:"-"`
	Verified           bool      `json:"verified" gorm:"default:false"`
	CreatedAt          time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"not null"`
}

func (SignUpInput) UsersTable() string {
	return "users"
}

// 👈 DBResponse struct
type DBResponse struct {
	UserID       			 uint   						`gorm:"primaryKey" json:"user_id"`
	Token               string             `json:"token"`
	Name               string             `json:"name"`
	Email              string             `json:"email"`
	Password           string             `json:"password"`
	PasswordConfirm    string             `json:"passwordConfirm,omitempty"`
	Role               string             `json:"role"`
	VerificationCode   string             `json:"verificationCode,omitempty"`
	ResetPasswordToken string             `json:"resetPasswordToken,omitempty"`
	ResetPasswordAt    time.Time          `json:"resetPasswordAt,omitempty"`
	Verified           bool               `json:"verified"`
	CreatedAt          time.Time          `json:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at"`
}

// 👈 ForgotPasswordToken struct
type ForgotPasswordToken struct {
	gorm.Model
	Token  string `json:"token" gorm:"uniqueIndex;not null"`
	UserID uint   `json:"userId" gorm:"not null"`
}

// 👈 UserResponse struct
type UserResponse struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func FilteredResponse(user *DBResponse) UserResponse {
	return UserResponse{
		ID:        user.UserID,
		Email:     user.Email,
		Name:      user.Name,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// 👈 PasswordReset struct
type PasswordReset struct {
	gorm.Model
	UserID uint   `json:"userId" gorm:"not null"`
	Token  string `json:"token" gorm:"uniqueIndex;not null"`
}

