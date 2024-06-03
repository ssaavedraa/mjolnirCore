package controllers

import (
	"hex/cms/pkg/services"
	"hex/cms/pkg/utils"
	"hex/cms/pkg/utils/logging"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductControllerImpl struct {
	ProductService services.ProductService
}

func NewProductController (productService services.ProductService) ProductController {
	return &ProductControllerImpl {
		ProductService: productService,
	}
}

func (pc *ProductControllerImpl) CreateProduct (c *gin.Context) {
	// userId, _ := c.Get("userId")
	// companyId, _ := c.Get("companyId")
	var productInput services.ProductInput

	if err := c.ShouldBindJSON(&productInput); err != nil {
		logging.Error(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
		})

		return
	}

	createdProduct, err := pc.ProductService.CreateProduct(productInput)

	if err != nil {
		logging.Error(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create product. Please try again later",
		})

		return
	}

	response := utils.ConvertToResponse(createdProduct, utils.ResponseFields{
		"id": createdProduct.ID,
		"name": createdProduct.Name,
		"description": createdProduct.Description,
		"price": createdProduct.Price,
		"imageUrl": createdProduct.ImageUrl,
	})

	c.JSON(http.StatusCreated, response)
}