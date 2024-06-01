package api

import (
	"log"
	"mime/multipart"
	"net/http"

	"hex/cms/internal_deprecated/email"

	"github.com/gin-gonic/gin"
)

type Email struct {
	To string `form:"to" binding:"required"`
	Cc string `form:"cc" binding:"required"`
	Subject string `form:"subject" binding:"required"`
	Body string `form:"body" binding:"required"`
	File *multipart.FileHeader `form:"file" binding:"required"`
}

func SendEmailHandler (context *gin.Context) {
	var newEmail Email


	if error := context.Bind(&newEmail); error != nil {
		log.Printf("Invalid request body: %v", error)

		context.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "Invalid request body"})
		return
	}

	// Open the uploaded file
	file, err := newEmail.File.Open()

	if err != nil {
			log.Printf("Failed to open uploaded file: %v", err)
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to open uploaded file"})
			return
	}
	defer file.Close()

	factory := email.GetSMTPClientFactory()
	service := email.NewService(factory)

	error := service.SendEmail(newEmail.To, newEmail.Cc, newEmail.Subject, newEmail.Body, file)

	if error != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message" : "Failed to send email"})
		return
	}

	context.IndentedJSON(http.StatusCreated, gin.H{"message" : "Email sent successfully"})
}

func ClockInHandler (context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "clocked in!"})
}