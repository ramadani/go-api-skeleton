package usecase

import "github.com/ramadani/go-api-skeleton/app/user/data"

// Usecase for user usecase interface
type Usecase interface {
	Paginate(page, limit uint) (data.UserPaginate, error)
	FindByID(id uint) (data.User, error)
	Create(name, email, password string) (data.User, error)
	Update(name string, id uint) (data.User, error)
	Delete(id uint) error
}
