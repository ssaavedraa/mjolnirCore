package services

import (
	"hex/mjolnir-core/pkg/models"
	"hex/mjolnir-core/pkg/repositories"

	"gorm.io/gorm"
)

type CompanyServiceImpl struct {
	CompanyRepository repositories.CompanyRepository
}

func NewCompanyService(
	companyRepository repositories.CompanyRepository,
) CompanyService {
	return &CompanyServiceImpl{
		CompanyRepository: companyRepository,
	}
}

func (cs *CompanyServiceImpl) UpdateDraftCompany(input OptionalCompanyInput) (models.Company, error) {
	company := models.Company{
		Model: gorm.Model{
			ID: input.Id,
		},
		Name:        input.Name,
		Domain:      input.Domain,
		Nit:         input.Nit,
		Address:     input.Address,
		PhoneNumber: input.PhoneNumber,
		IsDraft:     false,
	}

	updatedCompany, err := cs.CompanyRepository.Update(company)

	if err != nil {
		return models.Company{}, err
	}

	return updatedCompany, nil
}
