package services

import "hex/mjolnir-core/pkg/models"

type UserInput struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	Fullname    string `json:"fullname" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Address     string `json:"address" binding:"required"`
	CompanyId   uint   `json:"companyId" binding:"required"`
}

type UserInvite struct {
	Email       string `json:"email" binding:"required,email"`
	CompanyName string `json:"companyName" binding:"required"`
	Fullname    string `json:"fullname" binding:"required"`
}

type UserCredentials struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserService interface {
	CreateUser(input UserInput, creationMethod string) (models.User, error)
	Login(credentials UserCredentials) (models.User, string, error)
	InviteUser(invite UserInvite) (models.User, error)
	GetByInviteId(inviteId string) (models.User, error)
}
