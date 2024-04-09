package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"santiagosaavedra.com.co/invoices/internal/api"
)

func main() {
	env := os.Getenv("ENVIRONMENT")
	if env == "development" {
		if error := godotenv.Load(); error != nil {
			log.Fatalf("Error loading .env file: %v", error)
		}
	}

	router := gin.Default()

	api.SetupRoutes(router)

	port := os.Getenv("PORT")
	address := ":" + port

	if error := http.ListenAndServe(address, router); error != nil {
		log.Fatalf("Failed to start server: %v", error)
	}
}