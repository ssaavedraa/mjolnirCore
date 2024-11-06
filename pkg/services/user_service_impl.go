package services

import (
	"encoding/json"
	"fmt"
	"hex/mjolnir-core/pkg/config"
	"hex/mjolnir-core/pkg/dtos"
	"hex/mjolnir-core/pkg/interfaces"
	"hex/mjolnir-core/pkg/models"
	"hex/mjolnir-core/pkg/repositories"
	"hex/mjolnir-core/pkg/utils"
	"reflect"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserServiceImpl struct {
	UserRepository    repositories.UserRepository
	CompanyRepository repositories.CompanyRepository
	Bcrypt            interfaces.BcryptInterface
	Config            config.Config
	Jwt               interfaces.JwtInterface
	EmailSender       interfaces.EmailSender
}

func NewUserService(
	companyRepository repositories.CompanyRepository,
	userRepository repositories.UserRepository,
	bcrypt interfaces.BcryptInterface,
	jwt interfaces.JwtInterface,
	config config.Config,
	emailSender interfaces.EmailSender,
) UserService {
	return &UserServiceImpl{
		CompanyRepository: companyRepository,
		UserRepository:    userRepository,
		Bcrypt:            bcrypt,
		Config:            config,
		Jwt:               jwt,
		EmailSender:       emailSender,
	}
}

func (us *UserServiceImpl) CreateUser(input dtos.UserInput, creationMethod string) (*models.User, error) {
	emailTemplate := getEmailTemplateId(creationMethod)

	hash, err := us.Bcrypt.GenerateFromPassword([]byte(input.Password), 10)

	if err != nil {
		return nil, err
	}

	user := models.User{
		PhoneNumber: input.PhoneNumber,
		CompanyID:   input.CompanyId,
		Fullname:    input.Fullname,
		Address:     input.Address,
		Password:    string(hash),
		Email:       input.Email,
		IsDraft:     creationMethod == "hex-invite",
		TeamID:      input.TeamID,
		RoleID:      input.RoleId,
	}

	if creationMethod == "hex-invite" {
		user.Password = ""
		user.InviteId = utils.GenerateId()
	}

	createdUser, err := us.UserRepository.CreateUser(&user)

	if err != nil {
		return nil, err
	}

	if creationMethod == "hex-invite" {

		email := interfaces.Email{
			TemplateData: map[string]string{
				"RecipientName": strings.Split(createdUser.Fullname, " ")[0],
				"InviteId":      createdUser.InviteId,
			},
			SenderAddress:   "invoices@santiagosaavedra.com.co",
			ReceiverAddress: createdUser.Email,
			TemplateName:    emailTemplate,
			Subject:         "Welcome to Hex",
			Locale:          "en",
		}

		_, err = json.Marshal(email)

		if err != nil {
			return nil, err
		}

		err = us.EmailSender.Send(email)

		if err != nil {
			return nil, err
		}
	}

	return createdUser, nil
}

func (us *UserServiceImpl) Login(credentials dtos.UserCredentials) (*models.User, string, error) {
	user, err := us.UserRepository.GetUserByEmail(credentials.Email)

	if err != nil {
		return nil, "", err
	}

	if err := us.Bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(credentials.Password),
	); err != nil {
		return nil, "", err
	}

	token := us.Jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": user.ID,
			"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	)

	jwtSecret := us.Config.GetEnv("JWT_SECRET")

	tokenString, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return nil, "", err
	}

	return user, tokenString, nil
}

func (us *UserServiceImpl) InviteUser(invite dtos.UserInvite) (*models.User, error) {
	company, err := us.CompanyRepository.FindOrCreateCompanyByName(invite.CompanyName)

	if err != nil {
		return nil, err
	}

	user := dtos.UserInput{
		CompanyId: company.ID,
		Email:     invite.Email,
		Fullname:  invite.Fullname,
		TeamID:    invite.TeamID,
	}

	createdUser, err := us.CreateUser(user, "hex-invite")

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (us *UserServiceImpl) GetByInviteId(inviteId string) (*models.User, error) {
	user, err := us.UserRepository.GetUserByInviteId(inviteId)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserServiceImpl) UpdateUser(input dtos.OptionalUserInput) (*models.User, error) {
	existingUser, err := us.UserRepository.GetUserById(*input.Id)

	if err != nil {
		return nil, fmt.Errorf("error retieving user: %w", err)
	}

	inputValue := reflect.ValueOf(input)
	existingValue := reflect.ValueOf(&existingUser).Elem()

	for i := 0; i < inputValue.NumField(); i++ {
		inputField := inputValue.Field(i)
		fieldType := inputValue.Type().Field(i)

		if inputField.IsZero() {
			continue
		}

		if fieldType.Name == "Password" {
			hashedPassword, err := us.Bcrypt.GenerateFromPassword([]byte(*input.Password), 10)

			if err != nil {
				return nil, fmt.Errorf("error hashing password: %w", err)
			}

			existingValue.FieldByName(fieldType.Name).SetString(string(hashedPassword))
			continue
		}

		existingValue.FieldByName(fieldType.Name).Set(inputField)
	}

	updatedUser, err := us.UserRepository.UpdateUser(existingUser)

	if err != nil {
		return nil, fmt.Errorf("error updating user: %w", err)
	}

	return updatedUser, nil
}

func (us *UserServiceImpl) GetUserById(userId uint) (*models.User, error) {
	user, err := us.UserRepository.GetUserById(userId)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func getEmailTemplateId(creationMethod string) string {
	switch creationMethod {
	case "hex-invite":
		return "user_invite_mvp"
	default:
		return "user_created"
	}
}
