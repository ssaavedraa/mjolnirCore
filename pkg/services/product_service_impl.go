package services

import (
	"hex/cms/pkg/models"
	"hex/cms/pkg/repositories"
)

type ProductServiceImpl struct {
	ProductRepository repositories.ProductRepository
}

func NewProductService (
	productRepository repositories.ProductRepository,
) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
	}
}

func (ps *ProductServiceImpl) CreateProduct (input ProductInput) (models.Product, error) {
	product := models.Product{
		Name: input.Name,
		Description: input.Description,
		Price: input.Price,
		ImageUrl: input.ImageUrl,
		CreatedBy: input.UserId,
		CompanyID: input.CompanyId,
	}

	createdProduct, err := ps.ProductRepository.CreateProduct(product)

	if err != nil {
		return models.Product{}, err
	}

	return createdProduct, nil
}