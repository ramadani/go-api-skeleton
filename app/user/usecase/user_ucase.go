package usecase

import "github.com/ramadani/go-api-skeleton/app/user/data"

// UserUsecase contain the dependencies of user use case
type UserUsecase struct{}

// Paginate the user collection
func (ucase *UserUsecase) Paginate(page, limit int) (data.UserPaginate, error) {
	users := []data.User{
		data.User{
			ID:    1,
			Name:  "Ramadani",
			Email: "email.ramadani@gmail.com",
		},
		data.User{
			ID:    2,
			Name:  "Example",
			Email: "email.example@gmail.com",
		},
	}

	pagination := data.UserPaginate{
		Data:    users,
		Total:   int64(len(users)),
		PerPage: limit,
		Page:    page,
		Pages:   10,
	}

	return pagination, nil
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
func (ucase *UserUsecase) FindByID(id int64) (data.User, error) {
	user := data.User{
		ID:    1,
		Name:  "Ramadani",
		Email: "email.ramadani@gmail.com",
	}

	return user, nil
}

// Update an existing user
func (ucase *UserUsecase) Update(name string, id int64) (data.User, error) {
	user := data.User{
		ID:    id,
		Name:  name,
		Email: "email.ramadani@gmail.com",
	}

	return user, nil
}

// Delete a user by id
func (ucase *UserUsecase) Delete(id int64) error {
	return nil
}

// New user usecase
func New() *UserUsecase {
	return &UserUsecase{}
}
