package controllers

import "github.com/gin-gonic/gin"

type CompanyController interface {
	UpdateCompany(c *gin.Context)
	GetCompanyRoles(c *gin.Context)
	CreateCompanyRole(c *gin.Context)
}
