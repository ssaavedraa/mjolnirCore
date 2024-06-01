package services

import (
	"hex/cms/pkg/models"
	repositories "hex/cms/pkg/repositories/user_repository"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
}

type UserInput struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Address string `json:"address" binding:"required"`
}

func NewUserService (userRepository repositories.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (us *UserServiceImpl) CreateUser (input UserInput) (models.User, error) {
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

	createdUser, err := us.UserRepository.CreateUser(user)

	if err != nil {
		return models.User{}, err
	}

	return createdUser, nil
}