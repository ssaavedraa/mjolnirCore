package services

import "hex/mjolnir-core/pkg/models"

type UserInput struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	Fullname    string `json:"fullname" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Address     string `json:"address" binding:"required"`
	CompanyId   uint   `json:"companyId" binding:"required"`
	RoleId      uint   `json:"roleId" binding:"required"`
	TeamID      uint   `json:"teamId" binding:"required"`
}

type OptionalUserInput struct {
	Id          uint   `json:"id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Fullname    string `json:"fullname"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
	CompanyId   uint   `json:"companyId"`
	CompanyRole string `json:"companyRole"`
}

type UserInvite struct {
	Email       string `json:"email" binding:"required,email"`
	CompanyName string `json:"companyName" binding:"required"`
	Fullname    string `json:"fullname" binding:"required"`
	TeamID      uint   `json:"teamId" binding:"required"`
}

type UserCredentials struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserService interface {
	CreateUser(input UserInput, creationMethod string) (*models.User, error)
	Login(credentials UserCredentials) (*models.User, string, error)
	InviteUser(invite UserInvite) (*models.User, error)
	GetByInviteId(inviteId string) (*models.User, error)
	UpdateUser(input OptionalUserInput) (*models.User, error)
}
