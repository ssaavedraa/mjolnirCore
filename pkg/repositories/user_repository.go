package repositories

import "hex/mjolnir-core/pkg/models"

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetByInviteId(inviteId string) (models.User, error)
	GetById(id uint) (models.User, error)
	Update(user models.User) (models.User, error)
}
