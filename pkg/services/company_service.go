package services

import "hex/mjolnir-core/pkg/models"

type OptionalCompanyInput struct {
	Id          uint    `json:"id"`
	Name        string  `json:"name"`
	Domain      string  `json:"domain"`
	Nit         *string `json:"nit"`
	Address     string  `json:"address"`
	PhoneNumber string  `json:"phoneNumber"`
}

type CompanyService interface {
	UpdateDraftCompany(input OptionalCompanyInput) (models.Company, error)
}
