package service

import (
	"app/repository"
	"pkg/cache"
)

type IUser interface {
}

func NewUser(repo repository.IUser, cache cache.ICache) IUser {
	serv := new(UserProxy)
	serv.repo = repo
	serv.cache = cache
	return serv
}
