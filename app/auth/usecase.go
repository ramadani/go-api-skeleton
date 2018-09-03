package auth

import "github.com/ramadani/go-api-skeleton/app/auth/dto"

type UseCase interface {
	Attempt(email, password string) (string, error)
	Register(registrar *dto.RegistrarDto) (bool, error)
}
