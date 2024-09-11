package services_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/kevin07696/ecommerce/domain"
	"github.com/kevin07696/ecommerce/domain/auth/mocks"
	"github.com/kevin07696/ecommerce/domain/auth/models"
	"github.com/kevin07696/ecommerce/domain/auth/port"
	"github.com/kevin07696/ecommerce/domain/auth/services"
	"github.com/stretchr/testify/assert"
)

func TestSendEmailOTP(t *testing.T) {
	tc := []struct {
		Name               string
		Request            services.SendRegisterOTPReq
		NewEmailMock       func(email string) (models.Email, error)
		GetUserByEmailMock func(ctx context.Context, email models.Email) (models.User, error)
		SendEmailMock      func(ctx context.Context, to, subject, body string) error
		SetCacheMock       func(ctx context.Context, key, value string, exp time.Duration) error
		ExpectedError      *domain.CustomError
	}{
		{
			Name:    "Succeeds",
			Request: services.SendRegisterOTPReq{Email: "local+sub@domain.sub.tld"},
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
			SendEmailMock: func(ctx context.Context, to, subject, body string) error {
				return nil
			},
			SetCacheMock: func(ctx context.Context, key, value string, exp time.Duration) error {
				return nil
			},
		},
		{
			Name:    "FailsValidation_EmptyEmail",
			Request: services.SendRegisterOTPReq{Email: ""},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{}, models.ErrEmptyEmail
			},
			ExpectedError: domain.ValidationError(services.ErrMsgEmptyEmail),
		},
		{
			Name:    "FailsValidation_InvalidEmail",
			Request: services.SendRegisterOTPReq{Email: "local+$ub@domain.sub.tld"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{}, models.ErrInvalidEmail
			},
			ExpectedError: domain.ValidationError(services.ErrMsgInvalidEmail),
		},
		{
			Name:    "FailsRepo_DuplicateKey",
			Request: services.SendRegisterOTPReq{Email: "local+sub@domain.sub.tld"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{
					Local:      "local",
					SubAddress: "+sub",
					Domain:     "domain.sub.tld",
				}, nil
			},
			GetUserByEmailMock: func(ctx context.Context, email models.Email) (models.User, error) {
				return models.User{
					Email: models.Email{
						Local:      "local",
						SubAddress: "+sub",
						Domain:     "domain.sub.tld",
					},
					Username: models.Username("user"),
					Role:     models.Role("developer"),
				}, nil
			},
			ExpectedError: domain.CustomizeError(domain.ErrDuplicateKey, services.ErrMsgEmailNotUnique),
		},
		{
			Name:    "FailsRepo_InternalServer",
			Request: services.SendRegisterOTPReq{Email: "local+sub@domain.sub.tld"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{
					Local:      "local",
					SubAddress: "+sub",
					Domain:     "domain.sub.tld",
				}, nil
			},
			GetUserByEmailMock: func(ctx context.Context, email models.Email) (models.User, error) {
				return models.User{}, errors.New("")
			},
			ExpectedError: domain.CustomizeError(domain.ErrInternalServer, services.ErrMsgInternalServer),
		},
		{
			Name:    "FailsEmail_InternalServer",
			Request: services.SendRegisterOTPReq{Email: "local+sub@domain.sub.tld"},
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
			SendEmailMock: func(ctx context.Context, to, subject, body string) error {
				return domain.ErrInternalServer
			},
			SetCacheMock: func(ctx context.Context, key, value string, exp time.Duration) error {
				return nil
			},
			ExpectedError: domain.CustomizeError(domain.ErrInternalServer, services.ErrMsgInternalServer),
		},
		{
			Name:    "FailsCache_InternalServer",
			Request: services.SendRegisterOTPReq{Email: "local+sub@domain.sub.tld"},
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
			SendEmailMock: func(ctx context.Context, to, subject, body string) error {
				return nil
			},
			SetCacheMock: func(ctx context.Context, key, value string, exp time.Duration) error {
				return domain.ErrInternalServer
			},
			ExpectedError: domain.CustomizeError(domain.ErrInternalServer, services.ErrMsgInternalServer),
		},
	}
	for i := range tc {
		t.Run(tc[i].Name, func(t *testing.T) {
			var sessionManager port.ISessionManager
			var cacher port.ICache = mocks.MockCache{
				SetMock: tc[i].SetCacheMock,
			}
			var emailer port.IEmail = mocks.MockEmailer{
				SendEmailMock: tc[i].SendEmailMock,
			}
			var repositor port.IRepository = mocks.MockRepository{
				GetUserByEmailMock: tc[i].GetUserByEmailMock,
			}
			var modeler models.IModels = &mocks.MockModels{
				NewEmailMock: tc[i].NewEmailMock,
			}
			service := services.NewService(repositor, sessionManager, cacher, emailer, modeler)

			err := service.SendRegisterOTP(context.TODO(), tc[i].Request)

			assert.Equal(t, tc[i].ExpectedError, err)
		})
	}
}
