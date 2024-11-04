package services

import "hex/mjolnir-core/pkg/models"

type OptionalCompanyInput struct {
	Id          uint    `json:"id"`
	Name        string  `json:"name"`
	Domain      string  `json:"domain"`
	Nit         *string `json:"nit"`
	Address     string  `json:"address"`
	PhoneNumber string  `json:"phoneNumber"`
}

type CompanyRoleInput struct {
	RoleName string `json:"roleName" binding:"required"`
}

type CompanyService interface {
	UpdateCompany(input OptionalCompanyInput) (*models.Company, error)
	GetCompanyRoles(companyId uint) ([]models.Role, error)
	CreateCompanyRole(companyId uint, roleName string) error
}
