package mongodb

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/kevin07696/ecommerce/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ICollectionClient interface {
	FindOneAndDecode(ctx context.Context, filter, model interface{}, opts ...*options.FindOneOptions) error
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
}

type CollectionClient struct {
	collection *mongo.Collection
}

func NewCollectionClient(ctx context.Context, client *mongo.Client, database, collection string, indexModels []mongo.IndexModel) (CollectionClient, error) {
	collectionClient := client.Database(database).Collection(collection)
	_, err := collectionClient.Indexes().CreateMany(ctx, indexModels)
	return CollectionClient{collection: collectionClient}, err
}

func (cc CollectionClient) FindOneAndDecode(ctx context.Context, filter, model interface{}, opts ...*options.FindOneOptions) error {
	err := cc.collection.FindOne(ctx, filter, opts...).Decode(model)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			slog.ErrorContext(ctx, domain.ErrNotFound.Error(), slog.Any("error", err))
			return domain.ErrNotFound
		}
		msg := fmt.Sprintf("%v: For more information on mongo find one check out: https://www.mongodb.com/docs/manual/reference/command/find/: %v", domain.ErrInternalServer, err)
		slog.ErrorContext(ctx, msg, slog.Any("error", err))
		return domain.ErrInternalServer
	}
	return nil
}

func (cc CollectionClient) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return cc.collection.InsertOne(ctx, document, opts...)
}
