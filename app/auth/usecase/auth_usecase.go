package usecase

import (
	"github.com/ramadani/go-api-skeleton/app/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	rp user.Repository
}

func (uc *AuthUseCase) Attempt(email, password string) (string, error) {
	user, err := uc.rp.FindByEmail(email)

	return user.Email, err
}

func (uc *AuthUseCase) Register(name, email, password string) (bool, error) {
	bytes, errHash := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if errHash != nil {
		return false, errHash
	}

	passwordHashed := string(bytes)
	_, errCreate := uc.rp.Create(name, email, passwordHashed)
	if errCreate != nil {
		return false, errCreate
	}

	return true, nil
}

func NewUseCase(rp user.Repository) *AuthUseCase {
	return &AuthUseCase{rp}
}
