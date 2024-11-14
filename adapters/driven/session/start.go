package session

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-session/session"
	"github.com/kevin07696/ecommerce/domain"
)

const refreshKey = "refresh_timestamp"

func (s SessionManager) Start(ctx context.Context, w http.ResponseWriter, r *http.Request) (session.Store, error) {
	session, err := s.manager.Start(ctx, w, r)
	if err != nil {
		log.Printf("Fails to start session: %v, %v", domain.ErrInternalServer, err)
		return nil, domain.ErrInternalServer
	}
	if _, ok := session.Get(refreshKey); ok {
		return session, nil
	}
	if s.customOpts.refreshed > 0 {
		ts := time.Now().Add(time.Duration(time.Duration(s.customOpts.refreshed).Seconds()))
		session.Set(refreshKey, ts)
	}
	return session, nil
}
