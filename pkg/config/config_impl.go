package config

import (
	"fmt"
	"log"
	"os"

	"hex/mjolnir-core/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type ConfigImpl struct{}

func NewConfig() Config {
	return &ConfigImpl{}
}

func (c *ConfigImpl) LoadConfig() {
	env := c.GetEnv("ENVIRONMENT")

	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	if env == "development" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.GetEnv("DB_HOST"),
		c.GetEnv("DB_USER"),
		c.GetEnv("DB_PASSWORD"),
		c.GetEnv("DB_NAME"),
		c.GetEnv("DB_PORT"),
	)

	fmt.Printf("dsn: %+v", dsn)

	DbInstance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	DB = DbInstance

	err = DB.AutoMigrate(
		&models.User{},
		&models.Company{},
		&models.Invoice{},
		&models.Shift{},
		&models.InvoiceItem{},
		&models.Product{},
	)

	// Drop the automatically created unique index
	DB.Exec("DROP INDEX IF EXISTS idx_companies_nit;")

	// Create the partial unique index using raw SQL
	DB.Exec("CREATE UNIQUE INDEX idx_companies_nit ON companies (nit) WHERE nit IS NOT NULL;")

	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
}

func (c *ConfigImpl) GetEnv(key string) string {
	return os.Getenv(key)
}
