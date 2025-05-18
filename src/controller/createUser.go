package controller

import (
	"fmt"

	config "github.com/Railssa1/crud-go/src/config/errors"
	"github.com/Railssa1/crud-go/src/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest models.UserRequest

	if err := c.BindJSON(&userRequest); err != nil {
		restErr := config.NewBadRequestError("There are some incorrect fields")
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println("Usuário criado", userRequest)
}
