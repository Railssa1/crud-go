package entity

import (
	"github.com/Railssa1/crud-go/src/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserEntity struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Name     string             `bson:"name"`
	Age      int8               `bson:"age"`
}

func ConvertDomainToEntity(userDomain domain.UserDomainInterface) UserEntity {
	return UserEntity{
		Email:    userDomain.GetEmail(),
		Password: userDomain.GetPassword(),
		Name:     userDomain.GetName(),
		Age:      userDomain.GetAge(),
	}
}

func ConvertEntityToDomain(user UserEntity) domain.UserDomainInterface {
	userDomain := domain.NewUserDomain(user.Email, user.Password, user.Name, user.Age)
	userDomain.SetId(user.ID.Hex())
	return userDomain
}
