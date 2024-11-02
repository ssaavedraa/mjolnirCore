package repositories

import (
	"fmt"
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

	result := config.DB.Where("company_id = ?", companyId).Find(&teams)

	if result.Error != nil {
		return teams, result.Error
	}

	return teams, nil
}

func (repo *TeamRepositoryImpl) GetTeamMembers(companyId uint, teamName string) ([]TeamMember, error) {
	var teamMembers = []TeamMember{}

	result := config.DB.Debug().
		Model(&models.User{}).
		Select("users.*, company_roles.role_id, roles.name as role_name").
		Joins("JOIN teams ON teams.id = users.team_id").
		Joins("JOIN company_roles ON company_roles.id = users.company_role_id").
		Joins("JOIN roles ON roles.id = company_roles.role_id").
		Where("teams.name ILIKE ? AND users.company_id = ?", teamName, companyId).
		Preload("Team").
		Find(&teamMembers)

	if result.Error != nil {
		return teamMembers, result.Error
	}

	fmt.Printf("team members: %v\n", teamMembers)

	return teamMembers, nil
}
