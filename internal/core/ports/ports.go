package ports

import "github.com/HP-97/go-htmx-web-todomvc/internal/core/domain"

type TodoRepository interface {
	FindById(id uint) (*domain.TodoItem, error)
	All()([]domain.TodoItem, error)
	Save(data *domain.TodoItem) error
}

type TodoService interface {
	FindById(id uint) (*domain.TodoItem, error)
	All()([]domain.TodoItem, error)
	Save(data *domain.TodoItem) error
}
