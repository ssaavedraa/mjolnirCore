package models

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name      string `gorm:"not null"`
	CompanyID uint

	Company Company `gorm:"constraint:OnUpdate:CASCADE,onDelete:SET NULL;"`
	Users   []User  `gorm:"foreignKey:TeamID"`
}
