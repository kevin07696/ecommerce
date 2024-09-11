package services

import (
	"context"

	"github.com/kevin07696/ecommerce/domain"
	"github.com/kevin07696/ecommerce/domain/auth/models"
)

type SendRegisterOTPReq struct {
	Email string
}

const registerTask = "register"

func (s *Service) SendRegisterOTP(ctx context.Context, req SendRegisterOTPReq) *domain.CustomError {
	email, err := s.modeler.NewEmail(req.Email)
	if err != nil {
		if err == models.ErrEmptyEmail {
			return domain.ValidationError(ErrMsgEmptyEmail)
		}
		return domain.ValidationError(ErrMsgInvalidEmail)
	}

	_, err = s.repository.GetUserByEmail(ctx, email)
	if err == nil {
		return domain.CustomizeError(domain.ErrDuplicateKey, ErrMsgEmailNotUnique)
	}
	if err != domain.ErrNotFound {
		return domain.CustomizeError(domain.ErrInternalServer, ErrMsgInternalServer)
	}

	err = s.sendOTP(ctx, email, registerTask)
	if err != nil {
		return domain.CustomizeError(domain.ErrInternalServer, ErrMsgInternalServer)
	}

	return nil
}
