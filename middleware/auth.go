package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"santiagosaavedra.com.co/invoices/internal/db"
	"santiagosaavedra.com.co/invoices/internal/model"
)

func Auth (c *gin.Context) {
	tokenString, err := c.Cookie("session")

	if err != nil {
		log.Printf("Authentication error: %v", err)

		c.AbortWithStatus(http.StatusUnauthorized)

		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		log.Printf("Token parsing error: %v", err)

		c.AbortWithStatus(http.StatusUnauthorized)

		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check token expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		// Find user
		var user model.User
		database, dbErr := db.GetInstance()

		if dbErr != nil {
			log.Printf("Database unavailable: %v", err)

			c.AbortWithStatus(http.StatusInternalServerError)

			return
		}

		if err := database.First(&user, claims["sub"]).Error; err != nil {
			log.Printf("Failed to authenticate: %v", err)

			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", &user)

		c.Next()
	} else {
		log.Printf("Authentication error: %v", err)

		c.AbortWithStatus(http.StatusUnauthorized)

		return
	}
}