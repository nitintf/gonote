package models

import (
	"github.com/gofrs/uuid/v5"
)

type Todo struct {
	ID        uuid.UUID `json:"id"`
	Content   string    `json:"content"`
	Completed bool      `json:"completed"`
}

type TodoService interface {
	Find(id uuid.UUID) (Todo, error)
	Create(todo *Todo) error
	List() ([]Todo, error)
	Search(term string) ([]Todo, error)
	Delete(id uuid.UUID) error
	Update(todo *Todo) error
	SetCompleted(id uuid.UUID, completed bool) error
}
