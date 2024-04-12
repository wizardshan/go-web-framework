package repository

import (
	"app/domain"
	"app/repository/ent"
	"app/repository/ent/user"
	"context"
)

func (repo *User) Register(ctx context.Context, option func(builder *ent.UserCreate)) (*domain.User, error) {

	var domainUser *domain.User
	err := repo.withTx(ctx, repo.db, func(tx *ent.Tx) error {
		db := tx.Client()
		domainUser = repo.create(ctx, db, option)
		hashID := "abcd"
		repo.update(ctx, db, func(builder *ent.UserUpdate) {
			builder.SetHashID(hashID).Where(user.ID(domainUser.ID))
		})
		return nil
	})

	return domainUser, err
}

//.Order(
//func(selector *sql.Selector) {
//	selector.OrderBy("rand()")
//})
