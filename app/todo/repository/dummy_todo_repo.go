package repository

import (
	"fmt"
	"time"

	"github.com/ramadani/go-api-skeleton/app/todo/data"
)

type DummyTodoRepo struct {
	todos []data.Todo
}

func (rp *DummyTodoRepo) All() []data.Todo {
	return rp.todos
}

func (rp *DummyTodoRepo) Create(title, body string) data.Todo {
	todos := rp.todos
	id := uint(len(todos) + 1)
	todo := data.Todo{id, title, body, time.Now().Format(time.RFC3339)}
	todos = append(todos, todo)
	rp.todos = todos

	return todo
}

func (rp *DummyTodoRepo) Find(id uint) data.Todo {
	return rp.todos[id-1]
}

func (rp *DummyTodoRepo) Update(title, body string, id uint) data.Todo {
	todo := rp.Find(id)
	todo.Title = title
	todo.Body = body
	rp.todos[id-1] = todo

	return todo
}

func (rp *DummyTodoRepo) Delete(id uint) bool {
	rp.todos = append(rp.todos[:(id-1)], rp.todos[(id-1)+1:]...)

	return true
}

func NewDummyTodoRepo() *DummyTodoRepo {
	todos := []data.Todo{}
	for i := 1; i <= 5; i++ {
		title := fmt.Sprintf("Great %d", i)
		body := fmt.Sprintf("Great Body %d", i)
		todos = append(todos, data.Todo{uint(i), title, body, time.Now().Format(time.RFC3339)})
	}
	return &DummyTodoRepo{todos}
}
