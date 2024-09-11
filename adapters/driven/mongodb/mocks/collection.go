package mocks

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection struct {
	InsertOneMock        func(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	FindOneAndDecodeMock func(ctx context.Context, filter, model interface{}, opts ...*options.FindOneOptions) error
}

func (c Collection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return c.InsertOneMock(ctx, document, opts...)
}

func (c Collection) FindOneAndDecode(ctx context.Context, filter, model interface{}, opts ...*options.FindOneOptions) error {
	return c.FindOneAndDecodeMock(ctx, filter, model, opts...)
}
