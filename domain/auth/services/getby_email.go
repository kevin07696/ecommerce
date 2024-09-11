package services

import (
	"context"

	"github.com/kevin07696/ecommerce/domain"
	"github.com/kevin07696/ecommerce/domain/auth/models"
)

type GetUserByEmailReq struct {
	Email string
}

type GetUserByEmailResp struct {
	Username models.Username
}

func (s Service) GetUserByEmail(ctx context.Context, req GetUserByEmailReq) (GetUserByEmailResp, *domain.CustomError) {
	email, err := s.modeler.NewEmail(req.Email)
	if err != nil {
		if err == models.ErrEmptyEmail {
			return GetUserByEmailResp{}, domain.ValidationError(ErrMsgEmptyEmail)
		}
		return GetUserByEmailResp{}, domain.ValidationError(ErrMsgInvalidEmail)
	}

	user, err := s.repository.GetUserByEmail(ctx, email)
	if err != nil {
		if err == domain.ErrNotFound {
			return GetUserByEmailResp{}, domain.CustomizeError(err, ErrMsgEmailNotFound)
		}
		return GetUserByEmailResp{}, domain.CustomizeError(domain.ErrInternalServer, ErrMsgInternalServer)
	}

	return GetUserByEmailResp{Username: user.Username}, nil
}
