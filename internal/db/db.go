package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error


func GetInstance() (*gorm.DB, error){
	if db == nil {

		// Setup database
		dbUser := os.Getenv("POSTGRES_USER")
		dbPassword := os.Getenv("POSTGRES_PASSWORD")
		dbName := os.Getenv("POSTGRES_DB")

		dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5434 sslmode=disable TimeZone=Australia/Melbourne", dbUser, dbPassword, dbName)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			return nil, err
		}

	}

	return db, nil
}