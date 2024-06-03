package controllers

import (
	"errors"
	"net/http"

	"hex/cms/pkg/services"
	"hex/cms/pkg/utils"
	"hex/cms/pkg/utils/logging"

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

type UserResponse struct {
	ID uint `json:"id"`
	Fullname uint `json:"fullname"`
}

func (uc *UserControllerImpl) CreateUser (c *gin.Context) {
	var userInput services.UserInput

	if err := c.ShouldBindJSON(&userInput); err != nil {
		logging.Error(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
		})

		return
	}

	createdUser, err := uc.UserService.CreateUser(userInput)

	if err != nil {
		logging.Error(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create user. Please try again later",
		})

		return
	}

	response := utils.ConvertToResponse(createdUser, utils.ResponseFields{
		"id": createdUser.ID,
		"fullname": createdUser.Fullname,
		"companyId": createdUser.CompanyID,
	})

	c.JSON(http.StatusCreated, response)
}

func (uc *UserControllerImpl) Login (c *gin.Context) {
	var credentials services.UserCredentials

	if err := c.ShouldBindJSON(&credentials); err != nil {
		logging.Error(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
		})

		return
	}

	user, token, err := uc.UserService.Login(credentials)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid credentials",
			})

			return
		}
		logging.Error(err)

		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to login user. Please try again later",
		})

		return
	}

	response := utils.ConvertToResponse(user, utils.ResponseFields{
		"id": user.ID,
		"fullname": user.Fullname,
		"companyId": user.CompanyID,
		"token": token,
	})

	c.JSON(http.StatusOK, response)
}