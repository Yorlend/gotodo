package models

type TodoRepository interface {
	Create(title, desc string) error
	ReadAll() (Todos, error)
	Update(todo *Todo) error
	Delete(id int) error
}

type Todo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type Todos []Todo
