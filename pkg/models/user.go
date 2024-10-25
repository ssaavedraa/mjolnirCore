package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Fullname    string
	Email       string `gorm:"uniqueIndex"`
	PhoneNumber string
	Address     string
	Password    string
	CompanyRole string
	IsDraft     bool   `gorm:"default:false"`
	InviteId    string `gorm:"default:null"`

	CompanyID uint
	TeamID    uint
	Products  []Product `gorm:"foreignKey:CreatedBy"`
	Company   Company   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Team      Team      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:TeamID"`
}
