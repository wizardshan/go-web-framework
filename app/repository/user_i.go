package repository

import (
	"app/repository/ent"
	"pkg/cache"
)

type IUser interface {
}

func NewUser(db *ent.Client, cache cache.ICache) IUser {
	repo := new(UserProxy)
	repo.db = db
	repo.cache = cache
	return repo
}
