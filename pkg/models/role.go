package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name      string
	CompanyID uint

	Company Company `gorm:"constraint:OnDelete:CASCADE;"`
	User    []User  `gorm:"foreignKey:RoleID"`
}
