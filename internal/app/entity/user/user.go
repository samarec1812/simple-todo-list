package user

type User struct {
	Name     string `db:"name" validate:"required"`
	Username string `db:"username" validate:"required"`
}

func NewUser(name, username string) *User {
	return &User{
		Name:     name,
		Username: username,
	}
}

func (u *User) GetUserDBRecord() map[string]any {
	return map[string]any{
		"name":     u.Name,
		"username": u.Username,
	}
}
