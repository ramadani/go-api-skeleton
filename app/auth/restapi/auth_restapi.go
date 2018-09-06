package restapi

import (
	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/app/auth/jwt"
	"github.com/ramadani/go-api-skeleton/app/auth/usecase"
	gormUserRepo "github.com/ramadani/go-api-skeleton/app/user/repository"
	"github.com/ramadani/go-api-skeleton/db"
)

type AuthRest struct {
	e  *echo.Echo
	db *db.Database
}

func (ar *AuthRest) Boot() {
	jwt := jwt.New("secret")
	userRepo := gormUserRepo.NewGormRepo(ar.db)
	usecase := usecase.NewUseCase(userRepo, jwt)
	handler := NewHandler(usecase)

	ar.e.POST("/login", handler.Attempt)
	ar.e.POST("/register", handler.Register)
}

func New(e *echo.Echo, db *db.Database) *AuthRest {
	return &AuthRest{e, db}
}
