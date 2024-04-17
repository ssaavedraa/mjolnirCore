package api

import (
	"github.com/gin-gonic/gin"
	"santiagosaavedra.com.co/invoices/internal/controller"
	"santiagosaavedra.com.co/invoices/middleware"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/email", middleware.Auth, SendEmailHandler)
	router.POST("/clock-in", middleware.Auth,  controller.ClockIn)
	router.POST("/clock-out", middleware.Auth, controller.ClockOut)
	router.POST("/sign-up", controller.SignUp)
	router.POST("/log-in", controller.LogIn)
}