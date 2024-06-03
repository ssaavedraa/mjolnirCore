package main

import (
	"log"

	"hex/cms/pkg/config"
	"hex/cms/pkg/routes"
	utils "hex/cms/pkg/utils/wrappers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	config := config.NewConfig()

	config.LoadConfig()

	env := config.GetEnv("ENVIRONMENT")

	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	bcrypt := &utils.BcryptWrapper{}
	jwt := &utils.JwtWrapper{}

	r := routes.SetupRouter(bcrypt, jwt, config)

	port := config.GetEnv("PORT")

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}