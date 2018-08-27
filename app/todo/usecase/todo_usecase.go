package usecase

import (
	"github.com/ramadani/go-api-skeleton/app/todo"

	"github.com/ramadani/go-api-skeleton/app/todo/data"
)

type TodoUseCase struct {
	rp todo.Repository
}

func (td TodoUseCase) All() []data.Todo {
	return td.rp.All()
}

func NewTodoUseCase(rp todo.Repository) *TodoUseCase {
	return &TodoUseCase{rp}
}
