package data

import "github.com/yorlend/gotodo/internal/models"

type PgTodoRepo struct {
	db *Client
}

func NewPgTodoRepo(db *Client) models.TodoRepository {
	return &PgTodoRepo{db: db}
}

type Todo struct {
	Id    int
	Title string
	Descr string
	Done  bool
}

func (r *PgTodoRepo) Create(
	title, desc string,
) error {
	todo := &Todo{
		Title: title,
		Descr: desc,
		Done:  false,
	}

	return r.db.Insert(todo)
}

func (r *PgTodoRepo) ReadAll() (models.Todos, error) {
	var todos []Todo
	_, err := r.db.Query(&todos, `select * from todos`)

	var mappedTodos models.Todos
	for _, t := range todos {
		mappedTodos = append(mappedTodos, models.Todo{
			Id:          t.Id,
			Title:       t.Title,
			Description: t.Descr,
			Completed:   t.Done,
		})
	}

	return mappedTodos, err
}

func (r *PgTodoRepo) Update(
	todo *models.Todo,
) error {
	var entity = &Todo{
		Id:    todo.Id,
		Title: todo.Title,
		Descr: todo.Description,
		Done:  todo.Completed,
	}
	_, err := r.db.Model(entity).WherePK().UpdateNotNull()
	return err
}

func (r *PgTodoRepo) Delete(
	id int,
) error {
	_, err := r.db.Model(&Todo{Id: id}).WherePK().Delete()
	return err
}

func (r *PgTodoRepo) Read(
	id int,
) (*models.Todo, error) {
	todo := &Todo{Id: id}
	err := r.db.Select(todo)
	return &models.Todo{
		Id:          todo.Id,
		Title:       todo.Title,
		Description: todo.Descr,
		Completed:   todo.Done,
	}, err
}
