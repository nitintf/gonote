package web

import (
	"gonote/internals/db"

	"github.com/go-chi/chi/v5"
)

func InitRoutes(s *chi.Mux, ts *db.TodoService) {
	h := NewHandler(ts)
	s.Get("/", h.Index)
	s.Get("/todos", h.ListTodos)
	s.Post("/todos", h.CreateTodo)
	s.Delete("/todos/{id}", h.DeleteTodo)
	s.Get("/todos/search", h.ListSearchTodos)
}
