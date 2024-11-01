package controllers

import (
	"hex/mjolnir-core/pkg/services"
	"hex/mjolnir-core/pkg/utils"
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

func (cc *CompanyControllerImpl) UpdateCompany(c *gin.Context) {
	var companyInput services.OptionalCompanyInput

	if err := c.ShouldBindJSON(&companyInput); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request Payload", err)

		return
	}

	_, err := cc.CompanyService.UpdateCompany(companyInput)

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to update company, please try again later", err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
