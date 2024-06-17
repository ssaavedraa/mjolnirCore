package controller

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"hex/mjolnir-core/internal_deprecated/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SignUp(c *gin.Context) {
	var body struct {
		Email       string `json:"email" binding:"required"`
		Password    string `json:"password" binding:"required"`
		Fullname    string `json:"fullname" binding:"required"`
		Abn         string `json:"abn" binding:"required"`
		PhoneNumber string `json:"phoneNumber" binding:"required"`
		Address     string `json:"address" binding:"required"`
	}

	if err := c.Bind(&body); err != nil {
		log.Printf("Invalid request body: %v", err)

		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		log.Printf("Failed to hash password: %v", err)

		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to hash password",
		})

		return
	}

	user := model.User{
		Email:       body.Email,
		Password:    string(hash),
		Fullname:    body.Fullname,
		Abn:         body.Abn,
		PhoneNumber: body.PhoneNumber,
		Address:     body.Address,
	}

	database, err := db.GetInstance()

	if err != nil {
		log.Printf("Database unavailable: %v", err)

		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Database unavailable",
		})

		return
	}

	if err := database.Create(&user).Error; err != nil {
		log.Printf(" Failed to create new user: %v", err)

		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create new user",
		})

		return
	}

	signedToken, err := getToken(&user)

	if err != nil {
		log.Printf("Failed to sign token: %v", err)

		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to sign token",
		})
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("session", signedToken, 3600*24*7, "", "", false, true)

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Welcome!"})
}

func LogIn(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.Bind(&credentials); err != nil {
		log.Printf("Invalid request body: %v", err)

		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})

		return
	}

	database, err := db.GetInstance()

	if err != nil {
		log.Printf("Database Unavailable: %v", err)

		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Database unavailable",
		})

		return
	}

	var user model.User
	if err := database.First(&user, "email = ?", credentials.Email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"message": "Invalid credentials",
			})

			return
		}

		log.Printf("Query error: %v", err)

		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})

		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Invalid credentials",
		})

		return
	}

	signedToken, tokenErr := getToken(&user)

	if tokenErr != nil {
		c.AbortWithStatus(500)
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("session", signedToken, 3600*24*7, "", "", false, true)

	c.IndentedJSON(http.StatusCreated, gin.H{})
}

func LogOut(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("session", "", -1, "", "", false, true)

	c.IndentedJSON(http.StatusCreated, gin.H{})
}

func GetUserDetails(c *gin.Context) {
	user, _ := c.Get("user")

	log.Println(user)

	c.IndentedJSON(http.StatusOK, user)
}

func getToken(user *model.User) (signedToken string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	jwtSecret := os.Getenv("JWT_SECRET")

	tokenString, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		log.Printf("Failed to create token: %v", err)

		return "", err
	}

	return tokenString, nil
}
