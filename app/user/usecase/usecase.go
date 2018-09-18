package usecase

import "github.com/ramadani/go-api-skeleton/app/user/data"

// Usecase for user usecase interface
type Usecase interface {
	Paginate(limit int) (data.UserPaginate, error)
	FindByID(id int64) (data.User, error)
	Create(name, email, password string) (data.User, error)
	Update(name string, id int64) (data.User, error)
	Delete(id int64) error
}
