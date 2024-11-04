package services

import (
	"hex/mjolnir-core/pkg/models"
	"hex/mjolnir-core/pkg/repositories"

	"gorm.io/gorm"
)

type CompanyServiceImpl struct {
	CompanyRepository repositories.CompanyRepository
	RoleRepository    repositories.RoleRepository
}

func NewCompanyService(
	companyRepository repositories.CompanyRepository,
	roleRepository repositories.RoleRepository,
) CompanyService {
	return &CompanyServiceImpl{
		CompanyRepository: companyRepository,
		RoleRepository:    roleRepository,
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

func (cs *CompanyServiceImpl) GetCompanyRoles(companyId uint) ([]models.Role, error) {
	companyRoles, err := cs.RoleRepository.GetCompanyRoles(companyId)

	if err != nil {
		return nil, err
	}

	return companyRoles, nil
}

func (cs *CompanyServiceImpl) CreateCompanyRole(companyId uint, roleName string) error {
	_, err := cs.RoleRepository.FindOrCreateRoleByName(roleName, companyId)

	if err != nil {
		return err
	}

	return nil
}
