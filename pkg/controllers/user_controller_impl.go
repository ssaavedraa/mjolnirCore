package controllers

import (
	"errors"
	"fmt"
	"net/http"

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
		var userInput services.UserInput

		if err := c.ShouldBindJSON(&userInput); err != nil {
			utils.RespondWithError(c, http.StatusBadRequest, "Invalid request Payload", err)
			// logging.Error(err)

			// c.JSON(http.StatusBadRequest, gin.H{
			// 	"message": "Invalid request payload",
			// })

			return
		}

		createdUser, err := uc.UserService.CreateUser(userInput, creationMethod)

		if err != nil {
			utils.RespondWithError(c, http.StatusBadRequest, "Failed to create user. Please try again later", err)
			// logging.Error(err)

			// c.JSON(http.StatusInternalServerError, gin.H{
			// 	"message": "Failed to create user. Please try again later",
			// })

			return
		}

		response := utils.ConvertToResponse(createdUser, utils.ResponseFields{
			"id":        createdUser.ID,
			"fullname":  createdUser.Fullname,
			"companyId": createdUser.CompanyID,
		})

		c.JSON(http.StatusCreated, response)
	}

	if creationMethod == "hex-invite" {
		var userInvite services.UserInvite

		if err := c.ShouldBindJSON(&userInvite); err != nil {
			utils.RespondWithError(c, http.StatusBadRequest, "Invalid request Payload", err)
			// logging.Error(err)

			// 	c.JSON(http.StatusBadRequest, gin.H{
			// 		"message": "Invalid request payload",
			// 	})

			return
		}

		fmt.Printf("userInvite: %v \n", userInvite)

		_, err := uc.UserService.InviteUser(userInvite)

		if err != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, "Failed to invite user. Please try again later", err)
			// logging.Error(err)

			// 	c.JSON(http.StatusInternalServerError, gin.H{
			// 		"message": "Failed to invite user. Please try again later",
			// 	})

			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}

func (uc *UserControllerImpl) Login(c *gin.Context) {
	var credentials services.UserCredentials

	if err := c.ShouldBindJSON(&credentials); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request Payload", err)
		// logging.Error(err)

		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"message": "Invalid request payload",
		// })

		return
	}

	user, token, err := uc.UserService.Login(credentials)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			utils.RespondWithError(c, http.StatusBadRequest, "Invalid credentials", err)
			// c.JSON(http.StatusBadRequest, gin.H{
			// 		"message": "Invalid credentials",
			// 	})

			return
		}
		// logging.Error(err)
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to login user. Please try again later", err)

		// c.IndentedJSON(http.StatusInternalServerError, gin.H{
		// 	"message": "Failed to login user. Please try again later",
		// })

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
		// logging.Error(err)

		// c.JSON(http.StatusInternalServerError, gin.H{
		// 	"message": "Failed to find user by invite id. Please try again later",
		// })
		return
	}

	response := utils.ConvertToResponse(user, utils.ResponseFields{
		"id":          user.ID,
		"email":       user.Email,
		"fullname":    user.Fullname,
		"phoneNumber": user.PhoneNumber,
		"address":     user.Address,
		"companyRole": user.CompanyRole,
		"company": utils.ResponseFields{
			"id":          user.CompanyID,
			"name":        user.Company.Name,
			"domain":      user.Company.Domain,
			"nit":         user.Company.Nit,
			"address":     user.Company.Address,
			"phoneNumber": user.Company.PhoneNumber,
		},
	})

	c.JSON(http.StatusOK, response)
}

func (uc *UserControllerImpl) UpdateUser(c *gin.Context) {
	var userInput services.OptionalUserInput

	if err := c.ShouldBindJSON(&userInput); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request Payload", err)
		// logging.Error(err)

		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"message": "Invalid request payload",
		// })

		return
	}

	_, err := uc.UserService.UpdateUser(userInput)

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to update user. Please try again later", err)
		// logging.Error(err)

		// c.JSON(http.StatusInternalServerError, gin.H{
		// 	"message": "Failed to update user. Please try again later",
		// })
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
