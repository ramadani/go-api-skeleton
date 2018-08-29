package todo

import (
	"github.com/ramadani/go-api-skeleton/app/todo/data"
)

// Repository interface for concrete repository
type Repository interface {
	All() []data.Todo
	Create(title, body string) data.Todo
	Find(id uint) data.Todo
	Update(title, body string, id uint) data.Todo
	Delete(id uint) bool
}
