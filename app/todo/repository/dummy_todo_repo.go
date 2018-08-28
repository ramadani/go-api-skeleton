package repository

import (
	"time"

	"github.com/ramadani/go-api-skeleton/app/todo/data"
)

type DummyTodoRepo struct{}

func (rp *DummyTodoRepo) All() []data.Todo {
	todos := []data.Todo{}
	todos = append(todos, data.Todo{1, "Great 1", "Great Body 1", time.Now().Format(time.RFC3339)})
	todos = append(todos, data.Todo{2, "Great 2", "Great Body 2", time.Now().Format(time.RFC3339)})
	todos = append(todos, data.Todo{3, "Great 3", "Great Body 3", time.Now().Format(time.RFC3339)})

	return todos
}

func NewDummyTodoRepo() *DummyTodoRepo {
	return &DummyTodoRepo{}
}
