package service

import (
	"fmt"

	errors_api "github.com/Railssa1/crud-go/src/config/errors"
	"github.com/Railssa1/crud-go/src/config/logger"
	"github.com/Railssa1/crud-go/src/domain"
)

type UserDomainService interface {
	CreateUser(domain.UserDomainInterface) *errors_api.ApiErrors
}

type userDomainService struct{}

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

func (ud *userDomainService) CreateUser(userDomain domain.UserDomainInterface) *errors_api.ApiErrors {
	logger.Info("Init CreateUser domain")

	userDomain.EncryptPassword()

	fmt.Println(userDomain)

	return nil
}
