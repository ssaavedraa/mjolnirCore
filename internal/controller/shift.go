package controller

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"santiagosaavedra.com.co/invoices/internal/db"
	"santiagosaavedra.com.co/invoices/internal/model"
)

func ClockIn(c *gin.Context) {
	user, _ := c.Get("user")

	database, err := db.GetInstance()

	if err != nil {
		log.Printf("Database unavailable")

		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message" : "Database unavailable"})
		return
	}

	currentDate := time.Now().Format("2006-01-02")

	var userShift model.Shift
	if err := database.Where("user_id = ? AND to_char(clock_in, 'YYYY-MM-DD') = ?", user.(model.User).ID, currentDate).First(&userShift).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			database.Create(&userShift)

			c.IndentedJSON(http.StatusCreated, gin.H{"message": "Welcome back!"})
			return
		}

		log.Printf("Error retrieving shift: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving shift"})
		return
	}

	c.IndentedJSON(http.StatusAlreadyReported, gin.H{"message": "Already clocked in"})
}

func ClockOut(c *gin.Context) {
	user, _ := c.Get("user")

	db, err := db.GetInstance()

	if err != nil {
		log.Printf("Database unavailable")

		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message" : "Database unavailable"})
		return
	}

	currentDate := time.Now().Format("2006-01-02")

	var userShift model.Shift

	if err := db.Where("user_id = ? AND to_char(clock_in, 'YYYY-MM-DD') LIKE ? AND clock_out IS NULL", user.(model.User).ID, currentDate).First(&userShift).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Shift not found for user with ID %d", userShift.UserID)

			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Shift not found, have you clocked in?"})
			return
		}

		log.Printf("Error retrieving shift: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"megin jwt authssage": "Error retrieving shift"})
		return
	}

	if err := db.Model(&userShift).Update("clock_out", time.Now()).Error; err != nil {
		log.Printf("Error updating clock_out: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error updating clock_out"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "See you next time!"})
}