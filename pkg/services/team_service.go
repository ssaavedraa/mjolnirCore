package services

import (
	"hex/mjolnir-core/pkg/models"
	"hex/mjolnir-core/pkg/repositories"
)

type TeamService interface {
	GetTeams(companyId uint) ([]models.Team, error)
	GetTeamMembers(companyId uint, teamName string) ([]repositories.TeamMember, error)
}
