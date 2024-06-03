package routes

import (
	"time"

	"hex/cms/pkg/config"
	"hex/cms/pkg/controllers"
	"hex/cms/pkg/interfaces"
	"hex/cms/pkg/repositories"
	"hex/cms/pkg/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter (
	bcrypt interfaces.BcryptInterface,
	jwt interfaces.JwtInterface,
	config config.Config,
) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{config.GetEnv("DOMAIN")},
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
    AllowHeaders:     []string{"Origin", "Content-Type"},
    AllowCredentials: true,
    MaxAge:           12 * time.Hour,
	}))

	userRepository := repositories.NewUserRepository()
	productRepository := repositories.NewProductRepository()

	userService := services.NewUserService(
		userRepository,
		bcrypt,
		jwt,
		config,
	)
	productService := services.NewProductService(
		productRepository,
	)

	userController := controllers.NewUserController(userService)
	productController := controllers.NewProductController(productService)

	api := r.Group("/api")

	userApi := api.Group("/users")

	{
		userApi.POST("", userController.CreateUser)
		userApi.POST("/login", userController.Login)
	}

	productApi := api.Group("/products")

	{
		productApi.POST("", productController.CreateProduct)
		productApi.GET("", productController.GetAllProducts)
		productApi.GET("/:id", productController.GetProductById)
	}

	return r
}