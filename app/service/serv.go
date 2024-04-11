package service

import (
	"pkg/cache"
	"time"
)

type serv struct {
	cache cache.ICache
}

func (serv *serv) cacheDefaultExpiration() time.Duration {
	return 60 * time.Minute
}
