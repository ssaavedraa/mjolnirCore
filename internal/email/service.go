package email

import (
	"encoding/base64"
	"io"
	"log"
	"mime/multipart"
	"net/smtp"
	"os"
)

type Service struct {
	smtpFactory SMTPClientFactory
}

func NewService(factory SMTPClientFactory) *Service {
	return &Service{smtpFactory: factory}
}

func (service *Service) SendEmail(to, subject, body string, file multipart.File) error {
	senderEmail := os.Getenv("ICLOUD_SENDER_EMAIL")


	boundary := "boundary-123456789"
	message := []byte(
		"From: " + senderEmail + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: multipart/mixed; boundary=" + boundary + "\r\n\r\n" +

		"--" + boundary + "\r\n" +
		"Content-Type: text/plain; charset=utf-8\r\n" +
		"\r\n" +
		body + "\r\n",
	)

	fileContent, fileError := io.ReadAll(file)

	if fileError != nil {
		log.Printf("Failed to read file content")
		return fileError
	}

	encodedContent := base64.StdEncoding.EncodeToString(fileContent)

	attachment := []byte(
		"--" + boundary + "\r\n" +
		"Content-Type: application/pdf; name=\"invoice.pdf\"\r\n" +
		"Content-Disposition: attachment; filename=\"invoice.pdf\"\r\n" +
		"Content-Transfer-Encoding: base64\r\n" +
		"\r\n" +
		encodedContent + "\r\n" +
		"--" + boundary + "--\r\n",
	)

	message = append(message, attachment...)

	auth := service.smtpFactory.NewSMTPClient()

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	error := smtp.SendMail(smtpHost + ":" + smtpPort, auth, senderEmail, []string{to}, []byte(message))

	if error != nil {
		log.Printf("Failed to send email: %v", error)
		return error
	}

	log.Printf("Email sent to %s", to)
	return nil
}