package mocks

import (
	"context"
	"net/http"

	"github.com/kevin07696/ecommerce/domain"
	"github.com/kevin07696/ecommerce/domain/auth/services"
)

type MockAPI struct {
	RegisterUserMock     func(context.Context, services.RegisterUserReq) (services.RegisterUserResp, *domain.CustomError)
	GetUserByEmailMock   func(context.Context, services.GetUserByEmailReq) (services.GetUserByEmailResp, *domain.CustomError)
	LoginUserMock        func(context.Context, services.LoginUserReq) (services.LoginUserResp, *domain.CustomError)
	SendRegisterOTPMock  func(context.Context, services.SendRegisterOTPReq) *domain.CustomError
	SendLoginOTPMock     func(context.Context, services.SendLoginOTPReq) *domain.CustomError
	ValidateUsernameMock func(context.Context, services.ValidateUsernameReq) *domain.CustomError
	CreateSessionMock    func(ctx context.Context, w http.ResponseWriter, r *http.Request, req services.CreateSessionReq) *domain.CustomError
	UpdateSessionMock    func(ctx context.Context, w http.ResponseWriter, r *http.Request) *domain.CustomError
	DeleteSessionMock    func(ctx context.Context, w http.ResponseWriter, r *http.Request) *domain.CustomError
}

func (m *MockAPI) RegisterUser(ctx context.Context, req services.RegisterUserReq) (services.RegisterUserResp, *domain.CustomError) {
	return m.RegisterUserMock(ctx, req)
}

func (m *MockAPI) GetUserByEmail(ctx context.Context, req services.GetUserByEmailReq) (services.GetUserByEmailResp, *domain.CustomError) {
	return m.GetUserByEmailMock(ctx, req)
}

func (m *MockAPI) LoginUser(ctx context.Context, req services.LoginUserReq) (services.LoginUserResp, *domain.CustomError) {
	return m.LoginUserMock(ctx, req)
}

func (m *MockAPI) SendRegisterOTP(ctx context.Context, req services.SendRegisterOTPReq) *domain.CustomError {
	return m.SendRegisterOTPMock(ctx, req)
}

func (m *MockAPI) SendLoginOTP(ctx context.Context, req services.SendLoginOTPReq) *domain.CustomError {
	return m.SendLoginOTPMock(ctx, req)
}

func (m *MockAPI) ValidateUsername(ctx context.Context, req services.ValidateUsernameReq) *domain.CustomError {
	return m.ValidateUsernameMock(ctx, req)
}

func (m *MockAPI) CreateSession(ctx context.Context, w http.ResponseWriter, r *http.Request, req services.CreateSessionReq) *domain.CustomError {
	return m.CreateSessionMock(ctx, w, r, req)
}

func (m *MockAPI) UpdateSession(ctx context.Context, w http.ResponseWriter, r *http.Request) *domain.CustomError {
	return m.UpdateSessionMock(ctx, w, r)
}

func (m *MockAPI) DeleteSession(ctx context.Context, w http.ResponseWriter, r *http.Request) *domain.CustomError {
	return m.DeleteSessionMock(ctx, w, r)
}
