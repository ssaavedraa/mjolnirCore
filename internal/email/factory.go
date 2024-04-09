package email

import (
	"net/smtp"
	"os"
)

type SMTPClientFactory interface {
	NewSMTPClient() smtp.Auth
}

type iCloudSMTPClientFactory struct {
	authEmail string
	password string
	host string
}

func (f *iCloudSMTPClientFactory) NewSMTPClient() smtp.Auth {
	return smtp.PlainAuth("", f.authEmail, f.password, f.host)
}

func GetSMTPClientFactory() SMTPClientFactory {
	return &iCloudSMTPClientFactory{
		authEmail: os.Getenv("ICLOUD_AUTH_EMAIL"),
		password: os.Getenv("ICLOUD_PASSWORD"),
		host: os.Getenv("SMTP_HOST"),
	}
}