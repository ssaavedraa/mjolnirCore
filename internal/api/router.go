package api

import (
	"github.com/gin-gonic/gin"
	"santiagosaavedra.com.co/invoices/internal/controller"
	"santiagosaavedra.com.co/invoices/middleware"
)

func SetupRoutes(router *gin.Engine) {
	// Email routes
	router.POST("/email", middleware.Auth, SendEmailHandler)
	// Shift routes
	router.POST("/shift/clock-in", middleware.Auth,  controller.ClockIn)
	router.POST("/shift/clock-out", middleware.Auth, controller.ClockOut)
	router.GET("/shift", middleware.Auth, controller.GetAll)
	// User routes
	router.POST("/user/sign-up", controller.SignUp)
	router.POST("/user/log-in", controller.LogIn)
	router.POST("/user/log-out", controller.LogOut)
	router.GET("/user/me", middleware.Auth, controller.GetUserDetails)
}