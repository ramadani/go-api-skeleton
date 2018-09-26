package usecase

import (
	"math"

	"github.com/ramadani/go-api-skeleton/app/user/data"
	"github.com/ramadani/go-api-skeleton/app/user/repository"
)

// UserUsecase contain the dependencies of user use case
type UserUsecase struct {
	repo repository.Repository
}

// Paginate the user collection
func (ucase *UserUsecase) Paginate(page, limit uint) (data.UserPaginate, error) {
	offset := (page - 1) * limit
	users, total, err := ucase.repo.Paginate(offset, limit)
	pages := math.Ceil(float64(total / limit))

	userPaginate := data.UserPaginate{
		Data:    users,
		Total:   total,
		PerPage: limit,
		Page:    page,
		Pages:   uint(pages),
	}

	return userPaginate, err
}

// Create a new user
func (ucase *UserUsecase) Create(name, email, password string) (data.User, error) {
	user := data.User{
		ID:    1,
		Name:  name,
		Email: email,
	}

	return user, nil
}

// FindByID for existing user in the storage
func (ucase *UserUsecase) FindByID(id uint) (data.User, error) {
	user := data.User{
		ID:    1,
		Name:  "Ramadani",
		Email: "email.ramadani@gmail.com",
	}

	return user, nil
}

// Update an existing user
func (ucase *UserUsecase) Update(name string, id uint) (data.User, error) {
	user := data.User{
		ID:    id,
		Name:  name,
		Email: "email.ramadani@gmail.com",
	}

	return user, nil
}

// Delete a user by id
func (ucase *UserUsecase) Delete(id uint) error {
	return nil
}

// New user usecase
func New(repo repository.Repository) *UserUsecase {
	return &UserUsecase{repo: repo}
}
