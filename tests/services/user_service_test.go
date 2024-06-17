package tests

import (
	"errors"
	"hex/mjolnir-core/pkg/models"
	"hex/mjolnir-core/pkg/services"
	config_mocks "hex/mjolnir-core/tests/mocks/config"
	interfaces_mocks "hex/mjolnir-core/tests/mocks/interfaces"
	repositories_mocks "hex/mjolnir-core/tests/mocks/repositories"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type TestSetup struct {
	mockRepo    *repositories_mocks.MockUserRepository
	mockBcrypt  *interfaces_mocks.MockBcryptInterface
	mockJwt     *interfaces_mocks.MockJwtInterface
	mockConfig  *config_mocks.MockConfig
	userService services.UserService
}

func setup(_ *testing.T) *TestSetup {
	mockRepo := new(repositories_mocks.MockUserRepository)
	mockBcrypt := new(interfaces_mocks.MockBcryptInterface)
	mockJwt := new(interfaces_mocks.MockJwtInterface)
	mockConfig := new(config_mocks.MockConfig)
	userService := services.NewUserService(
		mockRepo,
		mockBcrypt,
		mockJwt,
		mockConfig,
	)

	return &TestSetup{
		mockRepo:    mockRepo,
		mockBcrypt:  mockBcrypt,
		mockJwt:     mockJwt,
		mockConfig:  mockConfig,
		userService: userService,
	}
}

func TestCreateUser_Success(t *testing.T) {
	ts := setup(t)

	input := services.UserInput{
		Email:       "test@mail.com",
		Password:    "password",
		Fullname:    "Test Name",
		PhoneNumber: "1234567890",
		Address:     "Test Address",
	}

	ts.mockRepo.On(
		"CreateUser",
		mock.AnythingOfType("models.User"),
	).Return(models.User{
		Model:       gorm.Model{ID: 1},
		Email:       input.Email,
		Password:    input.Password,
		Fullname:    input.Fullname,
		PhoneNumber: input.PhoneNumber,
		Address:     input.Address,
	}, nil)

	ts.mockBcrypt.On(
		"GenerateFromPassword",
		[]byte("password"),
		10,
	).Return([]byte("hashedPassword"), nil)

	createdUser, err := ts.userService.CreateUser(input)

	assert.NoError(t, err)
	assert.Equal(t, "test@mail.com", createdUser.Email)
	assert.Equal(t, "password", createdUser.Password)
	assert.Equal(t, "Test Name", createdUser.Fullname)
	assert.Equal(t, "1234567890", createdUser.PhoneNumber)
	assert.Equal(t, "Test Address", createdUser.Address)
}

func TestCreateUser_DatabaseError(t *testing.T) {
	ts := setup(t)

	input := services.UserInput{
		Email:       "test@mail.com",
		Password:    "password",
		Fullname:    "Test Name",
		PhoneNumber: "1234567890",
		Address:     "Test Address",
	}

	ts.mockRepo.On(
		"CreateUser",
		mock.AnythingOfType("models.User"),
	).Return(
		models.User{},
		errors.New("Database error"),
	)

	ts.mockBcrypt.On(
		"GenerateFromPassword",
		[]byte("password"),
		10,
	).Return([]byte("hashedPassword"), nil)

	_, err := ts.userService.CreateUser(input)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "Database error")
}

func TestCreateUser_BcryptError(t *testing.T) {
	ts := setup(t)

	input := services.UserInput{
		Email:       "test@mail.com",
		Password:    "password",
		Fullname:    "Test Name",
		PhoneNumber: "1234567890",
		Address:     "Test Address",
	}

	ts.mockRepo.On(
		"CreateUser",
		mock.AnythingOfType("models.User"),
	).Return(
		models.User{},
		errors.New("Database error"),
	)

	ts.mockBcrypt.On(
		"GenerateFromPassword",
		[]byte("password"),
		10,
	).Return(nil, errors.New("Bcrypt error"))

	_, err := ts.userService.CreateUser(input)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "Bcrypt error")
}

func TestLoginUser_Success(t *testing.T) {
	ts := setup(t)
	mockToken := new(interfaces_mocks.MockJwtTokenInterface)

	credentials := services.UserCredentials{
		Email:    "test@mail.com",
		Password: "password",
	}

	ts.mockRepo.On(
		"GetUserByEmail",
		credentials.Email,
	).Return(models.User{
		Model:    gorm.Model{ID: 1},
		Email:    credentials.Email,
		Password: "hashedPassword",
	}, nil)

	ts.mockBcrypt.On(
		"CompareHashAndPassword",
		[]byte("hashedPassword"),
		[]byte(credentials.Password),
	).Return(nil)

	ts.mockJwt.On(
		"NewWithClaims",
		jwt.SigningMethodHS256,
		mock.AnythingOfType("jwt.MapClaims"),
	).Return(mockToken)

	ts.mockConfig.On(
		"GetEnv",
		"JWT_SECRET",
	).Return("jwtSecret")

	mockToken.On(
		"SignedString",
		[]byte("jwtSecret"),
	).Return("tokenString", nil)

	user, token, err := ts.userService.Login(credentials)

	assert.NoError(t, err)
	assert.Equal(t, "test@mail.com", user.Email)
	assert.Equal(t, "tokenString", token)
}

func TestLoginUser_DatabaseError(t *testing.T) {
	ts := setup(t)

	credentials := services.UserCredentials{
		Email:    "test@mail.com",
		Password: "password",
	}

	ts.mockRepo.On(
		"GetUserByEmail",
		credentials.Email,
	).Return(models.User{}, errors.New("Database error"))

	_, _, err := ts.userService.Login(credentials)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "Database error")
}

func TestLoginUser_PasswordMissmatch(t *testing.T) {
	ts := setup(t)

	credentials := services.UserCredentials{
		Email:    "test@mail.com",
		Password: "password",
	}

	ts.mockRepo.On(
		"GetUserByEmail",
		credentials.Email,
	).Return(models.User{
		Model:    gorm.Model{ID: 1},
		Email:    credentials.Email,
		Password: "hashedPassword",
	}, nil)

	ts.mockBcrypt.On(
		"CompareHashAndPassword",
		[]byte("hashedPassword"),
		[]byte(credentials.Password),
	).Return(errors.New("Password missmatch"))

	_, _, err := ts.userService.Login(credentials)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "Password missmatch")
}

func TestLoginUser_TokenError(t *testing.T) {
	ts := setup(t)
	mockToken := new(interfaces_mocks.MockJwtTokenInterface)

	credentials := services.UserCredentials{
		Email:    "test@mail.com",
		Password: "password",
	}

	ts.mockRepo.On(
		"GetUserByEmail",
		credentials.Email,
	).Return(models.User{
		Model:    gorm.Model{ID: 1},
		Email:    credentials.Email,
		Password: "hashedPassword",
	}, nil)

	ts.mockBcrypt.On(
		"CompareHashAndPassword",
		[]byte("hashedPassword"),
		[]byte(credentials.Password),
	).Return(nil)

	ts.mockJwt.On(
		"NewWithClaims",
		jwt.SigningMethodHS256,
		mock.AnythingOfType("jwt.MapClaims"),
	).Return(mockToken)

	ts.mockConfig.On(
		"GetEnv",
		"JWT_SECRET",
	).Return("jwtSecret")

	mockToken.On(
		"SignedString",
		[]byte("jwtSecret"),
	).Return("", errors.New("Token error"))

	_, _, err := ts.userService.Login(credentials)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "Token error")
}
