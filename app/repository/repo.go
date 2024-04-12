package repository

import (
	"app/repository/ent"
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"pkg/cache"
)

type repoCache struct {
	cache cache.ICache
}

type repoDB struct {
	db *ent.Client
}

func (repo *repoDB) withTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", errRollback)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
}

func (repo *repoDB) Query(ctx context.Context, query string, args []any, option func(rows *sql.Rows)) error {
	return repo.query(ctx, repo.db, query, args, option)
}

func (repo *repoDB) query(ctx context.Context, db *ent.Client, query string, args []any, option func(rows *sql.Rows)) error {
	rows, err := db.DB().Query(query, args...)
	if err != nil {
		return err
	}

	for rows.Next() {
		option(rows)
	}
	return rows.Close()
}
