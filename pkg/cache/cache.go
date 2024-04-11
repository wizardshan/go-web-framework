package cache

import (
	"context"
	"time"
)

type ICache interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (interface{}, bool)
}

func New() ICache {

	var cache ICache
	return cache
}
