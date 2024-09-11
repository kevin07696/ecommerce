package services

import (
	"context"
	"log"
	"net/http"

	"github.com/kevin07696/ecommerce/domain"
)

func (s *Service) UpdateSession(ctx context.Context, w http.ResponseWriter, r *http.Request) *domain.CustomError {
	session, err := s.sessionManager.Start(ctx, w, r)
	if err != nil {
		return domain.CustomizeError(domain.ErrInternalServer, ErrMsgInternalServer)
	}
	if _, exists := session.Get("user"); !exists {
		return domain.CustomizeError(domain.ErrUnauthorized, ErrMsgUnauthorized)
	}
	if err := session.Save(); err != nil {
		log.Printf("Failed to save session: %v: %v", domain.ErrInternalServer, err)
		return domain.CustomizeError(domain.ErrInternalServer, ErrMsgInternalServer)
	}

	return nil
}
