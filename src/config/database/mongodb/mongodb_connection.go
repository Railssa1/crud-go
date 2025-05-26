package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitConnection() (*mongo.Database, error) {
	ctx := context.Background()
	dbConnection := os.Getenv("MONGODB_URL")
	dbName := os.Getenv("MONGODB_DATABASE")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbConnection))

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(dbName), nil
}
