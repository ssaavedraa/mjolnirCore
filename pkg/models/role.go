package models

import (
	"strings"

	"gorm.io/gorm"
)

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

func (role *Role) BeforeSave(tx *gorm.DB) (err error) {
	role.Name = strings.ToLower(role.Name)
	return
}
