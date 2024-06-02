package routes

import (
	"time"

	"hex/cms/pkg/config"
	controllers "hex/cms/pkg/controllers/user_controller"
	repositories "hex/cms/pkg/repositories/user_repository"
	services "hex/cms/pkg/services/user_service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter () *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{config.GetEnv("DOMAIN")},
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
    AllowHeaders:     []string{"Origin", "Content-Type"},
    AllowCredentials: true,
    MaxAge:           12 * time.Hour,
	}))

	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	api := r.Group("/api")

	userApi := api.Group("/users")

	{
		userApi.POST("/create", userController.CreateUser)
		userApi.POST("/login", userController.Login)
	}

	return r
}