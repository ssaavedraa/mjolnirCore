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
