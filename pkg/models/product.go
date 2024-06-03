package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name string
	Description string
	Price float64
	ImageUrl string

	CreatedBy uint
	CompanyID uint
}