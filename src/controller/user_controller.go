package controller

import (
	"github.com/Railssa1/crud-go/src/service"
	"github.com/gin-gonic/gin"
)

func NewUserController(service service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service,
	}
}

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
