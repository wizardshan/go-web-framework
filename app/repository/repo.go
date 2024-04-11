package repository

import (
	"app/repository/ent"
	"pkg/cache"
)

type repoCache struct {
	cache cache.ICache
}

type repoDB struct {
	db *ent.Client
}
