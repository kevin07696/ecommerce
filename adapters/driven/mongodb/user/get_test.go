package user_test

import (
	"context"
	"testing"

	"github.com/kevin07696/ecommerce/adapters/driven/mongodb"
	"github.com/kevin07696/ecommerce/adapters/driven/mongodb/mocks"
	"github.com/kevin07696/ecommerce/adapters/driven/mongodb/user"
	"github.com/kevin07696/ecommerce/domain"
	"github.com/kevin07696/ecommerce/domain/auth/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestGetUser(t *testing.T) {
	tc := []struct {
		Name                 string
		Email                models.Email
		Username             models.Username
		FindOneAndDecodeMock func(ctx context.Context, filter, model interface{}, opts ...*options.FindOneOptions) error
		ExpectedError        error
		ExpectedResponse     models.User
	}{
		{
			Name: "Succeeds_ByEmail",
			Email: models.Email{
				Local: "local",
				SubAddress: "+sub",
				Domain: "domain.sub.tld",
			},
			FindOneAndDecodeMock: func(ctx context.Context, filter, model interface{}, opts ...*options.FindOneOptions) error {
				model = models.User{
					Username: "user",
					Role: "developer",
					Email: models.Email{
						Local: "local",
						SubAddress: "+sub",
						Domain: "domain.sub.tld",
					},
				}
				return nil
			},
		},
		{
			Name: "Succeeds_ByUsername",
			Username: "user",
			FindOneAndDecodeMock: func(ctx context.Context, filter, model interface{}, opts ...*options.FindOneOptions) error {
				model = models.User{
					Username: "user",
					Role: "developer",
					Email: models.Email{
						Local: "local",
						SubAddress: "+sub",
						Domain: "domain.sub.tld",
					},
				}
				return nil
			},
		},
		{
			Name: "Fails_NotFound",
			Username: "user",
			Email: models.Email{
				Local: "local",
				SubAddress: "+sub",
				Domain: "domain.sub.tld",
			},
			FindOneAndDecodeMock: func(ctx context.Context, filter, model interface{}, opts ...*options.FindOneOptions) error {
				return domain.ErrNotFound
			},
			ExpectedError: domain.ErrNotFound,
		},
		{
			Name: "Fails_InternalServer",
			Username: "user",
			Email: models.Email{
				Local: "local",
				SubAddress: "+sub",
				Domain: "domain.sub.tld",
			},
			FindOneAndDecodeMock: func(ctx context.Context, filter, model interface{}, opts ...*options.FindOneOptions) error {
				return domain.ErrInternalServer
			},
			ExpectedError: domain.ErrInternalServer,
		},
	}

	t.Run(tc[0].Name, func(t *testing.T) {
		var collection mongodb.ICollectionClient = mocks.Collection{
			FindOneAndDecodeMock: tc[0].FindOneAndDecodeMock,
		}

		repo := user.UserRepository{
			Collection: collection,
		}

		response, _ := repo.GetUserByEmail(context.TODO(), tc[0].Email)

		assert.Equal(t, response, tc[0].ExpectedResponse)
	})

	t.Run(tc[1].Name, func(t *testing.T) {
		var collection mongodb.ICollectionClient = mocks.Collection{
			FindOneAndDecodeMock: tc[1].FindOneAndDecodeMock,
		}

		repo := user.UserRepository{
			Collection: collection,
		}

		response, _ := repo.GetUserByUsername(context.TODO(), tc[1].Username)

		assert.Equal(t, response, tc[1].ExpectedResponse)
	})

	for i := 2; i < len(tc); i++ {
		t.Run(tc[i].Name, func(t *testing.T) {
			var collection mongodb.ICollectionClient = mocks.Collection{
				FindOneAndDecodeMock: tc[i].FindOneAndDecodeMock,
			}
	
			repo := user.UserRepository{
				Collection: collection,
			}
	
			response, err := repo.GetUserByUsername(context.TODO(), tc[i].Username)
			
			assert.Equal(t, response, tc[i].ExpectedResponse)
			assert.Equal(t, err, tc[i].ExpectedError)
			
			response, err = repo.GetUserByEmail(context.TODO(), tc[i].Email)

			assert.Equal(t, response, tc[i].ExpectedResponse)
			assert.Equal(t, err, tc[i].ExpectedError)
		})
	}
}
