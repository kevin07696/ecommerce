package redis

import (
	"context"
	"log/slog"
	"time"

	"github.com/kevin07696/ecommerce/domain"
	"github.com/redis/go-redis/v9"
)

type RedisAdapter struct {
	client *redis.Client
}

func NewRedisAdapter(opts *redis.Options) RedisAdapter {
	return RedisAdapter{
		client: redis.NewClient(opts),
	}
}

func (r RedisAdapter) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			slog.LogAttrs(ctx, slog.LevelError, "Key does not exist in cache", slog.String("key", key))
			return "", domain.ErrNotFound
		}
		slog.LogAttrs(ctx, slog.LevelError, "Failed to get key from cache", slog.String("key", key), slog.Any("error", err))
		return "", domain.ErrInternalServer
	}
	return val, nil
}

func (r RedisAdapter) Set(ctx context.Context, key, value string, exp time.Duration) error {
	if err := r.client.Set(ctx, key, value, exp).Err(); err != nil {
		slog.LogAttrs(ctx, slog.LevelError, "Failed to set record in cache", slog.String("key", key), slog.String("value", value), slog.Any("error", err))
		return domain.ErrInternalServer
	}
	return nil
}

func (r RedisAdapter) Delete(ctx context.Context, key string) error {
	err := r.client.Expire(ctx, key, 0).Err()
	if err != nil {
		slog.LogAttrs(ctx, slog.LevelError, "Failed to expire key from cache", slog.String("key", key), slog.Any("error", err))
		return domain.ErrInternalServer
	}
	return nil
}
