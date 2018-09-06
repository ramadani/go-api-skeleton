package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ramadani/go-api-skeleton/app/user/model"
)

type Jwt struct {
	secret string
}

type UserClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func (j *Jwt) GenerateToken(user model.User) (string, error) {
	claims := &UserClaims{
		user.Name,
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(j.secret))

	return t, err
}

func New(secret string) *Jwt {
	return &Jwt{secret}
}
