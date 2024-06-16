package services

import (
	"hex/cms/pkg/config"
	"hex/cms/pkg/interfaces"
	"hex/cms/pkg/models"
	"hex/cms/pkg/repositories"
)

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
	Bcrypt         interfaces.BcryptInterface
	Jwt            interfaces.JwtInterface
	Config         config.Config
}

func NewUserService(
	userRepository repositories.UserRepository,
	bcrypt interfaces.BcryptInterface,
	config config.Config,
) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Bcrypt:         bcrypt,
		Config:         config,
	}
}

func (us *UserServiceImpl) CreateUser(input UserInput) (models.User, error) {
	hash, err := us.Bcrypt.GenerateFromPassword([]byte(input.Password), 10)

	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Email:       input.Email,
		Password:    string(hash),
		Fullname:    input.Fullname,
		PhoneNumber: input.PhoneNumber,
		Address:     input.Address,
		CompanyID:   input.CompanyId,
	}

	createdUser, err := us.UserRepository.CreateUser(user)

	if err != nil {
		return models.User{}, err
	}

	return createdUser, nil
}

func (us *UserServiceImpl) Login(credentials UserCredentials) (models.User, error) {
	user, err := us.UserRepository.GetUserByEmail(credentials.Email)

	if err != nil {
		return models.User{}, err
	}

	if err := us.Bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(credentials.Password),
	); err != nil {
		return models.User{}, err
	}

	return user, nil
}
