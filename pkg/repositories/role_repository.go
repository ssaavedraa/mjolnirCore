package repositories

import "hex/mjolnir-core/pkg/models"

type RoleRepository interface {
	FindOrCreateRoleByName(roleName string, companyId uint) (*models.Role, error)
	GetCompanyRoles(companyId uint) ([]models.Role, error)
}
