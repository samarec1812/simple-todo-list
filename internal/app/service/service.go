package service

import "context"

type TodoListRepository interface {
	Create(ctx context.Context, userID int) error
}

type UserRepository interface {
	Create(context.Context) error
}

type App interface {
	// User methods
	CreateUser(context.Context) error

	// CreateTodo
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

func (a *app) CreateUser(ctx context.Context) error {
	return a.userRepo.Create(ctx)
}

//func (a *app) SaveEvent(ctx context.Context, headers, body map[string]any) error {
//	err := entity.Validate(headers, body)
//	if err != nil {
//		return err
//	}
//	event := entity.NewEvent(headers, body)
//
//	return a.eventRepo.Save(ctx, *event)
//}
