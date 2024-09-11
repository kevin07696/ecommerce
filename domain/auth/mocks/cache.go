package mocks

import (
	"context"
	"time"
)

type MockCache struct {
	GetMock func(ctx context.Context, key string) (string, error)
	SetMock func(ctx context.Context, key, value string, exp time.Duration) error
	DeleteMock func(ctx context.Context, key string) error
}

func (c MockCache) Get(ctx context.Context, key string) (string, error) {
	return c.GetMock(ctx, key)
}

func (c MockCache) Set(ctx context.Context, key, value string, exp time.Duration) error {
	return c.SetMock(ctx, key, value, exp)
}

func (c MockCache) Delete(ctx context.Context, key string) error {
	return c.DeleteMock(ctx, key)
}
