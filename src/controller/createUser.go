package controller

import (
	"net/http"

	"github.com/Railssa1/crud-go/src/config/logger"
	"github.com/Railssa1/crud-go/src/config/validation"
	"github.com/Railssa1/crud-go/src/domain"
	"github.com/Railssa1/crud-go/src/dto"
	"github.com/Railssa1/crud-go/src/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller")
	var userRequest models.UserRequest

	if err := c.BindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	userDomain := domain.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	if err := uc.service.CreateUser(userDomain); err != nil {
		c.JSON(err.Code, err.Message)
		return
	}

	userResponse := dto.ConvertDomainToResponse(userDomain)
	logger.Info("User created successfully", zap.Any("userResponse", userResponse))

	c.JSON(http.StatusOK, userResponse)
}
