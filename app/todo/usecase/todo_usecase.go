package usecase

import (
	"github.com/ramadani/go-api-skeleton/app/todo"

	"github.com/ramadani/go-api-skeleton/app/todo/model"
)

type TodoUseCase struct {
	rp todo.Repository
}

func (uc *TodoUseCase) All() []model.Todo {
	return uc.rp.All()
}

func (uc *TodoUseCase) Create(title, body string) model.Todo {
	return uc.rp.Create(title, body)
}

func (uc *TodoUseCase) Find(id uint) (model.Todo, error) {
	return uc.rp.Find(id)
}

func (uc *TodoUseCase) Update(title, body string, id uint) (model.Todo, error) {
	return uc.rp.Update(title, body, id)
}

func (uc *TodoUseCase) Delete(id uint) error {
	return uc.rp.Delete(id)
}

func NewUseCase(rp todo.Repository) *TodoUseCase {
	return &TodoUseCase{rp}
}
