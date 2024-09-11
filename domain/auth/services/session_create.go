package services

import (
	"context"
	"log"
	"net/http"

	"github.com/kevin07696/ecommerce/domain"
	"github.com/kevin07696/ecommerce/domain/auth/models"
)

type CreateSessionReq struct {
	Username models.Username
}

func (s *Service) CreateSession(ctx context.Context, w http.ResponseWriter, r *http.Request, req CreateSessionReq) *domain.CustomError {
	session, err := s.sessionManager.Refresh(ctx, w, r)
	if err != nil {
		return domain.CustomizeError(domain.ErrInternalServer, ErrMsgInternalServer)
	}
	session.Set("user", req.Username)
	if err := session.Save(); err != nil {
		log.Printf("Failed to save session: %v: %v", domain.ErrInternalServer, err)
		return domain.CustomizeError(domain.ErrInternalServer, ErrMsgInternalServer)
	}

	return nil
}
