package service

import (
	"context"
	"fmt"

	"github.com/samarec1812/simple-todo-list/internal/app/entity/user"
)

func (a *app) CreateUser(ctx context.Context, name, username string) error {
	usr := user.NewUser(name, username)

	err := usr.Validate()
	if err != nil {
		return fmt.Errorf("error validation: %w", err)
	}

	err = a.userRepo.Create(ctx, *usr)
	if err != nil {
		return fmt.Errorf("error with create: %w", err)
	}

	return nil
}
