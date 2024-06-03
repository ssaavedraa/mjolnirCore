package repositories

import (
	"hex/cms/pkg/config"
	"hex/cms/pkg/models"

	"gorm.io/gorm"
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

func (repo *ProductRepositoryImpl) GetAllProducts () ([]models.Product, error) {
	var products []models.Product
	result := config.DB.Find(&products)

	if result.Error != nil {
		return []models.Product{}, result.Error
	}

	return products, nil
}

func (repo *ProductRepositoryImpl) GetProductById (id uint) (models.Product, error) {
	product := models.Product{
		Model: gorm.Model{
			ID: id,
		},
	}

	result := config.DB.Find(&product)

	if result.Error != nil {
		return models.Product{}, result.Error
	}

	return product, nil
}