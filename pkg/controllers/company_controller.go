package controllers

import "github.com/gin-gonic/gin"

type CompanyController interface {
	UpdateCompany(c *gin.Context)
}
