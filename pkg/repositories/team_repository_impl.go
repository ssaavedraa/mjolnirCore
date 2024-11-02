package repositories

import (
	"hex/mjolnir-core/pkg/models"

	"gorm.io/gorm"
)

type TeamRepositoryImpl struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) TeamRepository {
	return &TeamRepositoryImpl{
		db: db,
	}
}

func (repo *TeamRepositoryImpl) GetTeams(companyId uint) ([]*models.Team, error) {
	var teams = []*models.Team{}

	if err := repo.db.
		Where("company_id = ?", companyId).
		Find(&teams).Error; err != nil {
		return nil, err
	}

	return teams, nil
}

func (repo *TeamRepositoryImpl) GetTeamMembers(companyId uint, teamName string) ([]*TeamMember, error) {
	var teamMembers = []*TeamMember{}

	if err := repo.db.
		Model(&models.User{}).
		Select("users.*, company_roles.role_id, roles.name as role_name").
		Joins("JOIN teams ON teams.id = users.team_id").
		Joins("JOIN company_roles ON company_roles.id = users.company_role_id").
		Joins("JOIN roles ON roles.id = company_roles.role_id").
		Where("teams.name ILIKE ? AND users.company_id = ?", teamName, companyId).
		Preload("Team").
		Find(&teamMembers).Error; err != nil {
		return nil, err
	}

	return teamMembers, nil
}
