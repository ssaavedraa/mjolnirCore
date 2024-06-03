package repositories

import (
	"hex/cms/pkg/config"
	"hex/cms/pkg/models"
)

type ProductRepositoryImpl struct {}

func NewProductRepository () ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repo *ProductRepositoryImpl) CreateProduct (product models.Product) (models.Product, error) {
	result := config.DB.Create(&product)

	if result.Error != nil {
		return models.Product{}, result.Error
	}

	return product, nil
}