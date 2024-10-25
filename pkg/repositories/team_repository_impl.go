package repositories

import (
	"hex/mjolnir-core/pkg/config"
	"hex/mjolnir-core/pkg/models"
	"log"
)

type TeamRepositoryImpl struct{}

func NewTeamRepository() TeamRepository {
	return &TeamRepositoryImpl{}
}

func (repo *TeamRepositoryImpl) GetTeams(companyId uint) ([]models.Team, error) {
	var teams = []models.Team{}
	log.Printf("companyId: %v", companyId)

	result := config.DB.Debug().Where("company_id = ?", companyId).Find(&teams)

	if result.Error != nil {
		return teams, result.Error
	}

	return teams, nil
}

func (repo *TeamRepositoryImpl) GetTeamMembers(companyId uint, teamName string) ([]models.User, error) {
	var teamMembers = []models.User{}

	result := config.DB.Debug().
		Joins("JOIN teams ON teams.id = users.team_id").
		Where("teams.name ILIKE ? AND users.company_id = ?", teamName, companyId).
		Preload("Team").
		Find(&teamMembers)

	if result.Error != nil {
		return teamMembers, result.Error
	}

	return teamMembers, nil
}
