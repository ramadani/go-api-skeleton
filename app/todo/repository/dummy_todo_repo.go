package repository

import (
	"fmt"
	"time"

	"github.com/ramadani/go-api-skeleton/app/todo/model"
)

type DummyTodoRepo struct {
	todos []model.Todo
}

func (rp *DummyTodoRepo) All() []model.Todo {
	return rp.todos
}

func (rp *DummyTodoRepo) Create(title, body string) model.Todo {
	todos := rp.todos
	id := uint(len(todos) + 1)
	todo := model.Todo{
		ID:        id,
		Title:     title,
		Body:      body,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	todos = append(todos, todo)
	rp.todos = todos

	return todo
}

func (rp *DummyTodoRepo) Find(id uint) model.Todo {
	return rp.todos[id-1]
}

func (rp *DummyTodoRepo) Update(title, body string, id uint) model.Todo {
	todo := rp.Find(id)
	todo.Title = title
	todo.Body = body
	todo.UpdatedAt = time.Now()
	rp.todos[id-1] = todo

	return todo
}

func (rp *DummyTodoRepo) Delete(id uint) bool {
	rp.todos = append(rp.todos[:(id-1)], rp.todos[(id-1)+1:]...)

	return true
}

func NewDummyRepo() *DummyTodoRepo {
	todos := []model.Todo{}
	for i := 1; i <= 5; i++ {
		title := fmt.Sprintf("Great %d", i)
		body := fmt.Sprintf("Great Body %d", i)
		todos = append(todos, model.Todo{
			ID:        uint(i),
			Title:     title,
			Body:      body,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}
	return &DummyTodoRepo{todos}
}
