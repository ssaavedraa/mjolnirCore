package repositories

import "hex/mjolnir-core/pkg/models"

type TeamRepository interface {
	GetTeams(companyId uint) ([]models.Team, error)
	GetTeamMembers(companyId uint, teamName string) ([]models.User, error)
}
