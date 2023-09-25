package todo

import "github.com/go-playground/validator/v10"

func (u *TodoList) Validate() error {
	validate := validator.New()

	return validate.Struct(u)
}
