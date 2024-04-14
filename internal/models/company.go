package model

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name string
	Email string
	Address string
	Abn string
	PhoneNumber string
	UserId uint
}