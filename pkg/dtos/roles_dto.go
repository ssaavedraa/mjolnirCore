package dtos

type RoleResponse struct {
	Name *string `json:"name,omitempty"`

	Company *CompanyResponse `json:"company,omitempty"`
	Users   []UserResponse   `json:"users,omitempty"`
}
