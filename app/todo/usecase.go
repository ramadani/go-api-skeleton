package todo

import (
	"github.com/ramadani/go-api-skeleton/app/commons/response"
	"github.com/ramadani/go-api-skeleton/app/todo/data"
)

// UseCase for todo's logic
type UseCase interface {
	All() []data.Todo
	Create(title, body string) data.Todo
	Find(id uint) data.Todo
	Update(title, body string, id uint) data.Todo
	Delete(id uint) response.Message
}
