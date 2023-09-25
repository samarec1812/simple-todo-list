package service

import (
	"context"
	"fmt"

	"github.com/samarec1812/simple-todo-list/internal/app/entity/todo"
)

func (a *app) CreateTodoList(ctx context.Context, userID int64, title, description string) error {
	list := todo.NewTodoList(title, description)

	err := list.Validate()
	if err != nil {
		return fmt.Errorf("error validation: %w", err)
	}

	err = a.todoListRepo.Create(ctx, userID, *list)
	if err != nil {
		return fmt.Errorf("error with create: %w", err)
	}

	return nil
}
