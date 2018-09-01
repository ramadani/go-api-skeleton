package todo

import (
	"github.com/ramadani/go-api-skeleton/app/commons/response"
	"github.com/ramadani/go-api-skeleton/app/todo/model"
)

// UseCase for todo's logic
type UseCase interface {
	All() []model.Todo
	Create(title, body string) model.Todo
	Find(id uint) model.Todo
	Update(title, body string, id uint) model.Todo
	Delete(id uint) response.Message
}
