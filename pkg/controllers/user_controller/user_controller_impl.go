package controllers

import (
	"log"
	"net/http"

	services "hex/cms/pkg/services/user_service"
	"hex/cms/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) *UserControllerImpl {
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
		log.Println("Invalid request payload: ", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
		})

		return
	}

	createdUser, err := uc.UserService.CreateUser(userInput)

	if err != nil {
		log.Println("Failed to create user: ", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create user. Please try again later",
		})

		return
	}

	response := utils.ConvertToResponse(createdUser, utils.ResponseFields{
		"id": createdUser.ID,
		"fullname": createdUser.Fullname,
		"companyId": createdUser.Companies,
	})

	c.JSON(http.StatusCreated, response)
}