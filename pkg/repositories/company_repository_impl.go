package repositories

import (
	"hex/mjolnir-core/pkg/config"
	"hex/mjolnir-core/pkg/models"

	"gorm.io/gorm"
)

type CompanyRepositoryImpl struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return &CompanyRepositoryImpl{
		db: db,
	}
}

func (repo *CompanyRepositoryImpl) FindByNameOrCreate(name string) (models.Company, error) {
	var company = models.Company{
		Name: name,
	}

	err := repo.db.Where("name = ?", name).FirstOrCreate(&company).Error

	if err != nil {
		return models.Company{}, err
	}

	return company, nil
}

func (repo *CompanyRepositoryImpl) Update(company models.Company) (models.Company, error) {
	var existingCompany models.Company

	if err := repo.db.First(&existingCompany, company.ID).Error; err != nil {
		return models.Company{}, err
	}

	if err := config.DB.Model(&existingCompany).Updates(company).Error; err != nil {
		return models.Company{}, err
	}

	return company, nil
}

func (repo *CompanyRepositoryImpl) GetCompanyRoles(companyId uint) ([]models.CompanyRole, error) {
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
