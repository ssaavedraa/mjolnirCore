package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Fullname    string `json:"fullname"`
	Email       string `json:"email" gorm:"uniqueIndex"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
	Password    string `json:"-"`
	IsDraft     bool   `json:"-" gorm:"default:false"`
	InviteId    string `json:"-" gorm:"default:null"`

	RoleID    uint      `json:"-"`
	CompanyID uint      `json:"-"`
	TeamID    uint      `json:"-"`
	Role      Role      `json:"role"`
	Team      Team      `json:"team" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:TeamID"`
	Company   Company   `json:"company" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Products  []Product `json:"products,omitempty" gorm:"foreignKey:CreatedBy"`
}
