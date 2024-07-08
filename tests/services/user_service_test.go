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
	mockRepo        *repositories_mocks.MockUserRepository
	mockCompanyRepo *repositories_mocks.MockCompanyRepository
	mockKafka       *interfaces_mocks.MockKafkaProducerInterface
	mockBcrypt      *interfaces_mocks.MockBcryptInterface
	mockJwt         *interfaces_mocks.MockJwtInterface
	mockConfig      *config_mocks.MockConfig
	userService     services.UserService
}

func setup(_ *testing.T) *TestSetup {
	mockKafka := new(interfaces_mocks.MockKafkaProducerInterface)
	mockBcrypt := new(interfaces_mocks.MockBcryptInterface)
	mockJwt := new(interfaces_mocks.MockJwtInterface)

	mockRepo := new(repositories_mocks.MockUserRepository)
	mockCompanyRepo := new(repositories_mocks.MockCompanyRepository)

	mockConfig := new(config_mocks.MockConfig)

	userService := services.NewUserService(
		mockKafka,
		mockCompanyRepo,
		mockRepo,
		mockBcrypt,
		mockJwt,
		mockConfig,
	)

	return &TestSetup{
		mockCompanyRepo: mockCompanyRepo,
		userService:     userService,
		mockBcrypt:      mockBcrypt,
		mockConfig:      mockConfig,
		mockKafka:       mockKafka,
		mockRepo:        mockRepo,
		mockJwt:         mockJwt,
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

	creationMethod := ""

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

	ts.mockKafka.On(
		"SendMessageToKafka",
		"new_email",
		mock.AnythingOfType("[]uint8"),
	).Return(nil)

	createdUser, err := ts.userService.CreateUser(input, creationMethod)

	assert.NoError(t, err)
	assert.Equal(t, "test@mail.com", createdUser.Email)
	assert.Equal(t, "password", createdUser.Password)
	assert.Equal(t, "Test Name", createdUser.Fullname)
	assert.Equal(t, "1234567890", createdUser.PhoneNumber)
	assert.Equal(t, "Test Address", createdUser.Address)
}

func TestCreateUser_KafkaError(t *testing.T) {
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

	ts.mockKafka.On(
		"SendMessageToKafka",
		"new_email",
		mock.AnythingOfType("[]uint8"),
	).Return(errors.New("Kafka error"))

	_, err := ts.userService.CreateUser(input, "")

	assert.Error(t, err)
	assert.ErrorContains(t, err, "Kafka error")
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

	_, err := ts.userService.CreateUser(input, "")

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

	_, err := ts.userService.CreateUser(input, "")

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

func TestInviteUser_Success(t *testing.T) {
	ts := setup(t)

	invite := services.UserInvite{
		Email:       "test@mail.com",
		CompanyName: "test company",
		Fullname:    "test user",
	}

	ts.mockCompanyRepo.On(
		"FindByNameOrCreate",
		invite.CompanyName,
	).Return(models.Company{
		Model: gorm.Model{
			ID: 1,
		},
		Name:    invite.CompanyName,
		IsDraft: true,
	}, nil)

	ts.mockRepo.On(
		"CreateUser",
		mock.AnythingOfType("models.User"),
	).Return(models.User{
		Model:     gorm.Model{ID: 1},
		Email:     invite.Email,
		Fullname:  invite.Fullname,
		IsDraft:   true,
		CompanyID: 1,
		Company: models.Company{
			Model: gorm.Model{
				ID: 1,
			},
			Name:    invite.CompanyName,
			IsDraft: true,
		},
	}, nil)

	ts.mockBcrypt.On(
		"GenerateFromPassword",
		[]byte{},
		10,
	).Return([]byte("hashedPassword"), nil)

	ts.mockKafka.On(
		"SendMessageToKafka",
		"new_email",
		mock.AnythingOfType("[]uint8"),
	).Return(nil)

	createdUser, err := ts.userService.InviteUser(invite)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), createdUser.CompanyID)
	assert.Equal(t, true, createdUser.IsDraft)
	assert.Equal(t, "test@mail.com", createdUser.Email)
	assert.Equal(t, "test user", createdUser.Fullname)
}

func TestInviteUser_CompanyRepositoryError(t *testing.T) {
	ts := setup(t)

	invite := services.UserInvite{
		Email:       "test@mail.com",
		CompanyName: "test company",
		Fullname:    "test user",
	}

	ts.mockCompanyRepo.On(
		"FindByNameOrCreate",
		invite.CompanyName,
	).Return(models.Company{}, errors.New("Company repository error"))

	_, err := ts.userService.InviteUser(invite)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "Company repository error")
}

func TestInviteUser_UserRepositoryError(t *testing.T) {
	ts := setup(t)

	invite := services.UserInvite{
		Email:       "test@mail.com",
		CompanyName: "test company",
		Fullname:    "test user",
	}

	ts.mockCompanyRepo.On(
		"FindByNameOrCreate",
		invite.CompanyName,
	).Return(models.Company{
		Model: gorm.Model{
			ID: 1,
		},
		Name:    invite.CompanyName,
		IsDraft: true,
	}, nil)

	ts.mockRepo.On(
		"CreateUser",
		mock.AnythingOfType("models.User"),
	).Return(
		models.User{},
		errors.New("User repository error"),
	)

	ts.mockBcrypt.On(
		"GenerateFromPassword",
		[]byte{},
		10,
	).Return([]byte("hashedPassword"), nil)

	ts.mockKafka.On(
		"SendMessageToKafka",
		"new_email",
		mock.AnythingOfType("[]uint8"),
	).Return(nil)

	_, err := ts.userService.InviteUser(invite)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "User repository error")
}
