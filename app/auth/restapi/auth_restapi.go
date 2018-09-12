package restapi

import (
	"fmt"
	"net/http"

	"github.com/ramadani/go-api-skeleton/config"
	"github.com/ramadani/go-api-skeleton/middleware"

	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/app/auth/jwt"
	"github.com/ramadani/go-api-skeleton/app/auth/usecase"
	gormUserRepo "github.com/ramadani/go-api-skeleton/app/user/repository"
	"github.com/ramadani/go-api-skeleton/db"
)

// AuthRest contains dependencies to handle auth
type AuthRest struct {
	e   *echo.Echo
	db  *db.Database
	cog *config.Config
	md  *middleware.Middleware
}

// Boot the auth rest api
func (ar *AuthRest) Boot() {
	jwt := jwt.New(ar.cog.Config.GetString("jwt.key"))
	userRepo := gormUserRepo.NewGormRepo(ar.db)
	usecase := usecase.NewUseCase(userRepo, jwt)
	handler := NewHandler(usecase)

	ar.e.POST("/login", handler.Attempt)
	ar.e.POST("/register", handler.Register)
	ar.e.GET("/secret", func(c echo.Context) error {
		user := jwt.GetUser(c.Get("user"))
		return c.String(http.StatusOK, fmt.Sprintf("Hello %s", user.Name))
	}, ar.md.Jwt())
}

// New auth restapi
func New(
	e *echo.Echo,
	db *db.Database,
	cog *config.Config,
	md *middleware.Middleware,
) *AuthRest {
	return &AuthRest{e, db, cog, md}
}
