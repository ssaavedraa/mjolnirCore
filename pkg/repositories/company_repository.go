package repositories

import "hex/mjolnir-core/pkg/models"

type CompanyRepository interface {
	FindOrCreateCompanyByName(name string) (*models.Company, error)
	Update(company *models.Company) (*models.Company, error)
}
