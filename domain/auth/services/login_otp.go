package services

import (
	"context"

	"github.com/kevin07696/ecommerce/domain"
	"github.com/kevin07696/ecommerce/domain/auth/models"
)

type SendLoginOTPReq struct {
	UserId string
}

const loginTask = "login"

func (s *Service) SendLoginOTP(ctx context.Context, req SendLoginOTPReq) *domain.CustomError {
	email, err := s.modeler.NewEmail(req.UserId)
	if err == models.ErrEmptyEmail {
		return domain.ValidationError(ErrMsgEmptyUserId)
	}
	if err == nil {
		_, err = s.repository.GetUserByEmail(ctx, email)
		if err != nil {
			if err == domain.ErrNotFound {
				return domain.CustomizeError(err, ErrMsgUserIdNotFound)
			}
			return domain.CustomizeError(domain.ErrInternalServer, ErrMsgInternalServer)
		}
		err = s.sendOTP(ctx, email, loginTask)
		if err != nil {
			return domain.CustomizeError(domain.ErrInternalServer, ErrMsgInternalServer)
		}
		return nil
	}

	username, err := s.modeler.NewUsername(req.UserId)
	if err != nil {
		return domain.ValidationError(ErrMsgInvalidUserId)
	}
	user, err := s.repository.GetUserByUsername(ctx, username)
	if err != nil {
		if err == domain.ErrNotFound {
			return domain.CustomizeError(err, ErrMsgUserIdNotFound)
		}
		return domain.CustomizeError(domain.ErrInternalServer, ErrMsgInternalServer)
	}
	err = s.sendOTP(ctx, user.Email, loginTask)
	if err != nil {
		return domain.CustomizeError(domain.ErrInternalServer, ErrMsgInternalServer)
	}

	return nil
}
