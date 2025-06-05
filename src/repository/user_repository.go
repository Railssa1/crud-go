package repository

import (
	"context"
	"os"

	errors_api "github.com/Railssa1/crud-go/src/config/errors"
	"github.com/Railssa1/crud-go/src/config/logger"
	"github.com/Railssa1/crud-go/src/domain"
	"github.com/Railssa1/crud-go/src/repository/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(userDomain domain.UserDomainInterface) (domain.UserDomainInterface, *errors_api.ApiErrors)
}

type userRepository struct {
	dbConnection *mongo.Database
}

func NewUserRepository(dbConnection *mongo.Database) UserRepository {
	return &userRepository{
		dbConnection,
	}
}

func (ur *userRepository) CreateUser(userDomain domain.UserDomainInterface) (domain.UserDomainInterface, *errors_api.ApiErrors) {
	logger.Info("Init userRepository")

	collectionName := os.Getenv("MONGO_COLLECTION")

	collection := ur.dbConnection.Collection(collectionName)

	jsonValue := entity.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), jsonValue)
	if err != nil {
		return nil, errors_api.NewInternalServerError(err.Error())
	}

	jsonValue.ID = (result.InsertedID.(primitive.ObjectID))

	return entity.ConvertEntityToDomain(jsonValue), nil
}
