package session

import (
	"log"
	"time"

	"github.com/go-session/session"
	"github.com/kevin07696/ecommerce/domain"
)

type Option func(*options)

type options struct {
	refreshed int64
}

var defaultOptions = options{
	refreshed: 0,
}

func SetRefresh(interval int64) Option {
	return func(o *options) {
		o.refreshed = interval
	}
}

type SessionManager struct {
	manager    *session.Manager
	customOpts *options
}

func NewSessionManager(copts []Option, opts ...session.Option) SessionManager {
	customOpts := defaultOptions
	for _, o := range copts {
		o(&customOpts)
	}
	return SessionManager{
		manager:    session.NewManager(opts...),
		customOpts: &customOpts,
	}
}

func (s SessionManager) NeedsRefresh(session session.Store) bool {
	if s.customOpts.refreshed <= 0 {
		return false
	}
	val, _ := session.Get(refreshKey)
	ts, isParsed := val.(time.Time)
	if !isParsed {
		log.Printf("Fails to parse refresh timestamp: %v, %v", domain.ErrInternalServer, val)
		return true
	}
	return time.Now().After(ts)
}
