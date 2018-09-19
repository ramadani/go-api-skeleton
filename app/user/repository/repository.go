package repository

import "github.com/ramadani/go-api-skeleton/app/user/data"

// Repository interface of user repository
type Repository interface {
	Paginate(limit, offset int) ([]data.User, int64)
	Create(name, email, password string) (data.User, error)
	FindById(id int64) (data.User, error)
	Update(name string, id int64) (data.User, error)
	Delete(id int64) error
}
