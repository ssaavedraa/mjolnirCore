package repositories

import "hex/mjolnir-core/pkg/models"

type CompanyRoleRepository interface {
	GetCompanyRoles(companyId uint) ([]models.CompanyRole, error)
	CreateCompanyRole(companyRole *models.CompanyRole) error
}
