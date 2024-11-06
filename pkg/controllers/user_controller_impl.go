package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"hex/mjolnir-core/pkg/dtos"
	"hex/mjolnir-core/pkg/services"
	"hex/mjolnir-core/pkg/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserControllerImpl struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (uc *UserControllerImpl) CreateUser(c *gin.Context) {
	creationMethod := c.Query("method")

	if creationMethod == "" {
		var userInput dtos.UserInput

		if err := c.ShouldBindJSON(&userInput); err != nil {
			utils.RespondWithError(c, http.StatusBadRequest, "Invalid request Payload", err)

			return
		}

		_, err := uc.UserService.CreateUser(userInput, creationMethod)

		if err != nil {
			utils.RespondWithError(c, http.StatusBadRequest, "Failed to create user. Please try again later", err)

			return
		}

		c.JSON(http.StatusCreated, gin.H{})
	}

	if creationMethod == "hex-invite" {
		var userInvite dtos.UserInvite

		if err := c.ShouldBindJSON(&userInvite); err != nil {
			utils.RespondWithError(c, http.StatusBadRequest, "Invalid request Payload", err)

			return
		}

		_, err := uc.UserService.InviteUser(userInvite)

		if err != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, "Failed to invite user. Please try again later", err)

			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}

func (uc *UserControllerImpl) Login(c *gin.Context) {
	var credentials dtos.UserCredentials

	if err := c.ShouldBindJSON(&credentials); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request Payload", err)

		return
	}

	user, token, err := uc.UserService.Login(credentials)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			utils.RespondWithError(c, http.StatusBadRequest, "Invalid credentials", err)

			return
		}

		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to login user. Please try again later", err)

		return
	}

	response := utils.ConvertToResponse(user, utils.ResponseFields{
		"id":          user.ID,
		"name":        user.Fullname,
		"companyId":   user.CompanyID,
		"accessToken": token,
	})

	c.JSON(http.StatusCreated, response)
}

func (uc *UserControllerImpl) GetByInviteId(c *gin.Context) {
	inviteIdParam := c.Param("inviteId")

	user, err := uc.UserService.GetByInviteId(inviteIdParam)

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to find user by invite id. Please try again later", err)

		return
	}

	userResponse := dtos.UserResponse{
		ID:          &user.ID,
		Fullname:    &user.Fullname,
		Email:       &user.Email,
		PhoneNumber: &user.PhoneNumber,
		Address:     &user.Address,
		Role: &dtos.RoleResponse{
			Name: &user.Role.Name,
		},
		Company: &dtos.CompanyResponse{
			ID:          &user.CompanyID,
			Name:        &user.Company.Name,
			Domain:      &user.Company.Domain,
			Nit:         user.Company.Nit,
			Address:     &user.Company.Address,
			PhoneNumber: &user.Company.PhoneNumber,
		},
	}

	c.JSON(http.StatusOK, userResponse)
}

func (uc *UserControllerImpl) UpdateUser(c *gin.Context) {
	var userInput dtos.OptionalUserInput

	if err := c.ShouldBindJSON(&userInput); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request Payload", err)

		return
	}

	_, err := uc.UserService.UpdateUser(userInput)

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to update user. Please try again later", err)

		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (uc *UserControllerImpl) GetUserById(c *gin.Context) {
	var userIdParam = c.Param("userId")

	userId, err := strconv.ParseUint(userIdParam, 10, 64)

	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid url param", err)

		return
	}

	user, err := uc.UserService.GetUserById(uint(userId))

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to get user. Please try again later", err)

		return
	}

	userResponse := dtos.UserResponse{
		Fullname:    &user.Fullname,
		Email:       &user.Email,
		PhoneNumber: &user.PhoneNumber,
		Address:     &user.Address,
		Role: &dtos.RoleResponse{
			Name: &user.Role.Name,
		},
		Team: &dtos.TeamResponse{
			Name: &user.Team.Name,
		},
		Company: &dtos.CompanyResponse{
			Name: &user.Company.Name,
		},
	}

	c.JSON(http.StatusOK, userResponse)
}
