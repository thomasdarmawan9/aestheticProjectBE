package services

import "aesthetic/models"

type AuthService interface {
	SignUpUser(*models.SignUpInput) (*models.SignUpInput, error)
	SignInUser(*models.SignInInput) (*models.DBResponse, error)
}
