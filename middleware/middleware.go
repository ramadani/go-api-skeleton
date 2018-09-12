package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ramadani/go-api-skeleton/app/auth/jwt"
	"github.com/ramadani/go-api-skeleton/config"
)

// Middleware represent app middlewares
type Middleware struct {
	cog *config.Config
}

// Init returns middleware
func Init(cog *config.Config) *Middleware {
	return &Middleware{cog}
}

// Logger middleware
func (md *Middleware) Logger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			return !md.cog.Config.GetBool("app.debug")
		},
	})
}

// Recover middleware
func (md *Middleware) Recover() echo.MiddlewareFunc {
	return middleware.Recover()
}

// Jwt middleware
func (md *Middleware) Jwt() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &jwt.UserClaims{},
		SigningKey: []byte(md.cog.Config.GetString("jwt.key")),
	}

	return middleware.JWTWithConfig(config)
}
