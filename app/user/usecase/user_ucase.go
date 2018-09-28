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
	users, total, err := ucase.repo.Paginate(limit, offset)
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
	id, err := ucase.repo.Create(name, email, password)
	if err != nil {
		return data.User{}, err
	}

	user, err := ucase.repo.FindByID(id)
	return user, err
}

// FindByID for existing user in the storage
func (ucase *UserUsecase) FindByID(id uint) (data.User, error) {
	user, err := ucase.repo.FindByID(id)

	return user, err
}

// Update an existing user
func (ucase *UserUsecase) Update(name string, id uint) (data.User, error) {
	err := ucase.repo.Update(name, id)
	if err != nil {
		return data.User{}, err
	}

	user, err := ucase.FindByID(id)
	return user, err
}

// Delete a user by id
func (ucase *UserUsecase) Delete(id uint) error {
	err := ucase.repo.Delete(id)

	return err
}

// New user usecase
func New(repo repository.Repository) *UserUsecase {
	return &UserUsecase{repo: repo}
}
