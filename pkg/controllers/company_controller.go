package controllers

import "github.com/gin-gonic/gin"

type CompanyController interface {
	UpdateDraftCompany(c *gin.Context)
}
