package dto

import (
	"github.com/Railssa1/crud-go/src/domain"
	"github.com/Railssa1/crud-go/src/models"
)

func ConvertDomainToResponse(userDomain domain.UserDomainInterface) models.UserResponse {
	return models.UserResponse{
		Id:    1,
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
