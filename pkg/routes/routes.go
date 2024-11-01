package routes

import (
	"time"

	"hex/mjolnir-core/pkg/config"
	"hex/mjolnir-core/pkg/controllers"
	"hex/mjolnir-core/pkg/interfaces"
	"hex/mjolnir-core/pkg/repositories"
	"hex/mjolnir-core/pkg/services"
	"hex/mjolnir-core/pkg/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	bcrypt interfaces.BcryptInterface,
	jwt interfaces.JwtInterface,
	config config.Config,
) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowOrigins:     []string{config.GetEnv("FRONTEND_DOMAIN")},
		MaxAge:           12 * time.Hour,
		AllowCredentials: true,
	}))

	emailSender := utils.NewEmailSender(config)

	productRepository := repositories.NewProductRepository()
	companyRepository := repositories.NewCompanyRepository()
	userRepository := repositories.NewUserRepository()
	teamRepository := repositories.NewTeamRepository()

	userService := services.NewUserService(
		companyRepository,
		userRepository,
		bcrypt,
		jwt,
		config,
		emailSender,
	)
	companyService := services.NewCompanyService(
		companyRepository,
	)
	productService := services.NewProductService(
		productRepository,
	)
	teamService := services.NewTeamService(
		teamRepository,
	)

	userController := controllers.NewUserController(userService)
	companyController := controllers.NewCompanyController(companyService)
	productController := controllers.NewProductController(productService)
	teamController := controllers.NewTeamController(teamService)

	api := r.Group("/api")

	userApi := api.Group("/users")

	{
		userApi.POST("", userController.CreateUser)
		userApi.POST("/login", userController.Login)
		userApi.GET("/invite/:inviteId", userController.GetByInviteId)
		userApi.PUT("/:id", userController.UpdateUser)
	}

	companyApi := api.Group("/companies")

	{
		companyApi.PUT("", companyController.UpdateCompany)
	}

	productApi := api.Group("/products")

	{
		productApi.POST("", productController.CreateProduct)
		productApi.GET("", productController.GetAllProducts)
		productApi.GET("/:id", productController.GetProductById)
	}

	teamApi := api.Group("/companies/:companyId/teams")

	{
		teamApi.GET("/", teamController.GetTeams)
		teamApi.GET("/:teamName/members", teamController.GetTeamMembers)
	}

	return r
}
