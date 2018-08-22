package bootstrap

import (
	"github.com/labstack/echo/middleware"
)

type Middleware struct {
	app App
}

func (m Middleware) Boot() {
	m.app.e.Use(middleware.Logger())
	m.app.e.Use(middleware.Recover())
}

func InitMiddleware(app App) *Middleware {
	return &Middleware{app}
}
