package service

import (
	"context"
	"github.com/samarec1812/simple-todo-list/internal/app/entity/todo"
	"github.com/samarec1812/simple-todo-list/internal/app/entity/user"
)

type TodoListRepository interface {
	Create(context.Context, int64, todo.TodoList) error
}

type UserRepository interface {
	Create(context.Context, user.User) error
}

type App interface {
	// User methods
	CreateUser(context.Context, string, string) error

	// CreateTodo
	CreateTodoList(context.Context, int64, string, string) error
}

type app struct {
	userRepo     UserRepository
	todoListRepo TodoListRepository
}

func NewApp(userRepo UserRepository, todoListRepo TodoListRepository) App {
	return &app{
		userRepo:     userRepo,
		todoListRepo: todoListRepo,
	}
}
