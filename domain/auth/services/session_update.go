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
	username, exists := session.Get(userKey)
	if !exists {
		err := s.sessionManager.Delete(ctx, w, r)
		if err != nil {
			return domain.CustomizeError(domain.ErrInternalServer, ErrMsgInternalServer)
		}
		return domain.CustomizeError(domain.ErrUnauthorized, ErrMsgUnauthorized)
	}
	if s.sessionManager.NeedsRefresh(session) {
		session, err = s.sessionManager.Refresh(ctx, w, r)
		if err != nil {
			return domain.CustomizeError(domain.ErrInternalServer, ErrMsgInternalServer)
		}
		session.Set(userKey, username)
	}
	if err := session.Save(); err != nil {
		log.Printf("Failed to save session: %v: %v", domain.ErrInternalServer, err)
		return domain.CustomizeError(domain.ErrInternalServer, ErrMsgInternalServer)
	}

	return nil
}
