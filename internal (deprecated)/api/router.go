package api

import (
	"hex/cms/internal/controller"
	"hex/cms/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Email routes
	router.POST("/email", middleware.Auth, SendEmailHandler)
	// Shift routes
	router.POST("/shifts/clock-in", middleware.Auth,  controller.ClockIn)
	router.POST("/shifts/clock-out", middleware.Auth, controller.ClockOut)
	router.GET("/shifts", middleware.Auth, controller.GetAll)
	// User routes
	router.POST("/users/sign-up", controller.SignUp)
	router.GET("/users/me", middleware.Auth, controller.GetUserDetails)
	// Auth routes
	router.POST("/auth/login", controller.LogIn)
	router.POST("/auth/logout", controller.LogOut)
}