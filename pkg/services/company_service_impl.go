package services

import (
	"hex/mjolnir-core/pkg/models"
	"hex/mjolnir-core/pkg/repositories"

	"gorm.io/gorm"
)

type CompanyServiceImpl struct {
	CompanyRepository     repositories.CompanyRepository
	CompanyRoleRepository repositories.CompanyRoleRepository
}

func NewCompanyService(
	companyRepository repositories.CompanyRepository,
	companyRoleRepository repositories.CompanyRoleRepository,
) CompanyService {
	return &CompanyServiceImpl{
		CompanyRepository:     companyRepository,
		CompanyRoleRepository: companyRoleRepository,
	}
}

func (cs *CompanyServiceImpl) UpdateCompany(input OptionalCompanyInput) (*models.Company, error) {
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

	updatedCompany, err := cs.CompanyRepository.Update(&company)

	if err != nil {
		return nil, err
	}

	return updatedCompany, nil
}

func (cs *CompanyServiceImpl) GetCompanyRoles(companyId uint) ([]models.CompanyRole, error) {
	companyRoles, err := cs.CompanyRoleRepository.GetCompanyRoles(companyId)

	if err != nil {
		return []models.CompanyRole{}, err
	}

	return companyRoles, nil
}
