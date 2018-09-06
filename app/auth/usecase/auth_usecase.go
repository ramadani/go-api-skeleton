package usecase

import (
	"github.com/ramadani/go-api-skeleton/app/auth/jwt"
	"github.com/ramadani/go-api-skeleton/app/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	rp  user.Repository
	jwt *jwt.Jwt
}

func (uc *AuthUseCase) Attempt(email, password string) (string, error) {
	user, uErr := uc.rp.FindByEmail(email)
	if uErr != nil {
		// TODO with err
		return "", uErr
	}

	bErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if bErr != nil {
		// TODO with err
		return "", bErr
	}

	token, tErr := uc.jwt.GenerateToken(user)
	if tErr != nil {
		// TODO with err
		return "", tErr
	}

	return token, tErr
}

func (uc *AuthUseCase) Register(name, email, password string) (bool, error) {
	bytes, errHash := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if errHash != nil {
		// TODO with err
		return false, errHash
	}

	passwordHashed := string(bytes)
	_, errCreate := uc.rp.Create(name, email, passwordHashed)
	if errCreate != nil {
		// TODO with err
		return false, errCreate
	}

	return true, nil
}

func NewUseCase(rp user.Repository, jwt *jwt.Jwt) *AuthUseCase {
	return &AuthUseCase{rp, jwt}
}
