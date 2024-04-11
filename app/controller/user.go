package controller

import (
	"app/controller/request"
	"app/controller/response"
	"app/repository"
	"app/service"
	"github.com/gin-gonic/gin"
)

type User struct {
	repo repository.IUser
	serv service.IUser
}

func (ctr *User) Find(c *gin.Context) (response.Data, error) {
	user := ctr.serv.Find(c.Request.Context(), 1)

	return user, nil
}

func (ctr *User) EditProfile(c *gin.Context) (response.Data, error) {
	req := new(request.UserEditProfile)
	if err := c.ShouldBind(req); err != nil {
		return nil, err
	}

	ctr.serv.EditProfile(c.Request.Context(), req)

	return req, nil
}
