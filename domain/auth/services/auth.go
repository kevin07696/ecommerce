package services

import (
	"context"
	"net/http"

	"github.com/kevin07696/ecommerce/domain"
	"github.com/kevin07696/ecommerce/domain/auth/models"
	"github.com/kevin07696/ecommerce/domain/auth/port"
)

type API interface {
	RegisterUser(context.Context, RegisterUserReq) (RegisterUserResp, *domain.CustomError)
	LoginUser(ctx context.Context, req LoginUserReq) (LoginUserResp, *domain.CustomError)
	GetUserByEmail(context.Context, GetUserByEmailReq) (GetUserByEmailResp, *domain.CustomError)
	SendLoginOTP(ctx context.Context, req SendLoginOTPReq) *domain.CustomError
	SendRegisterOTP(context.Context, SendRegisterOTPReq) *domain.CustomError
	ValidateUsername(context.Context, ValidateUsernameReq) *domain.CustomError
	CreateSession(ctx context.Context, w http.ResponseWriter, r *http.Request, req CreateSessionReq) *domain.CustomError
	UpdateSession(ctx context.Context, w http.ResponseWriter, r *http.Request) *domain.CustomError
	DeleteSession(ctx context.Context, w http.ResponseWriter, r *http.Request) *domain.CustomError
}

type Service struct {
	repository     port.IRepository
	sessionManager port.ISessionManager
	cacher         port.ICache
	emailer        port.IEmail
	modeler        models.IModels
}

func NewService(userRepositor port.IRepository, sessionManager port.ISessionManager, cacher port.ICache, emailer port.IEmail, modeler models.IModels) *Service {
	return &Service{
		repository:     userRepositor,
		sessionManager: sessionManager,
		cacher:         cacher,
		emailer:        emailer,
		modeler:        modeler,
	}
}
