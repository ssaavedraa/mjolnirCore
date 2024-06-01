package services

import "hex/cms/pkg/models"

type UserService interface {
	CreateUser (input UserInput) (models.User, error)
}