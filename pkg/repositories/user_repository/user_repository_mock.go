package repositories

import "hex/cms/pkg/models"

type MockUserRepository struct {
	CreateUserFunc func(user models.User) (models.User, error)
}

func (m *MockUserRepository) CreateUser (user models.User) (models.User, error) {
	if (m.CreateUserFunc != nil) {
		return m.CreateUserFunc(user)
	}

	return models.User{}, nil
}