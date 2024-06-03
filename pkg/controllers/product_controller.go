package controllers

import "github.com/gin-gonic/gin"

type ProductController interface {
	CreateProduct (c *gin.Context)
}