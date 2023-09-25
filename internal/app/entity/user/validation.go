package user

import "github.com/go-playground/validator/v10"

func (u *User) Validate() error {
	validate := validator.New()

	return validate.Struct(u)
}
