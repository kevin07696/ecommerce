package port

import (
	"context"
	"time"
)

type ICache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, value string, exp time.Duration) error
	Delete(ctx context.Context, key string) error
}
