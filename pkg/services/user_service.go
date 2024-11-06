package services

import (
	"hex/mjolnir-core/pkg/dtos"
	"hex/mjolnir-core/pkg/models"
)

type UserService interface {
	CreateUser(input dtos.UserInput, creationMethod string) (*models.User, error)
	Login(credentials dtos.UserCredentials) (*models.User, string, error)
	InviteUser(invite dtos.UserInvite) (*models.User, error)
	GetByInviteId(inviteId string) (*models.User, error)
	UpdateUser(input dtos.OptionalUserInput) (*models.User, error)
	GetUserById(userId uint) (*models.User, error)
}
