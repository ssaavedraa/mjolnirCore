package controllers

import (
	"hex/mjolnir-core/pkg/services"
	"hex/mjolnir-core/pkg/utils/logging"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TeamcontrollerImpl struct {
	TeamService services.TeamService
}

func NewTeamController(teamService services.TeamService) TeamController {
	return &TeamcontrollerImpl{
		TeamService: teamService,
	}
}

func (tc *TeamcontrollerImpl) GetTeams(c *gin.Context) {
	companyIdParam := c.Param("companyId")

	companyId, err := strconv.ParseUint(companyIdParam, 10, 64)

	if err != nil {
		logging.Error(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid company id",
		})

		return
	}

	teams, err := tc.TeamService.GetTeams(uint(companyId))

	if err != nil {
		logging.Error(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch company teams. Please try again later",
		})

		return
	}

	var response []map[string]interface{}

	for _, team := range teams {
		response = append(response, map[string]interface{}{
			"id":   team.ID,
			"name": team.Name,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (tc *TeamcontrollerImpl) GetTeamMembers(c *gin.Context) {
	companyIdParam := c.Param("companyId")
	teamNameParam := c.Param("teamName")

	companyId, err := strconv.ParseUint(companyIdParam, 10, 64)

	if err != nil {
		logging.Error(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid company id",
		})

		return
	}

	teamMembers, err := tc.TeamService.GetTeamMembers(uint(companyId), teamNameParam)

	if err != nil {
		logging.Error(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get team members. Please try again later",
		})

		return
	}

	var response []map[string]interface{}

	for _, teamMember := range teamMembers {
		response = append(response, map[string]interface{}{
			"id":       teamMember.ID,
			"name":     teamMember.Fullname,
			"email":    teamMember.Email,
			"role":     teamMember.CompanyRole,
			"teamName": teamMember.Team.Name,
		})
	}

	c.JSON(http.StatusOK, response)
}
