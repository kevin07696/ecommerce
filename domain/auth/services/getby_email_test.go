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

func TestGetByEmail(t *testing.T) {
	tc := []struct {
		Name               string
		Request            services.GetUserByEmailReq
		NewEmailMock       func(email string) (models.Email, error)
		GetUserByEmailMock func(ctx context.Context, email models.Email) (models.User, error)
		ExpectedError      *domain.CustomError
		ExpectedResponse   services.GetUserByEmailResp
	}{
		{
			Name:    "Succeeds",
			Request: services.GetUserByEmailReq{Email: "local+sub@domain.sub.tld"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{
					Local:      "local",
					SubAddress: "+sub",
					Domain:     "domain.sub.tld",
				}, nil
			},
			GetUserByEmailMock: func(ctx context.Context, email models.Email) (models.User, error) {
				return models.User{
					Username: models.Username("user"),
					Email: models.Email{
						Local:      "local",
						SubAddress: "+sub",
						Domain:     "domain.sub.tld",
					},
					Role: models.Role("developer"),
				}, nil
			},
			ExpectedResponse: services.GetUserByEmailResp{Username: "user"},
		},
		{
			Name:    "FailsValidation_EmptyEmail",
			Request: services.GetUserByEmailReq{Email: ""},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{}, models.ErrEmptyEmail
			},
			ExpectedError: domain.ValidationError(services.ErrMsgEmptyEmail),
		},
		{
			Name:    "FailsValidation_InvalidEmail",
			Request: services.GetUserByEmailReq{Email: "local+$ub@domain.sub.tld"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{}, models.ErrInvalidEmail
			},
			ExpectedError: domain.ValidationError(services.ErrMsgInvalidEmail),
		},
		{
			Name:    "FailsRepo_NotFound",
			Request: services.GetUserByEmailReq{Email: "local+sub@domain.sub.tld"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{
					Local:      "local",
					SubAddress: "+sub",
					Domain:     "domain.sub.tld",
				}, nil
			},
			GetUserByEmailMock: func(ctx context.Context, email models.Email) (models.User, error) {
				return models.User{}, domain.ErrNotFound
			},
			ExpectedError: domain.CustomizeError(domain.ErrNotFound, services.ErrMsgEmailNotFound),
		},
		{
			Name:    "FailsRepo_InternalServer",
			Request: services.GetUserByEmailReq{Email: "local+sub@domain.sub.tld"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{
					Local:      "local",
					SubAddress: "+sub",
					Domain:     "domain.sub.tld",
				}, nil
			},
			GetUserByEmailMock: func(ctx context.Context, email models.Email) (models.User, error) {
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
			var Repositor port.IRepository = &mocks.MockRepository{
				GetUserByEmailMock: tc[i].GetUserByEmailMock,
			}
			var modeler models.IModels = &mocks.MockModels{
				NewEmailMock: tc[i].NewEmailMock,
			}

			service := services.NewService(Repositor, sessionManager, cacher, emailer, modeler)
			resp, err := service.GetUserByEmail(context.TODO(), tc[i].Request)

			assert.Equal(t, tc[i].ExpectedError, err)
			assert.Equal(t, tc[i].ExpectedResponse, resp)
		})
	}
}
