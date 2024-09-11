package user

import (
	"context"

	"github.com/kevin07696/ecommerce/adapters/driven/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	Collection mongodb.ICollectionClient
}

func NewUserRepository(ctx context.Context, client *mongo.Client, database string, collection string) (UserRepository, error) {
	indexModels := []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "local", Value: -1},
				{Key: "domain", Value: -1},
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "username", Value: -1}},
			Options: options.Index().SetUnique(true),
		},
	}
	collectionClient, err := mongodb.NewCollectionClient(ctx, client, database, collection, indexModels)
	if err != nil {
		return UserRepository{}, err
	}

	return UserRepository{
		Collection: collectionClient,
	}, nil
}
