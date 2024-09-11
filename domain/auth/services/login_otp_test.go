package services_test

import (
	"context"
	"testing"
	"time"

	"github.com/kevin07696/ecommerce/domain"
	"github.com/kevin07696/ecommerce/domain/auth/mocks"
	"github.com/kevin07696/ecommerce/domain/auth/models"
	"github.com/kevin07696/ecommerce/domain/auth/port"
	"github.com/kevin07696/ecommerce/domain/auth/services"
	"github.com/stretchr/testify/assert"
)

func TestSendLoginOTP(t *testing.T) {
	tc := []struct {
		Name                  string
		Request               services.SendLoginOTPReq
		NewEmailMock          func(email string) (models.Email, error)
		NewUsernameMock       func(username string) (models.Username, error)
		GetUserByEmailMock    func(ctx context.Context, email models.Email) (models.User, error)
		GetUserByUsernameMock func(ctx context.Context, username models.Username) (models.User, error)
		SetCacheMock          func(ctx context.Context, key, value string, exp time.Duration) error
		SendEmailMock         func(ctx context.Context, to, subject, body string) error
		ExpectedError         *domain.CustomError
	}{
		{
			Name:    "Succeeds_Username",
			Request: services.SendLoginOTPReq{UserId: "user"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{}, domain.ValidationError(services.ErrMsgInvalidEmail)
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return "user", nil
			},
			GetUserByUsernameMock: func(ctx context.Context, username models.Username) (models.User, error) {
				return models.User{
					Email: models.Email{
						Local:      "local",
						SubAddress: "+sub",
						Domain:     "domain.sub.tld",
					},
					Username: "user",
					Role:     "developer",
				}, nil
			},
			SetCacheMock: func(ctx context.Context, key, value string, exp time.Duration) error {
				return nil
			},
			SendEmailMock: func(ctx context.Context, to, subject, body string) error {
				return nil
			},
		},
		{
			Name:    "Succeeds_Email",
			Request: services.SendLoginOTPReq{UserId: "local+sub@domain.sub.tld"},
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
					Username: "user",
					Role:     "developer",
				}, nil
			},
			SetCacheMock: func(ctx context.Context, key, value string, exp time.Duration) error {
				return nil
			},
			SendEmailMock: func(ctx context.Context, to, subject, body string) error {
				return nil
			},
		},
		{
			Name:    "FailsValidation_UserIdEmpty",
			Request: services.SendLoginOTPReq{UserId: ""},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{}, models.ErrEmptyEmail
			},
			ExpectedError: domain.ValidationError(services.ErrMsgEmptyUserId),
		},
		{
			Name:    "FailsValidation_InvalidEmail",
			Request: services.SendLoginOTPReq{UserId: "local+$ub@domain.sub.tld"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{}, models.ErrInvalidEmail
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return models.Username(""), models.ErrInvalidUsername
			},
			ExpectedError: domain.ValidationError(services.ErrMsgInvalidUserId),
		},
		{
			Name:    "FailsValidation_InvalidUsername",
			Request: services.SendLoginOTPReq{UserId: "$user"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{}, models.ErrInvalidEmail
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return models.Username(""), models.ErrInvalidUsername
			},
			ExpectedError: domain.ValidationError(services.ErrMsgInvalidUserId),
		},
		{
			Name:    "FailsRepo_EmailNotFound",
			Request: services.SendLoginOTPReq{UserId: "local+sub@domain.sub.tld"},
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
			ExpectedError: domain.CustomizeError(domain.ErrNotFound, services.ErrMsgUserIdNotFound),
		},
		{
			Name:    "FailsRepo_UsernameNotFound",
			Request: services.SendLoginOTPReq{UserId: "user"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{}, domain.ValidationError(services.ErrMsgInvalidEmail)
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return "user", nil
			},
			GetUserByUsernameMock: func(ctx context.Context, username models.Username) (models.User, error) {
				return models.User{}, domain.ErrNotFound
			},
			ExpectedError: domain.CustomizeError(domain.ErrNotFound, services.ErrMsgUserIdNotFound),
		},
		{
			Name:    "FailsRepo_EmailInternalServer",
			Request: services.SendLoginOTPReq{UserId: "local+sub@domain.sub.tld"},
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
		{
			Name:    "FailsRepo_UsernameInternalServer",
			Request: services.SendLoginOTPReq{UserId: "user"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{}, domain.ValidationError(services.ErrMsgInvalidEmail)
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return "user", nil
			},
			GetUserByUsernameMock: func(ctx context.Context, username models.Username) (models.User, error) {
				return models.User{}, domain.ErrInternalServer
			},
			ExpectedError: domain.CustomizeError(domain.ErrInternalServer, services.ErrMsgInternalServer),
		},
		{
			Name:    "FailsCache_UsernameInternalServer",
			Request: services.SendLoginOTPReq{UserId: "user"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{}, domain.ValidationError(services.ErrMsgInvalidEmail)
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return "user", nil
			},
			GetUserByUsernameMock: func(ctx context.Context, username models.Username) (models.User, error) {
				return models.User{
					Email: models.Email{
						Local:      "local",
						SubAddress: "+sub",
						Domain:     "domain.sub.tld",
					},
					Username: "user",
					Role:     "developer",
				}, nil
			},
			SetCacheMock: func(ctx context.Context, key, value string, exp time.Duration) error {
				return domain.ErrInternalServer
			},
			SendEmailMock: func(ctx context.Context, to, subject, body string) error {
				return nil
			},
			ExpectedError: domain.CustomizeError(domain.ErrInternalServer, services.ErrMsgInternalServer),
		},
		{
			Name:    "FailsEmailer_UsernameInternalServer",
			Request: services.SendLoginOTPReq{UserId: "user"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{}, domain.ValidationError(services.ErrMsgInvalidEmail)
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return "user", nil
			},
			GetUserByUsernameMock: func(ctx context.Context, username models.Username) (models.User, error) {
				return models.User{
					Email: models.Email{
						Local:      "local",
						SubAddress: "+sub",
						Domain:     "domain.sub.tld",
					},
					Username: "user",
					Role:     "developer",
				}, nil
			},
			SetCacheMock: func(ctx context.Context, key, value string, exp time.Duration) error {
				return nil
			},
			SendEmailMock: func(ctx context.Context, to, subject, body string) error {
				return domain.ErrInternalServer
			},
			ExpectedError: domain.CustomizeError(domain.ErrInternalServer, services.ErrMsgInternalServer),
		},
		{
			Name:    "FailsCache_EmailInternalServer",
			Request: services.SendLoginOTPReq{UserId: "local+sub@domain.sub.tld"},
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
					Username: "user",
					Role:     "developer",
				}, nil
			},
			SetCacheMock: func(ctx context.Context, key, value string, exp time.Duration) error {
				return domain.ErrInternalServer
			},
			SendEmailMock: func(ctx context.Context, to, subject, body string) error {
				return nil
			},
			ExpectedError: domain.CustomizeError(domain.ErrInternalServer, services.ErrMsgInternalServer),
		},
		{
			Name:    "FailsEmailer_EmailInternalServer",
			Request: services.SendLoginOTPReq{UserId: "local+sub@domain.sub.tld"},
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
					Username: "user",
					Role:     "developer",
				}, nil
			},
			SetCacheMock: func(ctx context.Context, key, value string, exp time.Duration) error {
				return nil
			},
			SendEmailMock: func(ctx context.Context, to, subject, body string) error {
				return domain.ErrInternalServer
			},
			ExpectedError: domain.CustomizeError(domain.ErrInternalServer, services.ErrMsgInternalServer),
		},
	}

	for i := range tc {
		t.Run(tc[i].Name, func(t *testing.T) {
			var repositor port.IRepository = mocks.MockRepository{
				GetUserByEmailMock:    tc[i].GetUserByEmailMock,
				GetUserByUsernameMock: tc[i].GetUserByUsernameMock,
			}
			var sessionManager port.ISessionManager
			var cacher port.ICache = mocks.MockCache{
				SetMock: tc[i].SetCacheMock,
			}
			var emailer port.IEmail = mocks.MockEmailer{
				SendEmailMock: tc[i].SendEmailMock,
			}
			var modeler models.IModels = mocks.MockModels{
				NewEmailMock:    tc[i].NewEmailMock,
				NewUsernameMock: tc[i].NewUsernameMock,
			}

			service := services.NewService(repositor, sessionManager, cacher, emailer, modeler)
			err := service.SendLoginOTP(context.TODO(), tc[i].Request)

			assert.Equal(t, tc[i].ExpectedError, err)
		})
	}
}
