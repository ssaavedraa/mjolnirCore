package services

import (
	"hex/mjolnir-core/pkg/models"
	"hex/mjolnir-core/pkg/repositories"
)

type TeamServiceImpl struct {
	TeamRepository repositories.TeamRepository
}

func NewTeamService(
	teamRepository repositories.TeamRepository,
) TeamService {
	return &TeamServiceImpl{
		TeamRepository: teamRepository,
	}
}

func (ts *TeamServiceImpl) GetTeams(companyId uint) ([]*models.Team, error) {
	teams, err := ts.TeamRepository.GetTeams(companyId)

	if err != nil {
		return nil, err
	}

	return teams, nil
}

func (ts *TeamServiceImpl) GetTeamMembers(companyId uint, teamName string) ([]*repositories.TeamMember, error) {
	teamMembers, err := ts.TeamRepository.GetTeamMembers(companyId, teamName)

	if err != nil {
		return nil, err
	}

	return teamMembers, nil
}
