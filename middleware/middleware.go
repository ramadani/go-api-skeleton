package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Middleware represent app middlewares
type Middleware struct{}

// Init returns middleware
func Init() *Middleware {
	return &Middleware{}
}

// Logger middleware
func (md *Middleware) Logger() echo.MiddlewareFunc {
	return middleware.Logger()
}

// Recover middleware
func (md *Middleware) Recover() echo.MiddlewareFunc {
	return middleware.Recover()
}
