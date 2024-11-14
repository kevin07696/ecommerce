package session

import (
	"github.com/go-session/redis"
	"github.com/go-session/session"
)

func NewRedisStore(opts *redis.Options) session.ManagerStore {
	return redis.NewRedisStore(opts)
}
