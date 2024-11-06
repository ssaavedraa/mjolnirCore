package dtos

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
	Id          *uint   `json:"id,omitempty"`
	Email       *string `json:"email,omitempty"`
	Password    *string `json:"password,omitempty"`
	Fullname    *string `json:"fullname,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	Address     *string `json:"address,omitempty"`
	CompanyId   *uint   `json:"companyId,omitempty"`
	CompanyRole *string `json:"companyRole,omitempty"`
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

type UserResponse struct {
	ID          *uint   `json:"id,omitempty"`
	Fullname    *string `json:"fullname,omitempty"`
	Email       *string `json:"email,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	Address     *string `json:"address,omitempty"`

	Role     *RoleResponse    `json:"role,omitempty"`
	Team     *TeamResponse    `json:"team,omitempty"`
	Company  *CompanyResponse `json:"company,omitempty"`
	Products []models.Product `json:"products,omitempty"`
}
