package tests

import (
	"errors"
	"hex/cms/pkg/models"
	"hex/cms/pkg/services"
	interfaces_mocks "hex/cms/tests/mocks/interfaces"
	repositories_mocks "hex/cms/tests/mocks/repositories"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestCreateUser_Success (t *testing.T) {
	mockRepo := new(repositories_mocks.MockUserRepository)
	mockBcrypt := new(interfaces_mocks.MockBcryptInterface)
	mockJwt := new(interfaces_mocks.MockJwtInterface)
	userService := services.NewUserService(
		mockRepo,
		mockBcrypt,
		mockJwt,

	)

	input := services.UserInput{
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

func TestCreateUser_DatabaseError (t *testing.T) {
	mockRepo := new(repositories_mocks.MockUserRepository)
	mockBcrypt := new(interfaces_mocks.MockBcryptInterface)
	mockJwt := new(interfaces_mocks.MockJwtInterface)
	userService := services.NewUserService(
		mockRepo,
		mockBcrypt,
		mockJwt,

	)

	input := services.UserInput{
		Email: "test@mail.com",
		Password: "password",
		Fullname: "Test Name",
		PhoneNumber: "1234567890",
		Address: "Test Address",
	}

	mockRepo.On(
		"CreateUser",
		mock.AnythingOfType("models.User"),
	).Return(
		models.User{},
		errors.New("Database error"),
	)

	mockBcrypt.On(
		"GenerateFromPassword",
		[]byte("password"),
		10,
	).Return([]byte("hashedPassword"), nil)

	_, err := userService.CreateUser(input)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "Database error")
}


func TestCreateUser_BcryptError (t *testing.T) {
	mockRepo := new(repositories_mocks.MockUserRepository)
	mockBcrypt := new(interfaces_mocks.MockBcryptInterface)
	mockJwt := new(interfaces_mocks.MockJwtInterface)
	userService := services.NewUserService(
		mockRepo,
		mockBcrypt,
		mockJwt,

	)

	input := services.UserInput{
		Email: "test@mail.com",
		Password: "password",
		Fullname: "Test Name",
		PhoneNumber: "1234567890",
		Address: "Test Address",
	}

	mockRepo.On(
		"CreateUser",
		mock.AnythingOfType("models.User"),
	).Return(
		models.User{},
		errors.New("Database error"),
	)

	mockBcrypt.On(
		"GenerateFromPassword",
		[]byte("password"),
		10,
	).Return(nil , errors.New("Bcrypt error"))

	_, err := userService.CreateUser(input)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "Bcrypt error")
}



func TestLoginUser_Success (t *testing.T) {
	mockRepo := new(repositories_mocks.MockUserRepository)
	mockBcrypt := new(interfaces_mocks.MockBcryptInterface)
	mockJwt := new(interfaces_mocks.MockJwtInterface)
	mockToken := new(interfaces_mocks.MockJwtTokenInterface)
	userService := services.NewUserService(
		mockRepo,
		mockBcrypt,
		mockJwt,

	)

	credentials := services.UserCredentials{
		Email: "test@mail.com",
		Password: "password",
	}

	mockRepo.On(
		"GetUserByEmail",
		credentials.Email,
	).Return(models.User{
		Model: gorm.Model{ID: 1},
		Email: credentials.Email,
		Password: "hashedPassword",
	}, nil)

	mockBcrypt.On(
		"CompareHashAndPassword",
		[]byte("hashedPassword"),
		[]byte(credentials.Password),
	).Return(nil)

	mockToken.On(
		"SignedString",
		mock.Anything,
	).Return("tokenString")

	mockJwt.On(
		"NewWithClaims",
		jwt.SigningMethodHS256,
		mock.AnythingOfType("jwt.MapClaims"),
	).Return(mockToken)

	_, _, err := userService.Login(credentials)

	assert.NoError(t, err)
	// assert.Equal(t, "test@mail.com", user.Email)
	// assert.Equal(t, "tokenString", token)
}