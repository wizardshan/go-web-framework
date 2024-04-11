package repository

import (
	"app/domain"
	"app/repository/ent"
	"context"
	"pkg/cache"
)

type IUser interface {
	Find(ctx context.Context, id int) *domain.User
	Update(ctx context.Context, callbackFunc func(builder *ent.UserUpdate)) *domain.User
}

func NewUser(db *ent.Client, cache cache.ICache) IUser {
	repo := new(UserProxy)
	repo.db = db
	repo.cache = cache
	return repo
}
