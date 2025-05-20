package controller

import (
	"fmt"

	"github.com/Railssa1/crud-go/src/config/validation"
	"github.com/Railssa1/crud-go/src/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest models.UserRequest

	if err := c.BindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println("Usuário criado", userRequest)
}
