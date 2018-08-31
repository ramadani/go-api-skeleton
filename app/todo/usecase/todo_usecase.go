package usecase

import (
	"github.com/ramadani/go-api-skeleton/app/commons/response"
	"github.com/ramadani/go-api-skeleton/app/todo"

	"github.com/ramadani/go-api-skeleton/app/todo/resource"
)

type TodoUseCase struct {
	rp todo.Repository
}

func (uc *TodoUseCase) All() []resource.Todo {
	todos := uc.rp.All()

	return resource.Collection(todos)
}

func (uc *TodoUseCase) Create(title, body string) resource.Todo {
	todo := uc.rp.Create(title, body)

	return resource.Item(todo)
}

func (uc *TodoUseCase) Find(id uint) resource.Todo {
	todo := uc.rp.Find(id)

	return resource.Item(todo)
}

func (uc *TodoUseCase) Update(title, body string, id uint) resource.Todo {
	todo := uc.rp.Update(title, body, id)

	return resource.Item(todo)
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
