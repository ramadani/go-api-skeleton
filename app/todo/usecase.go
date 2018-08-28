package todo

import "github.com/ramadani/go-api-skeleton/app/todo/data"

type UseCase interface {
	All() []data.Todo
	Create() data.Todo
	// Find(id uint)
	// Update(id uint)
	// Delete(id uint)
}
