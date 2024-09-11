package services

import (
	"context"

	"github.com/kevin07696/ecommerce/domain"
	"github.com/kevin07696/ecommerce/domain/auth/models"
)

type LoginUserReq struct {
	UserId string
	OTP    string
}

type LoginUserResp struct {
	Username models.Username
	Role     models.Role
}

func (s *Service) LoginUser(ctx context.Context, req LoginUserReq) (LoginUserResp, *domain.CustomError) {
	otp, err := s.modeler.NewOTP(req.OTP)
	if err != nil {
		if err == models.ErrEmptyOTP {
			return LoginUserResp{}, domain.ValidationError(ErrMsgEmptyOTP)
		}
		return LoginUserResp{}, domain.ValidationError(ErrMsgInvalidOTP)
	}

	email, err := s.modeler.NewEmail(req.UserId)
	if err == models.ErrEmptyEmail {
		return LoginUserResp{}, domain.ValidationError(ErrMsgEmptyUserId)
	}
	if err == nil {
		err = s.processOTP(ctx, email, otp, loginTask)
		if err != nil {
			if err == domain.ErrUnauthorized {
				return LoginUserResp{}, domain.CustomizeError(domain.ErrUnauthorized, ErrMsgUnauthorized)
			}
			return LoginUserResp{}, domain.CustomizeError(domain.ErrInternalServer, ErrMsgInternalServer)
		}
		user, err := s.repository.GetUserByEmail(ctx, email)
		if err != nil {
			if err == domain.ErrNotFound {
				return LoginUserResp{}, domain.CustomizeError(err, ErrMsgUserIdNotFound)
			}
			return LoginUserResp{}, domain.CustomizeError(domain.ErrInternalServer, ErrMsgInternalServer)
		}
		return LoginUserResp{Username: user.Username, Role: user.Role}, nil
	}

	username, err := s.modeler.NewUsername(req.UserId)
	if err != nil {
		return LoginUserResp{}, domain.ValidationError(ErrMsgInvalidUserId)
	}

	user, err := s.repository.GetUserByUsername(ctx, username)
	if err != nil {
		if err == domain.ErrNotFound {
			return LoginUserResp{}, domain.CustomizeError(err, ErrMsgUserIdNotFound)
		}
		return LoginUserResp{}, domain.CustomizeError(domain.ErrInternalServer, ErrMsgInternalServer)
	}

	if err := s.processOTP(ctx, user.Email, otp, loginTask); err != nil {
		if err == domain.ErrUnauthorized {
			return LoginUserResp{}, domain.CustomizeError(domain.ErrUnauthorized, ErrMsgUnauthorized)
		}
		return LoginUserResp{}, domain.CustomizeError(domain.ErrInternalServer, ErrMsgInternalServer)
	}

	return LoginUserResp{Username: user.Username, Role: user.Role}, nil
}
