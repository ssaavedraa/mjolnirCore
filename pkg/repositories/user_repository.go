package repositories

import "hex/mjolnir-core/pkg/models"

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByInviteId(inviteId string) (*models.User, error)
	GetUserById(userId uint) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
}
