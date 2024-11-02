package repositories

import "hex/mjolnir-core/pkg/models"

type RoleRepository interface {
	FindOrCreateRoleByName(roleName string) (*models.Role, error)
}
