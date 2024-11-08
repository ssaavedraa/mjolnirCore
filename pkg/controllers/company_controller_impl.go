package controllers

import (
	"fmt"
	"hex/mjolnir-core/pkg/services"
	"hex/mjolnir-core/pkg/utils"
	"net/http"
	"strconv"

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

func (cc *CompanyControllerImpl) GetCompanyRoles(c *gin.Context) {
	companyIdParam := c.Param("companyId")

	companyId, err := strconv.ParseUint(companyIdParam, 10, 64)

	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid company id", err)

		return
	}

	companyRoles, err := cc.CompanyService.GetCompanyRoles(uint(companyId))

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to get company roles, Please try again later", err)

		return
	}

	var response []map[string]interface{}

	for _, role := range companyRoles {
		response = append(response, map[string]interface{}{
			"id":   role.ID,
			"name": role.Name,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (cc *CompanyControllerImpl) CreateCompanyRole(c *gin.Context) {
	companyIdParam := c.Param("companyId")

	companyId, err := strconv.ParseUint(companyIdParam, 10, 64)
	fmt.Printf("companyId: %d", companyId)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid company id", err)

		return
	}

	var companyRoleInput services.CompanyRoleInput

	if err = c.ShouldBindJSON(&companyRoleInput); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request Payload", err)

		return
	}

	if err = cc.CompanyService.CreateCompanyRole(uint(companyId), companyRoleInput.RoleName); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create comapny role, Please try again later", err)

		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
