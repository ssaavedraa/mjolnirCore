package config

import (
	"fmt"
	"log"
	"os"

	"hex/cms/pkg/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type ConfigImpl struct {}

func NewConfig () Config {
	return &ConfigImpl{}
}

func (c *ConfigImpl) LoadConfig () {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.GetEnv("DB_HOST"),
		c.GetEnv("DB_USER"),
		c.GetEnv("DB_PASSWORD"),
		c.GetEnv("DB_NAME"),
		c.GetEnv("DB_PORT"),
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
		&models.Product{},
	)

	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
}

func (c *ConfigImpl) GetEnv (key string) string {
	return os.Getenv(key)
}