package todo

type TodoList struct {
	Title       string `db:"title" validate:"required"`
	Description string `db:"description" validate:"required"`
}

func NewTodoList(title, description string) *TodoList {
	return &TodoList{
		Title:       title,
		Description: description,
	}
}

func (u *TodoList) GetTodoListDBRecord() map[string]any {
	return map[string]any{
		"title":       u.Title,
		"description": u.Description,
	}
}
