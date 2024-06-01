package services

import (
	"errors"
	"hex/cms/pkg/models"
	repositories "hex/cms/pkg/repositories/user_repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockBcrypt struct {
	mock.Mock
}

func TestCreateUser_Success (t *testing.T) {
	mockRepo := new(repositories.MockUserRepository)
	userService := NewUserService(mockRepo)
	input := UserInput{
		Email: "test@mail.com",
		Password: "password",
		Fullname: "Test Name",
		PhoneNumber: "1234567890",
		Address: "Test Address",
	}

	mockRepo.On(
		"CreateUser",
		mock.AnythingOfType("models.User"),
	).Return(models.User{
		Model: gorm.Model{ID: 1},
		Email: input.Email,
		Password: input.Password,
		Fullname: input.Fullname,
		PhoneNumber: input.PhoneNumber,
		Address: input.Address,
	}, nil)

	mockBcrypt := new(MockBcrypt)
	mockBcrypt.On(
		"GenerateFromPassword",
		[]byte("password"),
		10,
	).Return([]byte("hashedPassword"), nil)

	createdUser, err := userService.CreateUser(input)

	assert.NoError(t, err)
	assert.Equal(t, "test@mail.com", createdUser.Email)
	assert.Equal(t, "password", createdUser.Password)
	assert.Equal(t, "Test Name", createdUser.Fullname)
	assert.Equal(t, "1234567890", createdUser.PhoneNumber)
	assert.Equal(t, "Test Address", createdUser.Address)
}

func TestCreateUser_Error (t *testing.T) {
	mockrepo := new(repositories.MockUserRepository)
	userService := NewUserService(mockrepo)
	input := UserInput{
		Email: "test@mail.com",
		Password: "testPassword123!",
		Fullname: "Test Name",
		PhoneNumber: "1234567890",
		Address: "Test Address",
	}

	mockrepo.On(
		"CreateUser",
		mock.AnythingOfType("models.User"),
	).Return(
		models.User{},
		errors.New("Database error"),
	)

	_, err := userService.CreateUser(input)

	assert.Error(t, err)
}