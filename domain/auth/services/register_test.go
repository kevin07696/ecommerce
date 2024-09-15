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

func TestRegisterUser(t *testing.T) {
	tc := []struct {
		Name             string
		Request          services.RegisterUserReq
		NewEmailMock     func(email string) (models.Email, error)
		NewOTPMock       func(otp string) (models.OTP, error)
		GetCacheMock     func(ctx context.Context, key string) (string, error)
		NewUsernameMock  func(username string) (models.Username, error)
		NewRoleMock      func(role string) (models.Role, error)
		CreateUserMock   func(ctx context.Context, user models.User) error
		ExpectedError    *domain.CustomError
		ExpectedResponse services.RegisterUserResp
	}{
		{
			Name:    "Succeeds",
			Request: services.RegisterUserReq{Username: "user", Email: "local+sub@domain.sub.tld", Role: "developer", OTP: "12345678"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{
					Local:      "local",
					SubAddress: "+sub",
					Domain:     "domain.sub.tld",
				}, nil
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return models.OTP("12345678"), nil
			},
			GetCacheMock: func(ctx context.Context, key string) (string, error) {
				return "register12345678", nil
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return models.Username("user"), nil
			},
			NewRoleMock: func(role string) (models.Role, error) {
				return models.Role("developer"), nil
			},
			CreateUserMock: func(ctx context.Context, user models.User) error {
				return nil
			},
			ExpectedResponse: services.RegisterUserResp{Username: "user"},
		},
		{
			Name:    "FailsValidation_InvalidEmail",
			Request: services.RegisterUserReq{Username: "user", Email: "local+$ub@domain.sub.tld", Role: "developer", OTP: "12345678"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{}, models.ErrInvalidEmail
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return models.OTP("12345678"), nil
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return models.Username("user"), nil
			},
			NewRoleMock: func(role string) (models.Role, error) {
				return models.Role("developer"), nil
			},
			ExpectedError: domain.ValidationError(services.ErrMsgInvalidEmail),
		},
		{
			Name:    "FailsValidation_EmptyEmail",
			Request: services.RegisterUserReq{Username: "user", Email: "", Role: "developer", OTP: "12345678"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{}, models.ErrEmptyEmail
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return models.OTP("12345678"), nil
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return models.Username("user"), nil
			},
			NewRoleMock: func(role string) (models.Role, error) {
				return models.Role("developer"), nil
			},
			ExpectedError: domain.ValidationError(services.ErrMsgEmptyEmail),
		},
		{
			Name:    "FailsValidation_InvalidOTP",
			Request: services.RegisterUserReq{Username: "user", Email: "local+$ub@domain.sub.tld", Role: "developer", OTP: "1234567"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{}, models.ErrInvalidEmail
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return "", models.ErrInvalidOTP
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return models.Username("user"), nil
			},
			NewRoleMock: func(role string) (models.Role, error) {
				return models.Role("developer"), nil
			},
			ExpectedError: domain.ValidationError(services.ErrMsgInvalidOTP),
		},
		{
			Name:    "FailsValidation_EmptyOTP",
			Request: services.RegisterUserReq{Username: "user", Email: "", Role: "developer", OTP: ""},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{}, models.ErrEmptyEmail
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return "", models.ErrEmptyOTP
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return models.Username("user"), nil
			},
			NewRoleMock: func(role string) (models.Role, error) {
				return models.Role("developer"), nil
			},
			ExpectedError: domain.ValidationError(services.ErrMsgEmptyOTP),
		},
		{
			Name:    "FailsCache_InternalServer",
			Request: services.RegisterUserReq{Username: "user", Email: "local+sub@domain.sub.tld", Role: "developer", OTP: "12345678"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{
					Local:      "local",
					SubAddress: "+sub",
					Domain:     "domain.sub.tld",
				}, nil
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return models.OTP("12345678"), nil
			},
			GetCacheMock: func(ctx context.Context, key string) (string, error) {
				return "", domain.ErrInternalServer
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return models.Username("user"), nil
			},
			NewRoleMock: func(role string) (models.Role, error) {
				return models.Role("developer"), nil
			},
			ExpectedError: domain.CustomizeError(domain.ErrInternalServer, services.ErrMsgInternalServer),
		},
		{
			Name:    "FailsOTP_Mismatch",
			Request: services.RegisterUserReq{Username: "user", Email: "local+sub@domain.sub.tld", Role: "developer", OTP: "12345678"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{
					Local:      "local",
					SubAddress: "+sub",
					Domain:     "domain.sub.tld",
				}, nil
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return models.OTP("12345678"), nil
			},
			GetCacheMock: func(ctx context.Context, key string) (string, error) {
				return "register12345679", nil
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return models.Username("user"), nil
			},
			NewRoleMock: func(role string) (models.Role, error) {
				return models.Role("developer"), nil
			},
			ExpectedError: domain.CustomizeError(domain.ErrUnauthorized, services.ErrMsgUnauthorized),
		},
		{
			Name:    "FailsValidation_InvalidUsername",
			Request: services.RegisterUserReq{Username: "#user", Email: "local+sub@domain.sub.tld", Role: "developer", OTP: "12345678"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{
					Local:      "local",
					SubAddress: "+sub",
					Domain:     "domain.sub.tld",
				}, nil
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return models.OTP("12345678"), nil
			},
			GetCacheMock: func(ctx context.Context, key string) (string, error) {
				return "register12345678", nil
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return "", models.ErrInvalidUsername
			},
			NewRoleMock: func(role string) (models.Role, error) {
				return models.Role("developer"), nil
			},
			ExpectedError: domain.ValidationError(services.ErrMsgInvalidUsername),
		},
		{
			Name:    "FailsValidation_EmptyUsername",
			Request: services.RegisterUserReq{Username: "", Email: "local+sub@domain.sub.tld", Role: "developer", OTP: "12345678"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{
					Local:      "local",
					SubAddress: "+sub",
					Domain:     "domain.sub.tld",
				}, nil
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return models.OTP("12345678"), nil
			},
			GetCacheMock: func(ctx context.Context, key string) (string, error) {
				return "register12345678", nil
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return "", models.ErrEmptyUsername
			},
			NewRoleMock: func(role string) (models.Role, error) {
				return models.Role("developer"), nil
			},
			ExpectedError: domain.ValidationError(services.ErrMsgEmptyUsername),
		},
		{
			Name:    "FailsValidation_InvalidRole",
			Request: services.RegisterUserReq{Username: "user", Email: "local+sub@domain.sub.tld", Role: "Developer", OTP: "12345678"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{
					Local:      "local",
					SubAddress: "+sub",
					Domain:     "domain.sub.tld",
				}, nil
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return models.OTP("12345678"), nil
			},
			GetCacheMock: func(ctx context.Context, key string) (string, error) {
				return "register12345678", nil
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return models.Username("user"), nil
			},
			NewRoleMock: func(role string) (models.Role, error) {
				return "", models.ErrInvalidRole
			},
			ExpectedError: domain.ValidationError(services.ErrMsgInvalidRole),
		},
		{
			Name:    "FailsValidation_EmptyRole",
			Request: services.RegisterUserReq{Username: "user", Email: "local+sub@domain.sub.tld", Role: ""},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{
					Local:      "local",
					SubAddress: "+sub",
					Domain:     "domain.sub.tld",
				}, nil
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return models.OTP("12345678"), nil
			},
			GetCacheMock: func(ctx context.Context, key string) (string, error) {
				return "register12345678", nil
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return models.Username("user"), nil
			},
			NewRoleMock: func(role string) (models.Role, error) {
				return "", models.ErrEmptyRole
			},
			ExpectedError: domain.ValidationError(services.ErrMsgEmptyRole),
		},
		{
			Name:    "FailsRepo_DuplicateEmail",
			Request: services.RegisterUserReq{Username: "user", Email: "local+sub@domain.sub.tld", Role: "developer", OTP: "12345678"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{
					Local:      "local",
					SubAddress: "+sub",
					Domain:     "domain.sub.tld",
				}, nil
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return models.OTP("12345678"), nil
			},
			GetCacheMock: func(ctx context.Context, key string) (string, error) {
				return "register12345678", nil
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return models.Username("user"), nil
			},
			NewRoleMock: func(role string) (models.Role, error) {
				return models.Role("developer"), nil
			},
			CreateUserMock: func(ctx context.Context, user models.User) error {
				return domain.ErrDuplicateEmail
			},
			ExpectedError: domain.CustomizeError(domain.ErrDuplicateKey, services.ErrMsgEmailNotUnique),
		},
		{
			Name:    "FailsRepo_DuplicateUsername",
			Request: services.RegisterUserReq{Username: "user", Email: "local+sub@domain.sub.tld", Role: "developer", OTP: "12345678"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{
					Local:      "local",
					SubAddress: "+sub",
					Domain:     "domain.sub.tld",
				}, nil
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return models.OTP("12345678"), nil
			},
			GetCacheMock: func(ctx context.Context, key string) (string, error) {
				return "register12345678", nil
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return models.Username("user"), nil
			},
			NewRoleMock: func(role string) (models.Role, error) {
				return models.Role("developer"), nil
			},
			CreateUserMock: func(ctx context.Context, user models.User) error {
				return domain.ErrDuplicateUsername
			},
			ExpectedError: domain.CustomizeError(domain.ErrDuplicateKey, services.ErrMsgUsernameNotUnique),
		},
		{
			Name:    "FailsRepo_InternalServer",
			Request: services.RegisterUserReq{Username: "user", Email: "local+sub@domain.sub.tld", Role: "developer", OTP: "12345678"},
			NewEmailMock: func(email string) (models.Email, error) {
				return models.Email{
					Local:      "local",
					SubAddress: "+sub",
					Domain:     "domain.sub.tld",
				}, nil
			},
			NewOTPMock: func(otp string) (models.OTP, error) {
				return models.OTP("12345678"), nil
			},
			GetCacheMock: func(ctx context.Context, key string) (string, error) {
				return "register12345678", nil
			},
			NewUsernameMock: func(username string) (models.Username, error) {
				return models.Username("user"), nil
			},
			NewRoleMock: func(role string) (models.Role, error) {
				return models.Role("developer"), nil
			},
			CreateUserMock: func(ctx context.Context, user models.User) error {
				return domain.ErrInternalServer
			},
			ExpectedError: domain.CustomizeError(domain.ErrInternalServer, services.ErrMsgInternalServer),
		},
	}

	for i := range tc {
		t.Run(tc[i].Name, func(t *testing.T) {
			var sessionManager port.ISessionManager
			var cacher port.ICache = mocks.MockCache{
				GetMock: tc[i].GetCacheMock,
				DeleteMock: func(ctx context.Context, key string) error {
					return nil
				},
			}
			var emailer port.IEmail
			var repositor port.IRepository = mocks.MockRepository{
				CreateUserMock: tc[i].CreateUserMock,
			}
			var modeler models.IModels = mocks.MockModels{
				NewEmailMock:    tc[i].NewEmailMock,
				NewOTPMock:      tc[i].NewOTPMock,
				NewUsernameMock: tc[i].NewUsernameMock,
				NewRoleMock:     tc[i].NewRoleMock,
				NewUserMock: func(username models.Username, email models.Email, role models.Role) models.User {
					return models.User{
						Username: username,
						Email:    email,
						Role:     role,
					}
				},
			}
			service := services.NewService(repositor, sessionManager, cacher, emailer, modeler)
			resp, err := service.RegisterUser(context.TODO(), tc[i].Request)

			assert.Equal(t, tc[i].ExpectedError, err)
			assert.Equal(t, tc[i].ExpectedResponse, resp)
		})
	}
}
