package todo

import "github.com/ramadani/go-api-skeleton/app/todo/data"

type UseCase interface {
	All() []data.TodoData
	// Create()
	// Find(id uint)
	// Update(id uint)
	// Delete(id uint)
}
