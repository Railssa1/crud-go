package service

import (
	errors_api "github.com/Railssa1/crud-go/src/config/errors"
	"github.com/Railssa1/crud-go/src/config/logger"
	"github.com/Railssa1/crud-go/src/domain"
	"github.com/Railssa1/crud-go/src/repository"
)

type UserDomainService interface {
	CreateUser(domain.UserDomainInterface) (domain.UserDomainInterface, *errors_api.ApiErrors)
}

type userDomainService struct {
	repository repository.UserRepository
}

func NewUserDomainService(repository repository.UserRepository) UserDomainService {
	return &userDomainService{
		repository: repository,
	}
}

func (ud *userDomainService) CreateUser(userDomain domain.UserDomainInterface) (domain.UserDomainInterface, *errors_api.ApiErrors) {
	logger.Info("Init CreateUser domain")

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.repository.CreateUser(userDomain)
	if err != nil {
		return nil, err
	}

	return userDomainRepository, nil
}
