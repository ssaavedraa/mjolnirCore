package repositories

import (
	"hex/mjolnir-core/pkg/models"

	"gorm.io/gorm"
)

type CompanyRoleRepositoryImpl struct {
	db *gorm.DB
}

func NewCompanyRoleRepository(db *gorm.DB) CompanyRoleRepository {
	return &CompanyRoleRepositoryImpl{
		db: db,
	}
}

func (repo *CompanyRoleRepositoryImpl) GetCompanyRoles(companyId uint) ([]models.CompanyRole, error) {
	var companyRoles []models.CompanyRole

	err := repo.db.Debug().
		Joins("JOIN roles ON roles.id = company_roles.role_id").
		Where("company_roles.company_id = ?", companyId).
		Preload("Role").
		Find(&companyRoles).Error

	if err != nil {
		return nil, err
	}

	return companyRoles, nil
}

func (repo *CompanyRoleRepositoryImpl) CreateCompanyRole(companyRole *models.CompanyRole) error {
	if err := repo.db.
		Create(&companyRole).Error; err != nil {
		return err
	}

	return nil
}
