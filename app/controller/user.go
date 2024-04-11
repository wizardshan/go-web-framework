package controller

import (
	"app/controller/response"
	"app/repository"
	"app/service"
	"github.com/gin-gonic/gin"
)

type User struct {
	repo repository.IUser
	serv service.IUser
}

func (ctr *User) EditProfile(c *gin.Context) (response.Data, error) {

	return nil, nil
}
