package services

import (
	"context"

	"github.com/kevin07696/ecommerce/domain"
	"github.com/kevin07696/ecommerce/domain/auth/models"
)

type RegisterUserReq struct {
	OTP      string
	Username string
	Email    string
	Role     string
}

type RegisterUserResp struct {
	Username models.Username
}

func (s *Service) RegisterUser(ctx context.Context, req RegisterUserReq) (RegisterUserResp, *domain.CustomError) {
	otp, err := s.modeler.NewOTP(req.OTP)
	if err != nil {
		if err == models.ErrEmptyOTP {
			return RegisterUserResp{}, domain.ValidationError(ErrMsgEmptyOTP)
		}
		return RegisterUserResp{}, domain.ValidationError(ErrMsgInvalidOTP)
	}

	email, err := s.modeler.NewEmail(req.Email)
	if err != nil {
		if err == models.ErrEmptyEmail {
			return RegisterUserResp{}, domain.ValidationError(ErrMsgEmptyEmail)
		}
		return RegisterUserResp{}, domain.ValidationError(ErrMsgInvalidEmail)
	}

	err = s.processOTP(ctx, email, otp, registerTask)
	if err != nil {
		if err == domain.ErrUnauthorized {
			return RegisterUserResp{}, domain.CustomizeError(domain.ErrUnauthorized, ErrMsgUnauthorized)
		}
		return RegisterUserResp{}, domain.CustomizeError(domain.ErrInternalServer, ErrMsgInternalServer)
	}

	username, err := s.modeler.NewUsername(req.Username)
	if err != nil {
		if err == models.ErrEmptyUsername {
			return RegisterUserResp{}, domain.ValidationError(ErrMsgEmptyUsername)
		}
		return RegisterUserResp{}, domain.ValidationError(ErrMsgInvalidUsername)
	}

	role, err := s.modeler.NewRole(req.Role)
	if err != nil {
		if err == models.ErrEmptyRole {
			return RegisterUserResp{}, domain.ValidationError(ErrMsgEmptyRole)
		}
		return RegisterUserResp{}, domain.ValidationError(ErrMsgInvalidRole)
	}

	user := s.modeler.NewUser(username, email, role)

	err = s.repository.CreateUser(ctx, user)
	if err != nil {
		if err == domain.ErrDuplicateEmail {
			return RegisterUserResp{}, domain.CustomizeError(domain.ErrDuplicateKey, ErrMsgEmailNotUnique)
		}
		if err == domain.ErrDuplicateUsername {
			return RegisterUserResp{}, domain.CustomizeError(domain.ErrDuplicateKey, ErrMsgUsernameNotUnique)
		}
		return RegisterUserResp{}, domain.CustomizeError(domain.ErrInternalServer, ErrMsgInternalServer)
	}

	return RegisterUserResp{Username: user.Username}, nil
}
