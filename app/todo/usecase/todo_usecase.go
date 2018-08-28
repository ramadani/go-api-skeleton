package usecase

import (
	"github.com/ramadani/go-api-skeleton/app/todo"

	"github.com/ramadani/go-api-skeleton/app/todo/data"
)

type TodoUseCase struct {
	rp todo.Repository
}

func (uc *TodoUseCase) All() []data.Todo {
	return uc.rp.All()
}

func (uc *TodoUseCase) Create() data.Todo {
	return uc.rp.Create()
}

func NewTodoUseCase(rp todo.Repository) *TodoUseCase {
	return &TodoUseCase{rp}
}
