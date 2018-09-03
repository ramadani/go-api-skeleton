package auth

type UseCase interface {
	Attempt(email, password string) (string, error)
	Register(name, email, password string) (bool, error)
}
