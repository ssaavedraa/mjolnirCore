package controllers

import "github.com/gin-gonic/gin"

type UserController interface {
	CreateUser(c *gin.Context)
	Login(c *gin.Context)
	GetByInviteId(c *gin.Context)
	UpdateUser(c *gin.Context)
	GetUserById(c *gin.Context)
}
