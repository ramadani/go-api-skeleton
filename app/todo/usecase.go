package todo

import (
	"github.com/ramadani/go-api-skeleton/app/commons/response"
	"github.com/ramadani/go-api-skeleton/app/todo/resource"
)

// UseCase for todo's logic
type UseCase interface {
	All() []resource.Todo
	Create(title, body string) resource.Todo
	Find(id uint) resource.Todo
	Update(title, body string, id uint) resource.Todo
	Delete(id uint) response.Message
}
