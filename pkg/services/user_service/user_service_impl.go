package services

import (
	"hex/cms/pkg/config"
	"hex/cms/pkg/models"
	repositories "hex/cms/pkg/repositories/user_repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
}

func NewUserService (userRepository repositories.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (us *UserServiceImpl) CreateUser (input UserInput) (models.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)

	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Email: input.Email,
		Password: string(hash),
		Fullname: input.Fullname,
		PhoneNumber: input.PhoneNumber,
		Address: input.Address,
	}

	createdUser, err := us.UserRepository.CreateUser(user)

	if err != nil {
		return models.User{}, err
	}

	return createdUser, nil
}

func (us *UserServiceImpl) Login (credentials UserCredentials) (models.User, string, error) {
	user, err := us.UserRepository.GetUserByEmail(credentials.Email)

	if err != nil {
		return models.User{}, "",err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(credentials.Password),
	); err != nil {
		return models.User{}, "", err
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": user.ID,
			"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	)

	jwtSecret := config.GetEnv("JWT_SECRET")

	tokenString, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return models.User{}, "", err
	}

	return user, tokenString, nil
}