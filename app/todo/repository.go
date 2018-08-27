package todo

import "github.com/ramadani/go-api-skeleton/app/todo/data"

// Repository interface for concrete repository
type Repository interface {
	All() []data.Todo
	// Create()
	// Find(id uint)
	// Update(id uint)
	// Delete(id uint)
}
