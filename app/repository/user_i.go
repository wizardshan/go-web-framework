package repository

import (
	"app/domain"
	"app/repository/ent"
	"app/repository/ent/user"
	"context"
	"database/sql"
	"pkg/cache"
)

type IUser interface {
	Get(ctx context.Context, id int) *domain.User
	First(ctx context.Context, option func(builder *ent.UserQuery)) *domain.User
	All(ctx context.Context, option func(builder *ent.UserQuery)) domain.Users
	Query(ctx context.Context, query string, args []any, option func(rows *sql.Rows)) error
	Count(ctx context.Context, option func(builder *ent.UserQuery)) int
	Create(ctx context.Context, option func(builder *ent.UserCreate)) *domain.User
	Update(ctx context.Context, option func(builder *ent.UserUpdate)) int
}

func NewUser(db *ent.Client, cache cache.ICache) IUser {
	repo := new(UserProxy)
	repo.db = db
	repo.cache = cache
	return repo
}

type User struct {
	repoDB
}

func (repo *User) Get(ctx context.Context, id int) *domain.User {
	return repo.first(ctx, repo.db, func(builder *ent.UserQuery) {
		builder.Where(user.ID(id))
	})
}

func (repo *User) First(ctx context.Context, option func(builder *ent.UserQuery)) *domain.User {
	return repo.first(ctx, repo.db, option)
}

func (repo *User) first(ctx context.Context, db *ent.Client, option func(builder *ent.UserQuery)) *domain.User {
	builder := db.User.Query()
	option(builder)
	return builder.FirstX(ctx).Mapper()
}

func (repo *User) All(ctx context.Context, option func(builder *ent.UserQuery)) domain.Users {
	return repo.all(ctx, repo.db, option)
}

func (repo *User) all(ctx context.Context, db *ent.Client, option func(builder *ent.UserQuery)) domain.Users {
	builder := db.User.Query()
	option(builder)
	var ents ent.Users = builder.AllX(ctx)
	return ents.Mapper()
}

func (repo *User) Count(ctx context.Context, option func(builder *ent.UserQuery)) int {
	return repo.count(ctx, repo.db, option)
}

func (repo *User) count(ctx context.Context, db *ent.Client, option func(builder *ent.UserQuery)) int {
	builder := db.User.Query()
	option(builder)
	return builder.CountX(ctx)
}

func (repo *User) Update(ctx context.Context, option func(builder *ent.UserUpdate)) int {
	return repo.update(ctx, repo.db, option)
}

func (repo *User) update(ctx context.Context, db *ent.Client, option func(builder *ent.UserUpdate)) int {
	builder := db.User.Update()
	option(builder)
	return builder.SaveX(ctx)
}

func (repo *User) Create(ctx context.Context, option func(builder *ent.UserCreate)) *domain.User {
	return repo.create(ctx, repo.db, option)
}

func (repo *User) create(ctx context.Context, db *ent.Client, option func(builder *ent.UserCreate)) *domain.User {
	builder := db.User.Create()
	option(builder)
	return builder.SaveX(ctx).Mapper()
}
