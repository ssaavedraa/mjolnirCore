package repositories

import "hex/mjolnir-core/pkg/models"

type CompanyRepository interface {
	FindByNameOrCreate(name string) (models.Company, error)
}
