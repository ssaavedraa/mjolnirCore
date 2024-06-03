package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name string
	Domain string
	Nit string
	Address string
	PhoneNumber string

	Products []Product
	TeamMembers []User
}