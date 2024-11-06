package dtos

type CompanyResponse struct {
	ID          *uint   `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Domain      *string `json:"domain,omitempty"`
	Nit         *string `json:"nit,omitempty"`
	Address     *string `json:"address,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}
