package web

import (
	"gonote/internals/db"
	"gonote/internals/models"
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
)

type Handler struct {
	t *db.TodoService
}

func NewHandler(t *db.TodoService) *Handler {
	return &Handler{t}
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", map[string]interface{}{})
}

func (h *Handler) ListTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.t.List()
	if err != nil {
		renderTemplate(w, "error.html", map[string]interface{}{
			"error":     err.Error(),
			"errorCode": http.StatusInternalServerError,
		})
		return
	}

	renderTemplate(w, "todos.html", map[string]interface{}{
		"todos": todos,
	})
}

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	todo := models.Todo{}

	err := r.ParseForm()

	if err != nil {
		renderTemplate(w, "error.html", map[string]interface{}{
			"error":     err.Error(),
			"errorCode": http.StatusInternalServerError,
		})
		return
	}

	content := r.PostForm.Get("Content")

	todo.Content = content
	err = h.t.Create(&todo)

	if err != nil {
		renderTemplate(w, "error.html", map[string]interface{}{
			"error":     err.Error(),
			"errorCode": http.StatusInternalServerError,
		})
		return
	}

	todos, err := h.t.List()
	if err != nil {
		renderTemplate(w, "error.html", map[string]interface{}{
			"error":     err.Error(),
			"errorCode": http.StatusInternalServerError,
		})
		return
	}

	renderTemplate(w, "todos.html", map[string]interface{}{
		"todos": todos,
	})
}

func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := uuid.FromStringOrNil(chi.URLParam(r, "id"))

	err := h.t.Delete(id)

	if err != nil {
		renderTemplate(w, "error.html", map[string]interface{}{
			"error":     err.Error(),
			"errorCode": http.StatusInternalServerError,
		})
		return
	}

	todos, err := h.t.List()
	if err != nil {
		renderTemplate(w, "error.html", map[string]interface{}{
			"error":     err.Error(),
			"errorCode": http.StatusInternalServerError,
		})
		return
	}

	renderTemplate(w, "todos.html", map[string]interface{}{
		"todos": todos,
	})
}

func (h *Handler) ListSearchTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.t.Search(r.FormValue("keyword"))
	if err != nil {
		renderTemplate(w, "error.html", map[string]interface{}{
			"error":     err.Error(),
			"errorCode": http.StatusInternalServerError,
		})
		return
	}

	renderTemplate(w, "todos.html", map[string]interface{}{
		"todos": todos,
	})
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("public/templates/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
