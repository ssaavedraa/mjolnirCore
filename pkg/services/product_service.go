package services

import "hex/cms/pkg/models"

type ProductInput struct {
	Name string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price float64 `json:"price" binding:"required,gt=0"`
	ImageUrl string `json:"imageUrl" binding:"required"`
	// TODO; CHANGE THIS VALUES TO BE PROVIDED BY MIDDLEWARE
	UserId uint `json:"userId" binding:"required"`
	CompanyId uint `json:"companyId" binding:"required"`
}

type ProductService interface {
	CreateProduct (input ProductInput) (models.Product, error)
}