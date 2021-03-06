package repository

import "github.com/ramadani/go-api-skeleton/app/user/data"

// Repository interface of user repository
type Repository interface {
	Paginate(limit, offset uint) ([]data.User, uint, error)
	Create(name, email, password string) (uint, error)
	FindByID(id uint) (data.User, error)
	Update(name string, id uint) error
	Delete(id uint) error
}
