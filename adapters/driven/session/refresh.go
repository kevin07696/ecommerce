package session

import (
	"context"
	"log"
	"net/http"

	"github.com/go-session/session"
	"github.com/kevin07696/ecommerce/domain"
)

func (s SessionManager) Refresh(ctx context.Context, w http.ResponseWriter, r *http.Request) (session.Store, error) {
	session, err := s.manager.Refresh(ctx, w, r)
	if err != nil {
		log.Printf("Fails to refresh session: %v, %v", domain.ErrInternalServer, err)
		return nil, domain.ErrInternalServer
	}

	return session, nil
}
