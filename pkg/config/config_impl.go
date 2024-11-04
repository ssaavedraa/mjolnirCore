package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

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
	sqlDB, err := db.DB()

	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %w", err)
	}

	projectDir, err := os.Getwd()

	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	migrationDir := filepath.Join(projectDir, "pkg", "config", "db", "migrations")

	migrationFiles, err := os.ReadDir(migrationDir)

	if err != nil {
		return fmt.Errorf("failed to read migration directory: %w", err)
	}

	for _, file := range migrationFiles {
		if file.IsDir() {
			continue
		}

		if filepath.Ext(file.Name()) == ".sql" && file.Name() == file.Name()[:len(file.Name())-7]+".up.sql" {
			content, err := os.ReadFile(filepath.Join(migrationDir, file.Name()))

			if err != nil {
				return fmt.Errorf("failed to read migration file %s: %w", file.Name(), err)
			}
			_, err = sqlDB.Exec(string(content))

			if err != nil {
				return fmt.Errorf("failed to execute migration %s: %w", file.Name(), err)
			}

			log.Printf("Executed migration: %s", file.Name())
		}
	}

	return nil
}

func rollbackMigration(db *gorm.DB, migrationName string) error {
	sqlDB, err := db.DB()

	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %w", err)
	}

	projectDir, err := os.Getwd()

	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	migrationDir := filepath.Join(projectDir, "pkg", "config", "db", "migrations")

	migrationFiles, err := os.ReadDir(migrationDir)

	if err != nil {
		return fmt.Errorf("failed to read migration directory: %w", err)
	}

	for _, file := range migrationFiles {
		if file.IsDir() {
			continue
		}

		if filepath.Ext(file.Name()) == ".sql" && file.Name() == file.Name()[:len(file.Name())-9]+".down.sql" {
			content, err := os.ReadFile(filepath.Join(migrationDir, file.Name()))

			if err != nil {
				return fmt.Errorf("failed to read migration file %s: %w", file.Name(), err)
			}
			_, err = sqlDB.Exec(string(content))

			if err != nil {
				return fmt.Errorf("failed to execute migration %s: %w", file.Name(), err)
			}

			log.Printf("Executed rollback: %s", file.Name())
		}
	}

	return nil
}
