package controllers

import (
	"hex/mjolnir-core/pkg/services"
	"hex/mjolnir-core/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductControllerImpl struct {
	ProductService services.ProductService
}

func NewProductController(productService services.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (pc *ProductControllerImpl) CreateProduct(c *gin.Context) {
	// userId, _ := c.Get("userId")
	// companyId, _ := c.Get("companyId")
	var productInput services.ProductInput

	if err := c.ShouldBindJSON(&productInput); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request Payload", err)
		// logging.Error(err)

		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"message": "Invalid request payload",
		// })

		return
	}

	createdProduct, err := pc.ProductService.CreateProduct(productInput)

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create product. Please try again later", err)
		// logging.Error(err)

		// c.JSON(http.StatusInternalServerError, gin.H{
		// 	"message": "Failed to create product. Please try again later",
		// })

		return
	}

	response := utils.ConvertToResponse(createdProduct, utils.ResponseFields{
		"id":          createdProduct.ID,
		"name":        createdProduct.Name,
		"description": createdProduct.Description,
		"price":       createdProduct.Price,
		"imageUrl":    createdProduct.ImageUrl,
	})

	c.JSON(http.StatusCreated, response)
}

func (pc *ProductControllerImpl) GetAllProducts(c *gin.Context) {
	products, err := pc.ProductService.GetAllProducts()

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to get products. Please try again later", err)

		return
	}

	var productsResponse []any

	for _, product := range products {
		formattedProductResponse := utils.ConvertToResponse(product, utils.ResponseFields{
			"id":          product.ID,
			"name":        product.Name,
			"description": product.Description,
			"price":       product.Price,
			"imageUrl":    product.ImageUrl,
		})
		productsResponse = append(productsResponse, formattedProductResponse)
	}

	c.JSON(http.StatusCreated, productsResponse)
}

func (pc *ProductControllerImpl) GetProductById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid product id", err)

		return
	}

	productId := uint(id)

	product, err := pc.ProductService.GetProductById(uint(productId))

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to get product. Please try again later", err)

		return
	}

	response := utils.ConvertToResponse(product, utils.ResponseFields{
		"id":          product.ID,
		"name":        product.Name,
		"description": product.Description,
		"price":       product.Price,
		"imageUrl":    product.ImageUrl,
	})

	c.JSON(http.StatusCreated, response)
}
