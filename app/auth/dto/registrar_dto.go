package dto

type RegistrarDto struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
}
