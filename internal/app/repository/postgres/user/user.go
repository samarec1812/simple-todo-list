package user

import (
	"context"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context) error {
	return nil
}
