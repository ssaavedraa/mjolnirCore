package repositories

import (
	"hex/cms/pkg/config"
	"hex/cms/pkg/models"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repo *UserRepositoryImpl) CreateUser (user models.User) (models.User, error) {
	result := config.DB.Create(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}