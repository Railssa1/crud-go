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

	response, err := uc.service.CreateUser(userDomain)
	if err != nil {
		c.JSON(err.Code, err.Message)
		return
	}

	logger.Info("User created successfully", zap.Any("userResponse", response))

	c.JSON(http.StatusOK, dto.ConvertDomainToResponse(response))
}
