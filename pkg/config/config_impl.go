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
	} else if env == "development" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	DB = initDatabase(c)
}

func (c *ConfigImpl) GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s not set", key)
	}
	return value
}

func initDatabase(c *ConfigImpl) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.GetEnv("DB_HOST"),
		c.GetEnv("DB_USER"),
		c.GetEnv("DB_PASSWORD"),
		c.GetEnv("DB_NAME"),
		c.GetEnv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err := migrateDatabase(db); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	return db
}

func migrateDatabase(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&models.User{},
		&models.Company{},
		&models.Product{},
		&models.Team{},
	); err != nil {
		return fmt.Errorf("error during migration: %w", err)
	}

	// Drop and recreate the unique index with a partial constraint
	db.Exec("DROP INDEX IF EXISTS idx_companies_nit;")
	db.Exec("CREATE UNIQUE INDEX idx_companies_nit ON companies (nit) WHERE nit IS NOT NULL;")
	db.Exec("CREATE UNIQUE INDEX idx_roles_name_ci ON roles (LOWER(name));")

	return nil
}
