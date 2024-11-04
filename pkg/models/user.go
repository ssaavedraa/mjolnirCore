package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Fullname    string
	Email       string `gorm:"uniqueIndex"`
	PhoneNumber string
	Address     string
	Password    string
	IsDraft     bool   `gorm:"default:false"`
	InviteId    string `gorm:"default:null"`
	RoleID      uint

	CompanyID uint
	TeamID    uint
	Role      Role
	Team      Team      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:TeamID"`
	Company   Company   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Products  []Product `gorm:"foreignKey:CreatedBy"`
}
