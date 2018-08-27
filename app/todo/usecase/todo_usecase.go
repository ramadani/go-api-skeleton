package usecase

import (
	"time"

	"github.com/ramadani/go-api-skeleton/app/todo/data"
)

type TodoUseCase struct{}

func (todo TodoUseCase) All() []data.TodoData {
	todos := []data.TodoData{}
	todos = append(todos, data.TodoData{1, "Great 1", "Great Body 1", time.Now().Format(time.RFC3339)})
	todos = append(todos, data.TodoData{2, "Great 2", "Great Body 2", time.Now().Format(time.RFC3339)})
	todos = append(todos, data.TodoData{3, "Great 3", "Great Body 3", time.Now().Format(time.RFC3339)})

	return todos
}

func NewTodoUseCase() *TodoUseCase {
	return &TodoUseCase{}
}
