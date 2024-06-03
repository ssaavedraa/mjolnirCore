package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `json:"email" binding:"required" gorm:"uniqueIndex"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Address string `json:"address" binding:"required"`

	CompanyID uint
	Shifts []Shift `gorm:"foreignKey:UserID"`
	Invoices []Invoice `gorm:"foreignKey:UserID"`
	Products []Product `gorm:"foreignKey:CreatedBy"`
	Company     Company `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}