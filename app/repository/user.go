package repository

import (
	"app/domain"
	"app/repository/ent"
	"app/repository/ent/user"
	"context"
)

type User struct {
	repoDB
}

func (repo *User) Find(ctx context.Context, id int) *domain.User {

	entUser := repo.db.User.Query().Where(user.ID(id)).FirstX(ctx)

	var domainUser domain.User
	domainUser.ID = entUser.ID

	return &domainUser
}

func (repo *User) Update(ctx context.Context, callbackFunc func(builder *ent.UserUpdate)) *domain.User {
	builder := repo.db.User.Update()
	callbackFunc(builder)
	builder.SaveX(ctx)
	return nil
}
