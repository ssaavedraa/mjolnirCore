package controllers

import "github.com/gin-gonic/gin"

type TeamController interface {
	GetTeams(c *gin.Context)
	GetTeamMembers(c *gin.Context)
}
