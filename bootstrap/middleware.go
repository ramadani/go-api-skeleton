package bootstrap

import "github.com/labstack/echo/middleware"

type MiddlewareBoot struct {
	app App
}

func (b MiddlewareBoot) Boot() {
	b.app.e.Use(middleware.Logger())
	b.app.e.Use(middleware.Recover())
}

func NewMiddlewareBoot(app App) *MiddlewareBoot {
	return &MiddlewareBoot{app}
}
