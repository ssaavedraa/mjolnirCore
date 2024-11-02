package repositories

import (
	"hex/mjolnir-core/pkg/models"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		db: db,
	}
}

func (repo *ProductRepositoryImpl) CreateProduct(product *models.Product) (*models.Product, error) {
	if err := repo.db.Create(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (repo *ProductRepositoryImpl) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := repo.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (repo *ProductRepositoryImpl) GetProductById(id uint) (*models.Product, error) {
	product := models.Product{
		Model: gorm.Model{
			ID: id,
		},
	}

	if err := repo.db.Find(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}
