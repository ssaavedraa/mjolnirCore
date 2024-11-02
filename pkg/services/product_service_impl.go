package services

import (
	"hex/mjolnir-core/pkg/models"
	"hex/mjolnir-core/pkg/repositories"
)

type ProductServiceImpl struct {
	ProductRepository repositories.ProductRepository
}

func NewProductService(
	productRepository repositories.ProductRepository,
) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
	}
}

func (ps *ProductServiceImpl) CreateProduct(input ProductInput) (*models.Product, error) {
	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		ImageUrl:    input.ImageUrl,
		CreatedBy:   input.UserId,
		CompanyID:   input.CompanyId,
	}

	createdProduct, err := ps.ProductRepository.CreateProduct(&product)

	if err != nil {
		return nil, err
	}

	return createdProduct, nil
}

func (ps *ProductServiceImpl) GetAllProducts() ([]models.Product, error) {
	products, err := ps.ProductRepository.GetAllProducts()

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (ps *ProductServiceImpl) GetProductById(id uint) (*models.Product, error) {
	product, err := ps.ProductRepository.GetProductById(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}
