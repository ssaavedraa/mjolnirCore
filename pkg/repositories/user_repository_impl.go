package repositories

import (
	"fmt"
	"hex/mjolnir-core/pkg/models"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (repo *UserRepositoryImpl) CreateUser(user models.User) (models.User, error) {
	result := repo.db.Create(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func (repo *UserRepositoryImpl) GetUserByEmail(email string) (models.User, error) {
	var user = models.User{}

	result := repo.db.First(&user, "email = ?", email)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (repo *UserRepositoryImpl) GetByInviteId(inviteId string) (models.User, error) {
	var user = models.User{}

	result := repo.db.Preload("Company").Where("invite_id = ?", inviteId).First(&user)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (repo *UserRepositoryImpl) GetById(id uint) (models.User, error) {
	var existingUser models.User

	if err := repo.db.First(&existingUser, id).Error; err != nil {
		return models.User{}, fmt.Errorf("error retrieving user with ID %d, %w", id, err)
	}

	return existingUser, nil
}

func (repo *UserRepositoryImpl) Update(user models.User) (models.User, error) {
	var existingUser models.User

	if err := repo.db.First(&existingUser, user.ID).Error; err != nil {
		return models.User{}, fmt.Errorf("error retriieving user with ID %d, %w", user.ID, err)
	}

	if err := repo.db.Model(&existingUser).Updates(user).Error; err != nil {
		return models.User{}, fmt.Errorf("error updating user with ID %d, %w", user.ID, err)
	}

	return user, nil
}
