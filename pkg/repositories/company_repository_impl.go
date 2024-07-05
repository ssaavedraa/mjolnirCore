package repositories

import (
	"hex/mjolnir-core/pkg/config"
	"hex/mjolnir-core/pkg/models"
)

type CompanyRepositoryImpl struct{}

func NewCompanyRepository() CompanyRepository {
	return &CompanyRepositoryImpl{}
}

func (repo *CompanyRepositoryImpl) FindByNameOrCreate(name string) (models.Company, error) {
	var company = models.Company{
		Name: name,
	}

	result := config.DB.Where("name = ?", name).FirstOrCreate(&company)

	if result.Error != nil {
		return models.Company{}, result.Error
	}

	return company, nil
}

func (repo *CompanyRepositoryImpl) Update(company models.Company) (models.Company, error) {
	var existingCompany models.Company

	existingCompanyResult := config.DB.First(&existingCompany, company.ID)

	if existingCompanyResult.Error != nil {
		return existingCompany, existingCompanyResult.Error
	}

	updatedCompanyResult := config.DB.Model(&existingCompany).Updates(company)

	if updatedCompanyResult.Error != nil {
		return existingCompany, updatedCompanyResult.Error
	}

	return company, nil
}
