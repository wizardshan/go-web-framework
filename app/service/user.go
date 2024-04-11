package service

import (
	"app/controller/request"
	"app/domain"
	"app/repository"
	"app/repository/ent"
	"app/repository/ent/user"
	"context"
)

type User struct {
	repo repository.IUser
}

func (serv *User) Find(ctx context.Context, id int) *domain.User {
	return serv.repo.Find(ctx, id)
}

func (serv *User) EditProfile(ctx context.Context, req *request.UserEditProfile) {
	serv.repo.Update(ctx, func(builder *ent.UserUpdate) {
		builder.SetBio(req.Bio).SetMobile(req.Mobile).SetNickname(req.Nickname).Where(user.ID(1))
	})
}
