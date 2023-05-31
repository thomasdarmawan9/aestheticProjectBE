package models

import (
	"time"

	"gorm.io/gorm"
)

// 👈 User struct
type Tb_Customers struct {
	UserID       						uint   		`gorm:"primaryKey" json:"userId"`
	CustomerSerial      		string 		`gorm:"unique;not null" json:"customerSerial,omitempty"`
	Name     								string 		`gorm:"not null" json:"fullName"`
	Email        						string 		`gorm:"unique;not null" json:"email"`
	Password 								string 		`gorm:"not null" json:"password" binding:"required"`
	PhoneNumber 						string 		`gorm:"null" json:"phoneNumber"`
	Address      						string 		`gorm:"null" json:"address"`
	Age      								string 		`gorm:"null" json:"age"`
	DateOfBirth  						string 		`gorm:"null" json:"dateOfBirth"`
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
	UserID       						uint   		`gorm:"primaryKey" json:"userId"`
	CustomerSerial      		string 		`gorm:"unique;not null" json:"customerSerial,omitempty"`
	Name     								string 		`gorm:"not null" json:"fullName"`
	Email        						string 		`gorm:"unique;not null" json:"email"`
	Password 								string 		`gorm:"not null" json:"password" binding:"required"`
	PasswordConfirm    			string    `json:"passwordConfirm" gorm:"-"`
	PhoneNumber 						string 		`gorm:"null" json:"phoneNumber"`
	Address      						string 		`gorm:"null" json:"address"`
	Age      								string 		`gorm:"null" json:"age"`
	DateOfBirth  						string 		`gorm:"null" json:"dateOfBirth"`
	VerificationCode   			string    `json:"verificationCode,omitempty"`
	ResetPasswordToken 			string    `json:"resetPasswordToken,omitempty"`
	ResetPasswordAt    			time.Time `json:"resetPasswordAt,omitempty"`
	Verified           			bool      `json:"verified"`
	CreatedAt          			time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt          			time.Time `json:"updated_at" gorm:"not null"`
}

func (SignUpInput) UsersTable() string {
	return "tb_customers"
}

// 👈 DBResponse struct
type DBResponse struct {
	UserID       						uint   		`gorm:"primaryKey" json:"userId"`
	Token               		string     `json:"token"`
	CustomerSerial      		string 		`gorm:"unique;not null" json:"customerSerial,omitempty"`
	Name     								string 		`gorm:"not null" json:"fullName"`
	Email        						string 		`gorm:"unique;not null" json:"email"`
	Password 								string 		`gorm:"not null" json:"password" binding:"required"`
	PhoneNumber 						string 		`gorm:"null" json:"phoneNumber"`
	Address      						string 		`gorm:"null" json:"address"`
	Age      								string 		`gorm:"null" json:"age"`
	DateOfBirth  						string 		`gorm:"null" json:"dateOfBirth"`
	VerificationCode   			string    `json:"verificationCode,omitempty"`
	ResetPasswordToken 			string    `json:"resetPasswordToken,omitempty"`
	ResetPasswordAt    			time.Time `json:"resetPasswordAt,omitempty"`
	Verified           			bool      `json:"verified"`
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
	}
}

// 👈 PasswordReset struct
type PasswordReset struct {
	gorm.Model
	UserID uint   `json:"userId" gorm:"not null"`
	Token  string `json:"token" gorm:"uniqueIndex;not null"`
}

