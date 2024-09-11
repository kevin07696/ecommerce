package user_test

import (
	"context"
	"errors"
	"testing"

	"github.com/kevin07696/ecommerce/adapters/driven/mongodb/mocks"
	"github.com/kevin07696/ecommerce/adapters/driven/mongodb/user"
	"github.com/kevin07696/ecommerce/domain"
	"github.com/kevin07696/ecommerce/domain/auth/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCreateUser(t *testing.T) {
	tc := []struct {
		Name          string
		InsertOneMock func(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
		ExpectedError error
	}{
		{
			Name: "Succeeds",
			InsertOneMock: func(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
				return nil, nil
			},
		},
		{
			Name: "Fails_DuplicateEmail",
			InsertOneMock: func(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
				return nil, mocks.MockServerError{
					Name: "E11000",
					Code: 16460,
					Message: "MongoError: E11000 duplicate key error collection: db.users index: local_-1_domain_-1 dup key: { local: null, domain: null }",
				}
			},
			ExpectedError: domain.ErrDuplicateEmail,
		},
		{
			Name: "Fails_DuplicateUsername",
			InsertOneMock: func(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
				return nil, mocks.MockServerError{
					Name: "E11000",
					Code: 16460,
					Message: "Error: E11000 db.users index: username_-1 dup key: { username: null }",
				}
			},
			ExpectedError: domain.ErrDuplicateUsername,
		},
		{
			Name: "Fails_UnknownDuplicateKey",
			InsertOneMock: func(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
				return nil, mocks.MockServerError{
					Name: "E11000",
					Code: 16460,
					Message: "MongoError: E11000 duplicate key error collection: db.users index: index_-1 dup key: { : null }",
				}
			},
			ExpectedError: domain.ErrInternalServer,
		},
		{
			Name: "Fails_InternalServer",
			InsertOneMock: func(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
				return nil, errors.New("context deadline exceeded")
			},
			ExpectedError: domain.ErrInternalServer,
		},
	}

	for i := range tc {
		t.Run(tc[i].Name, func(t *testing.T) {
			collection := mocks.Collection{
				InsertOneMock: tc[i].InsertOneMock,
			}
			repo := user.UserRepository{
				Collection: collection,
			}
			user := models.User{
				Email: models.Email{
					Local: "local",
					SubAddress: "+sub",
					Domain: "domain.sub.tld",
				},
				Username: "user",
				Role: "developer",
			}
			err := repo.CreateUser(context.TODO(), user)

			assert.Equal(t, tc[i].ExpectedError, err)
		})
	}
}
