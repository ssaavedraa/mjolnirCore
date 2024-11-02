package repositories

import "hex/mjolnir-core/pkg/models"

type TeamMember struct {
	models.User
	RoleName string
}

type TeamRepository interface {
	GetTeams(companyId uint) ([]*models.Team, error)
	GetTeamMembers(companyId uint, teamName string) ([]*TeamMember, error)
}
