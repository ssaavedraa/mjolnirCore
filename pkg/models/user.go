package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email       string `gorm:"uniqueIndex"`
	Password    string
	Fullname    string
	PhoneNumber string
	Address     string
	CompanyRole string
	IsDraft     bool   `gorm:"default:false"`
	InviteId    string `gorm:"default:null"`

	CompanyID uint
	Shifts    []Shift   `gorm:"foreignKey:UserID"`
	Invoices  []Invoice `gorm:"foreignKey:UserID"`
	Products  []Product `gorm:"foreignKey:CreatedBy"`
	Company   Company   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
