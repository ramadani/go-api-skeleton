package repository

import "github.com/ramadani/go-api-skeleton/app/user/data"

// Repository interface of user repository
type Repository interface {
	Paginate(offset, limit uint) ([]data.User, uint, error)
	Create(name, email, password string) (data.User, error)
	FindByID(id uint) (data.User, error)
	Update(name string, id uint) (data.User, error)
	Delete(id uint) error
}
