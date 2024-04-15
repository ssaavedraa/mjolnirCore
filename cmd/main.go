package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"santiagosaavedra.com.co/invoices/internal/api"
	"santiagosaavedra.com.co/invoices/internal/db"
	"santiagosaavedra.com.co/invoices/internal/model"
)

func main() {
	env := os.Getenv("ENVIRONMENT")

	if env == "development" {
		if error := godotenv.Load(); error != nil {
			log.Fatalf("Error loading .env file: %v", error)
		}
	}

	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	frontendUrl := os.Getenv("FRONTEND_URL")

	db, error := db.GetInstance()

	if error != nil {
		log.Fatalf("Database connection failed: %v", error)
	}

	db.AutoMigrate(&model.User{}, &model.Company{}, &model.Invoice{}, &model.Shift{})

	router := gin.Default()

	router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{frontendUrl},
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
    AllowHeaders:     []string{"Origin", "Content-Type"},
    AllowCredentials: true,
    MaxAge:           12 * time.Hour,
}))

	api.SetupRoutes(router)

	port := os.Getenv("PORT")
	address := ":" + port

	if error := http.ListenAndServe(address, router); error != nil {
		log.Fatalf("Failed to start server: %v", error)
	}
}