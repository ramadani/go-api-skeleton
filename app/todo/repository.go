package todo

import (
	"github.com/ramadani/go-api-skeleton/app/todo/model"
)

// Repository interface for concrete repository
type Repository interface {
	All() []model.Todo
	Create(title, body string) model.Todo
	Find(id uint) model.Todo
	Update(title, body string, id uint) model.Todo
	Delete(id uint) bool
}
