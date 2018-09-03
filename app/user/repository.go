package user

import "github.com/ramadani/go-api-skeleton/app/user/model"

type Repository interface {
	FindByEmail(email string) (model.User, error)
}
