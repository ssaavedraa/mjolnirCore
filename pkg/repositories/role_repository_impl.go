package repositories

import (
	"hex/mjolnir-core/pkg/models"
	"strings"

	"gorm.io/gorm"
)

type RoleRepositoryImpl struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &RoleRepositoryImpl{
		db: db,
	}
}

func (repo *RoleRepositoryImpl) FindOrCreateRoleByName(roleName string, companyId uint) (*models.Role, error) {
	var role models.Role

	result := repo.db.
		Where("LOWER(name) = ?", strings.ToLower(roleName)).
		First(&role)

	if result.Error == nil {
		return &role, nil
	} else if result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}

	newRole := models.Role{
		Name:      roleName,
		CompanyID: companyId,
	}

	if err := repo.db.Create(&newRole).Error; err != nil {
		return nil, err
	}

	return &newRole, nil
}

func (repo *RoleRepositoryImpl) GetCompanyRoles(companyId uint) ([]models.Role, error) {
	var companyRoles []models.Role

	if err := repo.db.
		Where("company_id = ?", companyId).
		Find(&companyRoles).Error; err != nil {
		return nil, err
	}

	return companyRoles, nil

}
