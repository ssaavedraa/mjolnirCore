package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name      string
	CompanyID uint

	CompanyRoles []CompanyRole
}

type CompanyRole struct {
	gorm.Model
	CompanyID uint
	RoleId    uint

	Users   []User  `gorm:"foreignKey:CompanyRoleId"`
	Company Company `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Role    Role    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
