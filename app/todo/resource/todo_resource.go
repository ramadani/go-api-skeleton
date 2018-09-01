package resource

import (
	"time"

	"github.com/ramadani/go-api-skeleton/app/todo/model"
)

type Todo struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func Item(item model.Todo) Todo {
	return Todo{
		item.ID,
		item.Title,
		item.Body,
		item.CreatedAt.Format(time.RFC3339),
		item.UpdatedAt.Format(time.RFC3339),
	}
}

func Collection(items []model.Todo) []Todo {
	var collection []Todo
	for _, item := range items {
		collection = append(collection, Item(item))
	}

	return collection
}
