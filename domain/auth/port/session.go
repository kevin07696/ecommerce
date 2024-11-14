package port

import (
	"context"
	"net/http"

	"github.com/go-session/session"
)

type ISessionManager interface {
	Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error
	Refresh(ctx context.Context, w http.ResponseWriter, r *http.Request) (session.Store, error)
	Start(ctx context.Context, w http.ResponseWriter, r *http.Request) (session.Store, error)
	NeedsRefresh(session session.Store) bool
}
