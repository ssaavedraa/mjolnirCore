package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex"`
	Domain      string
	Nit         *string `gorm:"uniqueIndex"`
	Address     string
	PhoneNumber string
	IsDraft     bool `gorm:"default:true"`

	Products    []Product
	TeamMembers []User
}
