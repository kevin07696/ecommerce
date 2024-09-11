package services

import (
	"context"

	"github.com/kevin07696/ecommerce/domain"
	"github.com/kevin07696/ecommerce/domain/auth/models"
)

type ValidateUsernameReq struct {
	Username string
}

func (s *Service) ValidateUsername(ctx context.Context, req ValidateUsernameReq) *domain.CustomError {
	username, err := s.modeler.NewUsername(req.Username)
	if err != nil {
		if err == models.ErrEmptyUsername {
			return domain.ValidationError(ErrMsgEmptyUsername)
		}
		return domain.ValidationError(ErrMsgInvalidUsername)
	}

	_, err = s.repository.GetUserByUsername(ctx, username)
	if err == nil {
		return domain.CustomizeError(domain.ErrDuplicateKey, ErrMsgUsernameNotUnique)
	}
	if err != domain.ErrNotFound {
		return domain.CustomizeError(domain.ErrInternalServer, ErrMsgInternalServer)

	}

	return nil
}
