package user

import "github.com/ramadani/go-api-skeleton/app/user/model"

type Repository interface {
	Create(name, email, password string) (model.User, error)
	FindByEmail(email string) (model.User, error)
}
