package repositories

import (
	"santiagosaavedra.com.co/invoices/pkg/config"
	"santiagosaavedra.com.co/invoices/pkg/models"
)

func CreateUser (user models.User) (models.User, error) {
	result := config.DB.Create(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}