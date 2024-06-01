package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"santiagosaavedra.com.co/invoices/pkg/config"
	"santiagosaavedra.com.co/invoices/pkg/routes"
)

func main() {
	config.LoadConfig()

	env := config.GetEnv("ENVIRONMENT")

	if env == "development" {
		if error := godotenv.Load(); error != nil {
			log.Fatalf("Error loading .env file: %v", error)
		}
	}

	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := routes.SetupRouter()

	port := config.GetEnv("PORT")

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}