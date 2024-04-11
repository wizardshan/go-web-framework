package service

import (
	"app/controller/request"
	"app/domain"
	"app/repository"
	"context"
	"pkg/cache"
)

type IUser interface {
	Find(ctx context.Context, id int) *domain.User
	EditProfile(ctx context.Context, req *request.UserEditProfile)
}

func NewUser(repo repository.IUser, cache cache.ICache) IUser {
	serv := new(UserProxy)
	serv.repo = repo
	serv.cache = cache
	return serv
}
