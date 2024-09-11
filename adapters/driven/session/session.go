package session

import (
	"github.com/go-session/redis"
	"github.com/go-session/session"
)

func NewRedisStore(opts *redis.Options) session.ManagerStore {
	return redis.NewRedisStore(opts)
}

func NewSessionManager(opts ...session.Option) SessionManager {
	return SessionManager{
		manager: session.NewManager(opts...),
	}
}

type SessionManager struct {
	manager *session.Manager
}
