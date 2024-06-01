package services

import (
	"golang.org/x/crypto/bcrypt"
	"santiagosaavedra.com.co/invoices/pkg/models"
	"santiagosaavedra.com.co/invoices/pkg/repositories"
)

type UserInput struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Address string `json:"address" binding:"required"`
}

func CreateUser (input UserInput) (models.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)

	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Email: input.Email,
		Password: string(hash),
		Fullname: input.Fullname,
		PhoneNumber: input.PhoneNumber,
		Address: input.Address,
	}

	createdUser, err := repositories.CreateUser(user)

	if err != nil {
		return models.User{}, err
	}

	return createdUser, nil
}