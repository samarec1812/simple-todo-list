package todo

import (
	"context"
	"database/sql"
	"github.com/samarec1812/simple-todo-list/internal/app/entity/todo"
)

type Repository struct {
	db *sql.DB
}

func NewTodoListRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, userID int64, list todo.TodoList) error {
	return nil
}
