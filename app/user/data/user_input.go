package data

// UserInput for validate the input
type UserInput struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

// UserUpdateInput for validate the input
type UserUpdateInput struct {
	Name string `validate:"required"`
}
