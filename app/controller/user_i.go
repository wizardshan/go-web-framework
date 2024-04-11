package controller

import (
	"app/controller/response"
	"app/repository"
	"app/service"
	"github.com/gin-gonic/gin"
	"pkg/cache"
)

type IUser interface {
	EditProfile(c *gin.Context) (response.Data, error)
	Find(c *gin.Context) (response.Data, error)
}

func NewUser(repo repository.IUser, serv service.IUser, cache cache.ICache) IUser {

	ctr := new(UserProxy)
	ctr.repo = repo
	ctr.serv = serv
	ctr.cache = cache

	return ctr
}
