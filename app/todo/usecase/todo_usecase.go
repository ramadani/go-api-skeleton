package usecase

import (
	"github.com/ramadani/go-api-skeleton/app/commons/response"
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

func (uc *TodoUseCase) Find(id uint) model.Todo {
	return uc.rp.Find(id)
}

func (uc *TodoUseCase) Update(title, body string, id uint) model.Todo {
	return uc.rp.Update(title, body, id)
}

func (uc *TodoUseCase) Delete(id uint) response.Message {
	if isDeleted := uc.rp.Delete(id); isDeleted == true {
		return response.Message{"Todo has been deleted"}
	}

	return response.Message{"Failed"}
}

func NewUseCase(rp todo.Repository) *TodoUseCase {
	return &TodoUseCase{rp}
}
