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

func (rp *DummyTodoRepo) Create() data.Todo {
	todos := rp.todos
	id := uint(len(todos) + 1)
	todo := data.Todo{id, "Great Gan", "Great Gan Body", time.Now().Format(time.RFC3339)}
	todos = append(todos, todo)
	rp.todos = todos

	return todo
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
