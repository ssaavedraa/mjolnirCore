package repositories

import "hex/cms/pkg/models"

type ProductRepository interface {
	CreateProduct (product models.Product) (models.Product, error)
	GetAllProducts () ([]models.Product, error)
	GetProductById (id uint) (models.Product, error)
}