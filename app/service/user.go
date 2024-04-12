package service

import (
	"app/controller/request"
	"app/domain"
	"app/repository"
	"app/repository/ent"
	"app/repository/ent/user"
	"context"
	"database/sql"
)

type User struct {
	repo repository.IUser
}

func (serv *User) Get(ctx context.Context, id int) *domain.User {
	return serv.repo.Get(ctx, id)
}

func (serv *User) EditProfile(ctx context.Context, req *request.UserEditProfile) {
	serv.repo.Update(ctx, func(builder *ent.UserUpdate) {
		builder.SetBio(req.Bio).SetMobile(req.Mobile).SetNickname(req.Nickname).Where(user.ID(1))
	})
	serv.repo.Create(ctx, func(builder *ent.UserCreate) {
		builder.SetBio(req.Bio).SetMobile(req.Mobile).SetNickname(req.Nickname)
	})
}

func (serv *User) AllSql(ctx context.Context) (domain.Users, error) {
	query := "SELECT id, mobile FROM users WHERE mobile=? and id=?"
	var users domain.Users
	err := serv.repo.Query(ctx, query, []any{"13061919209", 9}, func(rows *sql.Rows) {
		var entUser ent.User
		rows.Scan(&entUser.ID, &entUser.Mobile)
		users = append(users, entUser.Mapper())
	})

	return users, err
}
