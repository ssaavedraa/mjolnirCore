package services

import (
	"encoding/json"
	"hex/mjolnir-core/pkg/config"
	"hex/mjolnir-core/pkg/interfaces"
	"hex/mjolnir-core/pkg/models"
	"hex/mjolnir-core/pkg/repositories"
	"hex/mjolnir-core/pkg/utils"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
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

func (us *UserServiceImpl) CreateUser(input UserInput, creationMethod string) (models.User, error) {
	emailTemplate := getEmailTemplateId(creationMethod)

	hash, err := us.Bcrypt.GenerateFromPassword([]byte(input.Password), 10)

	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		PhoneNumber: input.PhoneNumber,
		CompanyID:   input.CompanyId,
		Fullname:    input.Fullname,
		Address:     input.Address,
		Password:    string(hash),
		Email:       input.Email,
		IsDraft:     creationMethod == "hex-invite",
	}

	if creationMethod == "hex-invite" {
		user.Password = ""
		user.InviteId = utils.GenerateId()
	}

	createdUser, err := us.UserRepository.CreateUser(user)

	if err != nil {
		return models.User{}, err
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
			return models.User{}, err
		}

		err = us.EmailSender.Send(email)

		if err != nil {
			return models.User{}, err
		}
	}

	return createdUser, nil
}

func (us *UserServiceImpl) Login(credentials UserCredentials) (models.User, string, error) {
	user, err := us.UserRepository.GetUserByEmail(credentials.Email)

	if err != nil {
		return models.User{}, "", err
	}

	if err := us.Bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(credentials.Password),
	); err != nil {
		return models.User{}, "", err
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
		return models.User{}, "", err
	}

	return user, tokenString, nil
}

func (us *UserServiceImpl) InviteUser(invite UserInvite) (models.User, error) {
	company, err := us.CompanyRepository.FindByNameOrCreate(invite.CompanyName)

	if err != nil {
		return models.User{}, err
	}

	user := UserInput{
		CompanyId: company.ID,
		Email:     invite.Email,
		Fullname:  invite.Fullname,
	}

	createdUser, err := us.CreateUser(user, "hex-invite")

	if err != nil {
		return models.User{}, err
	}

	return createdUser, nil
}

func (us *UserServiceImpl) GetByInviteId(inviteId string) (models.User, error) {
	user, err := us.UserRepository.GetByInviteId(inviteId)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (us *UserServiceImpl) UpdateUser(input OptionalUserInput) (models.User, error) {
	hash, err := us.Bcrypt.GenerateFromPassword([]byte(input.Password), 10)

	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Model: gorm.Model{
			ID: input.Id,
		},
		Email:       input.Email,
		Password:    string(hash),
		Fullname:    input.Fullname,
		PhoneNumber: input.PhoneNumber,
		Address:     input.Address,
		CompanyRole: input.CompanyRole,
		CompanyID:   input.CompanyId,
		IsDraft:     false,
	}

	updatedUser, err := us.UserRepository.Update(user)

	if err != nil {
		return models.User{}, err
	}

	return updatedUser, nil
}

func getEmailTemplateId(creationMethod string) string {
	switch creationMethod {
	case "hex-invite":
		return "user_invite_mvp"
	default:
		return "user_created"
	}
}
