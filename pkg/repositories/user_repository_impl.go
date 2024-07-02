package repositories

import (
	"hex/mjolnir-core/pkg/config"
	"hex/mjolnir-core/pkg/models"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repo *UserRepositoryImpl) CreateUser(user models.User) (models.User, error) {
	result := config.DB.Create(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func (repo *UserRepositoryImpl) GetUserByEmail(email string) (models.User, error) {
	var user = models.User{}

	result := config.DB.First(&user, "email = ?", email)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (repo *UserRepositoryImpl) GetByInviteId(inviteId string) (models.User, error) {
	var user = models.User{}

	result := config.DB.Preload("Company").Where("invite_id = ?", inviteId).First(&user)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
