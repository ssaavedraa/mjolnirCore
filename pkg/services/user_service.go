package services

import "hex/cms/pkg/models"

type UserInput struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Address string `json:"address" binding:"required"`
	CompanyId uint `json:"companyId" binding:"required"`
}

type UserCredentials struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserService interface {
	CreateUser (input UserInput) (models.User, error)
	Login (credentials UserCredentials) (models.User, string, error)
}