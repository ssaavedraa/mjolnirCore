package controllers

import (
	"fmt"
	"hex/mjolnir-core/pkg/services"
	"hex/mjolnir-core/pkg/utils"
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
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid company id", err)

		return
	}

	teams, err := tc.TeamService.GetTeams(uint(companyId))

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to fetch company teams. Please try again later", err)

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

	fmt.Printf("teamNameParam: %v\n", teamNameParam)

	companyId, err := strconv.ParseUint(companyIdParam, 10, 64)

	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid company id", err)

		return
	}

	teamMembers, err := tc.TeamService.GetTeamMembers(uint(companyId), teamNameParam)

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to get team members. Please try again later", err)

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
