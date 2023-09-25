package user

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"

	"github.com/samarec1812/simple-todo-list/internal/app/entity/user"
)

const (
	usersTable = "users"
)

type Repository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, usr user.User) error {
	query, args, err := sq.Insert(usersTable).SetMap(usr.GetUserDBRecord()).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
