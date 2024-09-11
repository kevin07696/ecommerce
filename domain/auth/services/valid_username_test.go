package services_test

import (
	"context"
	"testing"

	"github.com/kevin07696/ecommerce/domain"
	"github.com/kevin07696/ecommerce/domain/auth/mocks"
	"github.com/kevin07696/ecommerce/domain/auth/models"
	"github.com/kevin07696/ecommerce/domain/auth/port"
	"github.com/kevin07696/ecommerce/domain/auth/services"
	"github.com/stretchr/testify/assert"
)

func TestValidateUsername(t *testing.T) {
	tc := []struct {
		Name                  string
		Request               services.ValidateUsernameReq
		NewUsernameMock       func(username string) (models.Username, error)
		GetUserByUsernameMock func(ctx context.Context, username models.Username) (models.User, error)
		ExpectedError         *domain.CustomError
	}{
		{
			Name:    "Succeeds",
			Request: services.ValidateUsernameReq{Username: "user"},
			NewUsernameMock: func(username string) (models.Username, error) {
				return models.Username("user"), nil
			},
			GetUserByUsernameMock: func(ctx context.Context, username models.Username) (models.User, error) {
				return models.User{}, domain.ErrNotFound
			},
		},
		{
			Name:    "FailsValidation_EmptyUsername",
			Request: services.ValidateUsernameReq{Username: ""},
			NewUsernameMock: func(username string) (models.Username, error) {
				return models.Username(""), models.ErrEmptyUsername
			},
			ExpectedError: domain.ValidationError(services.ErrMsgEmptyUsername),
		},
		{
			Name:    "FailsValidation_InvalidUsername",
			Request: services.ValidateUsernameReq{Username: "$user"},
			NewUsernameMock: func(username string) (models.Username, error) {
				return models.Username(""), models.ErrInvalidUsername
			},
			ExpectedError: domain.ValidationError(services.ErrMsgInvalidUsername),
		},
		{
			Name: "FailsRepo_DuplicateKey",
			Request: services.ValidateUsernameReq{Username: "user"},
			NewUsernameMock: func(username string) (models.Username, error) {
				return models.Username("user"), nil
			},
			GetUserByUsernameMock: func(ctx context.Context, username models.Username) (models.User, error) {
				return models.User{
					Username: models.Username("user"),
					Email: models.Email{
						Local: "local",
						SubAddress: "+sub",
						Domain: "domain.sub.tld",
					},
					Role: "developer",
				}, nil
			},
			ExpectedError: domain.CustomizeError(domain.ErrDuplicateKey, services.ErrMsgUsernameNotUnique),
		},
		{
			Name:    "FailsRepo_InternalServer",
			Request: services.ValidateUsernameReq{Username: "user"},
			NewUsernameMock: func(username string) (models.Username, error) {
				return models.Username("user"), nil
			},
			GetUserByUsernameMock: func(ctx context.Context, username models.Username) (models.User, error) {
				return models.User{}, domain.ErrInternalServer
			},
			ExpectedError: domain.CustomizeError(domain.ErrInternalServer, services.ErrMsgInternalServer),
		},
	}
	for i := range tc {
		t.Run(tc[i].Name, func(t *testing.T) {
			var sessionManager port.ISessionManager
			var cacher port.ICache
			var emailer port.IEmail
			var repositor port.IRepository = &mocks.MockRepository{
				GetUserByUsernameMock: tc[i].GetUserByUsernameMock,
			}
			var modeler models.IModels = &mocks.MockModels{
				NewUsernameMock: tc[i].NewUsernameMock,
			}
			service := services.NewService(repositor, sessionManager, cacher, emailer, modeler)

			err := service.ValidateUsername(context.TODO(), tc[i].Request)

			assert.Equal(t, tc[i].ExpectedError, err)
		})
	}
}
