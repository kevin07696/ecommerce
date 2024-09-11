package session

import (
	"context"
	"log"
	"net/http"

	"github.com/go-session/session"
	"github.com/kevin07696/ecommerce/domain"
)

func (s SessionManager) Start(ctx context.Context, w http.ResponseWriter, r *http.Request) (session.Store, error) {
	session, err := s.manager.Start(ctx, w, r)
	if err != nil {
		log.Printf("Fails to start session: %v, %v", domain.ErrInternalServer, err)
		return nil, domain.ErrInternalServer
	}
	return session, nil
}
