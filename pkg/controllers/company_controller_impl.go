package controllers

import (
	"hex/mjolnir-core/pkg/services"
	"hex/mjolnir-core/pkg/utils/logging"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompanyControllerImpl struct {
	CompanyService services.CompanyService
}

func NewCompanyController(
	companyService services.CompanyService,
) CompanyController {
	return &CompanyControllerImpl{
		CompanyService: companyService,
	}
}

func (cc *CompanyControllerImpl) UpdateDraftCompany(c *gin.Context) {
	var companyInput services.OptionalCompanyInput

	if err := c.ShouldBindJSON(&companyInput); err != nil {
		logging.Error(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
		})

		return
	}

	_, err := cc.CompanyService.UpdateDraftCompany(companyInput)

	if err != nil {
		logging.Error(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update draft company, please try again later",
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
