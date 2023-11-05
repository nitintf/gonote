package db

import (
	"database/sql"
	"gonote/internals/models"

	"github.com/gofrs/uuid/v5"
)

var _ models.TodoService = (*TodoService)(nil)

type TodoService struct {
	db *sql.DB
}

func NewTodoService(db *sql.DB) *TodoService {
	return &TodoService{db: db}
}

func (s *TodoService) Find(id uuid.UUID) (todo models.Todo, err error) {
	err = s.db.QueryRow("SELECT * FROM todos WHERE id = $1", id).Scan(&todo.ID, &todo.Content)
	return todo, err
}

func (s *TodoService) Create(todo *models.Todo) error {
	todo.ID = uuid.Must(uuid.NewV4())
	_, err := s.db.Exec("INSERT INTO todos (id, content, completed) VALUES (?, ?, ?)", todo.ID, todo.Content, false)
	return err
}

func (s *TodoService) List() ([]models.Todo, error) {
	rows, err := s.db.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Content, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (s *TodoService) Search(term string) ([]models.Todo, error) {
	rows, err := s.db.Query("SELECT * FROM todos WHERE content LIKE $1", "%"+term+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Content, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil

}

func (s *TodoService) Delete(id uuid.UUID) error {
	_, err := s.db.Exec("DELETE FROM todos WHERE id = $1", id)
	return err
}

func (s *TodoService) Update(todo *models.Todo) error {
	_, err := s.db.Exec("UPDATE todos SET content = :content WHERE id = :id", todo)
	return err
}

func (s *TodoService) SetCompleted(id uuid.UUID, completed bool) error {
	_, err := s.db.Exec("UPDATE todos SET completed = $1 WHERE id = $2", completed, id)
	return err
}
