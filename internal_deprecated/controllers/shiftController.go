package controller

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"hex/cms/internal_deprecated/model"
	shiftService "hex/cms/internal_deprecated/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ClockIn(c *gin.Context) {
	//TODO: Migrate to services layer
	user, _ := c.Get("user")

	database, err := db.GetInstance()

	if err != nil {
		log.Printf("Database unavailable")

		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Database unavailable"})
		return
	}

	currentDate := time.Now().Format("2006-01-02")

	var userShift model.Shift
	if err := database.Where("user_id = ? AND to_char(clock_in, 'YYYY-MM-DD') = ?", user.(model.User).ID, currentDate).First(&userShift).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			database.Create(&model.Shift{UserID: user.(model.User).ID})

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
	//TODO: Migrate to services layer
	user, _ := c.Get("user")

	db, err := db.GetInstance()

	if err != nil {
		log.Printf("Database unavailable")

		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Database unavailable"})
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
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving shift"})
		return
	}

	if err := db.Model(&userShift).Update("clock_out", time.Now()).Error; err != nil {
		log.Printf("Error updating clock_out: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error updating clock_out"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "See you next time!"})
}

func GetAll(c *gin.Context) {
	user, _ := c.Get("user")
	daysForwardQuery := c.Query("df")
	daysBackwardQuery := c.Query("db")

	if daysForwardQuery == "" && daysBackwardQuery == "" {
		//TODO: Migrate to services layer
		db, err := db.GetInstance()

		if err != nil {
			log.Printf("Database unavailable: %v", err)

			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Database unavailable"})
			return
		}

		var shifts []model.Shift
		db.Where("user_id = ?", user.(model.User).ID).Find(&shifts)

		c.IndentedJSON(http.StatusOK, shifts)
	}

	daysForward, _ := strconv.Atoi(daysForwardQuery)

	if daysForwardQuery != "" && daysForward <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Days forward must be a positive integer",
		})

		return
	}

	daysBackward, _ := strconv.Atoi(daysBackwardQuery)

	if daysBackwardQuery != "" && daysBackward <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Days backward must be a positive integer",
		})

		return
	}

	shifts, err := shiftService.GetShiftsInRage(daysBackward, daysForward, int(user.(model.User).ID))

	if err != nil {
		log.Printf("Failed to fetch shifts in range: %v", err)

		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch shifts in range",
		})
	}

	c.IndentedJSON(http.StatusOK, shifts)
}
