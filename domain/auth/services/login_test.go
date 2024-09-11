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

func TestLogin(t *testing.T) {
	tc := []struct {
		Name                  string
		Request               services.LoginUserReq
		NewEmailMock          func(email string) (models.Email, error)
		NewUsernameMock       func(username string) (models.Username, error)
		NewOTPMock            func(otp string) (models.OTP, error)
		GetCacheMock          func(ctx context.Context, key string) (string, error)
		GetUserByEmailMock    func(ctx context.Context, email models.Email) (models.User, error)
		GetUserByUsernameMock func(ctx context.Context, username models.Username) (models.User, error)
		ExpectedError         *domain.CustomError
		ExpectedResponse      services.LoginUserResp
	}{
		{
			Name:    "Succeeds_Email",
			Request: services.LoginUserReq{UserId: "local+sub@domain.sub.tld", OTP: "12345678"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{
					Local:      "local",
					SubAddress: "+sub",
					Domain:     "domain.sub.tld",
				}, nil
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return "12345678", nil
			},
			GetCacheMock: func(ctx context.Context, key string) (string, error) {
				return "login12345678", nil
			},
			GetUserByEmailMock: func(ctx context.Context, email models.Email) (models.User, error) {
				return models.User{
					Email: models.Email{
						Local:      "local",
						SubAddress: "+sub",
						Domain:     "domain.sub.tld",
					},
					Role:     "developer",
					Username: "user",
				}, nil
			},
			ExpectedResponse: services.LoginUserResp{Username: "user", Role: "developer"},
		},
		{
			Name:    "Succeeds_Username",
			Request: services.LoginUserReq{UserId: "user", OTP: "12345678"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{}, models.ErrInvalidEmail
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return models.Username("user"), nil
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return "12345678", nil
			},
			GetCacheMock: func(ctx context.Context, key string) (string, error) {
				return "login12345678", nil
			},
			GetUserByUsernameMock: func(ctx context.Context, username models.Username) (models.User, error) {
				return models.User{
					Email: models.Email{
						Local:      "local",
						SubAddress: "+sub",
						Domain:     "domain.sub.tld",
					},
					Role:     "developer",
					Username: "user",
				}, nil
			},
			ExpectedResponse: services.LoginUserResp{Username: "user", Role: "developer"},
		},
		{
			Name:    "FailsValidation_EmptyOTP",
			Request: services.LoginUserReq{UserId: "local+sub@domain.sub.tld", OTP: "12345678"},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return "", models.ErrEmptyOTP
			},
			ExpectedError: domain.ValidationError(services.ErrMsgEmptyOTP),
		},
		{
			Name:    "FailsValidation_InvalidOTP",
			Request: services.LoginUserReq{UserId: "local+sub@domain.sub.tld", OTP: "12345678"},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return "", models.ErrInvalidOTP
			},
			ExpectedError: domain.ValidationError(services.ErrMsgInvalidOTP),
		},
		{
			Name:    "FailsValidation_EmptyUserId",
			Request: services.LoginUserReq{UserId: "", OTP: "12345678"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{}, models.ErrEmptyEmail
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return "12345678", nil
			},
			ExpectedError: domain.ValidationError(services.ErrMsgEmptyUserId),
		},
		{
			Name:    "FailsValidation_InvalidEmail",
			Request: services.LoginUserReq{UserId: "local+$ub@domain.sub.tld", OTP: "12345678"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{}, models.ErrInvalidEmail
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return models.Username(""), models.ErrInvalidUsername
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return "12345678", nil
			},
			ExpectedError: domain.ValidationError(services.ErrMsgInvalidUserId),
		},
		{
			Name:    "FailsValidation_InvalidUsername",
			Request: services.LoginUserReq{UserId: "user", OTP: "12345678"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{}, models.ErrInvalidEmail
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return models.Username(""), models.ErrInvalidUsername
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return "12345678", nil
			},
			ExpectedError: domain.ValidationError(services.ErrMsgInvalidUserId),
		},
		{
			Name: "FailsCache_UsernameInternalServer",
			Request: services.LoginUserReq{UserId: "user", OTP: "12345678"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{}, models.ErrInvalidEmail
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return models.Username("user"), nil
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return "12345678", nil
			},
			GetUserByUsernameMock: func(ctx context.Context, username models.Username) (models.User, error) {
				return models.User{
					Email: models.Email{
						Local:      "local",
						SubAddress: "+sub",
						Domain:     "domain.sub.tld",
					},
					Role:     "developer",
					Username: "user",
				}, nil
			},
			GetCacheMock: func(ctx context.Context, key string) (string, error) {
				return "", domain.ErrInternalServer
			},
			ExpectedError: domain.CustomizeError(domain.ErrInternalServer, services.ErrMsgInternalServer),
		},
		{
			Name:    "FailsCache_EmailInternalServer",
			Request: services.LoginUserReq{UserId: "local+sub@domain.sub.tld", OTP: "12345678"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{
					Local:      "local",
					SubAddress: "+sub",
					Domain:     "domain.sub.tld",
				}, nil
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return "12345678", nil
			},
			GetCacheMock: func(ctx context.Context, key string) (string, error) {
				return "", domain.ErrInternalServer
			},
			ExpectedError: domain.CustomizeError(domain.ErrInternalServer, services.ErrMsgInternalServer),
		},
		{
			Name:    "FailsOTP_EmailMismatch",
			Request: services.LoginUserReq{UserId: "local+sub@domain.sub.tld", OTP: "12345678"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{
					Local:      "local",
					SubAddress: "+sub",
					Domain:     "domain.sub.tld",
				}, nil
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return "12345678", nil
			},
			GetCacheMock: func(ctx context.Context, key string) (string, error) {
				return "login12345679", nil
			},
			ExpectedError: domain.CustomizeError(domain.ErrUnauthorized, services.ErrMsgUnauthorized),
		},
		{
			Name:    "FailsOTP_UsernameMismatch",
			Request: services.LoginUserReq{UserId: "user", OTP: "12345678"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{}, models.ErrInvalidEmail
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return models.Username("user"), nil
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return "12345678", nil
			},
			GetCacheMock: func(ctx context.Context, key string) (string, error) {
				return "login12345679", nil
			},
			GetUserByUsernameMock: func(ctx context.Context, username models.Username) (models.User, error) {
				return models.User{
					Email: models.Email{
						Local:      "local",
						SubAddress: "+sub",
						Domain:     "domain.sub.tld",
					},
					Role:     "developer",
					Username: "user",
				}, nil
			},
			ExpectedError: domain.CustomizeError(domain.ErrUnauthorized, services.ErrMsgUnauthorized),
		},
		{
			Name: "FailsRepo_UsernameNotFound",
			Request: services.LoginUserReq{UserId: "user", OTP: "12345678"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{}, models.ErrInvalidEmail
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return models.Username("user"), nil
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return "12345678", nil
			},
			GetUserByUsernameMock: func(ctx context.Context, username models.Username) (models.User, error) {
				return models.User{}, domain.ErrNotFound
			},
			ExpectedError: domain.CustomizeError(domain.ErrNotFound, services.ErrMsgUserIdNotFound),
		},
		{
			Name: "FailsRepo_UsernameInternalServer",
			Request: services.LoginUserReq{UserId: "user", OTP: "12345678"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{}, models.ErrInvalidEmail
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return models.Username("user"), nil
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return "12345678", nil
			},
			GetUserByUsernameMock: func(ctx context.Context, username models.Username) (models.User, error) {
				return models.User{}, domain.ErrInternalServer
			},
			ExpectedError: domain.CustomizeError(domain.ErrInternalServer, services.ErrMsgInternalServer),

		},
		{
			Name: "FailsRepo_EmailNotFound",
			Request: services.LoginUserReq{UserId: "local+sub@domain.sub.tld", OTP: "12345678"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{
					Local:      "local",
					SubAddress: "+sub",
					Domain:     "domain.sub.tld",
				}, nil
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return "12345678", nil
			},
			GetCacheMock: func(ctx context.Context, key string) (string, error) {
				return "login12345678", nil
			},
			GetUserByEmailMock: func(ctx context.Context, email models.Email) (models.User, error) {
				return models.User{}, domain.ErrNotFound
			},
			ExpectedError: domain.CustomizeError(domain.ErrNotFound, services.ErrMsgUserIdNotFound),
		},
		{
			Name: "FailsRepo_EmailInternalServer",
			Request: services.LoginUserReq{UserId: "local+sub@domain.sub.tld", OTP: "12345678"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{
					Local:      "local",
					SubAddress: "+sub",
					Domain:     "domain.sub.tld",
				}, nil
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return "12345678", nil
			},
			GetCacheMock: func(ctx context.Context, key string) (string, error) {
				return "login12345678", nil
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
			var cacher port.ICache = mocks.MockCache{
				GetMock: tc[i].GetCacheMock,
			}
			var emailer port.IEmail
			var repositor port.IRepository = mocks.MockRepository{
				GetUserByEmailMock: tc[i].GetUserByEmailMock,
				GetUserByUsernameMock: tc[i].GetUserByUsernameMock,
			}
			var modeler models.IModels = &mocks.MockModels{
				NewEmailMock: tc[i].NewEmailMock,
				NewUsernameMock: tc[i].NewUsernameMock,
				NewOTPMock:   tc[i].NewOTPMock,
			}
			service := services.NewService(repositor, sessionManager, cacher, emailer, modeler)

			loginUserResp, err := service.LoginUser(context.TODO(), tc[i].Request)

			assert.Equal(t, tc[i].ExpectedError, err)
			assert.Equal(t, tc[i].ExpectedResponse, loginUserResp)
		})
	}
}
