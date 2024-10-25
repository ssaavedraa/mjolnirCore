package services

import "hex/mjolnir-core/pkg/models"

type TeamService interface {
	GetTeams(companyId uint) ([]models.Team, error)
	GetTeamMembers(companyId uint, teamName string) ([]models.User, error)
}
