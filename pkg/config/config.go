package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"santiagosaavedra.com.co/invoices/pkg/models"
)

var DB *gorm.DB

func LoadConfig () {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		GetEnv("DB_HOST"),
		GetEnv("DB_USER"),
		GetEnv("DB_PASSWORD"),
		GetEnv("DB_NAME"),
		GetEnv("DB_PORT"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}


	err = DB.AutoMigrate(
		&models.User{},
		&models.Company{},
		&models.Invoice{},
		&models.Shift{},
		&models.InvoiceItem{},
	)

	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
}

func GetEnv (key string) string {
	return os.Getenv(key)
}