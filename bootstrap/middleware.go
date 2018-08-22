package bootstrap

import (
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

type Middleware struct {
	app App
}

func (m Middleware) Boot() {
	if viper.GetBool("debug") {
		m.app.e.Use(middleware.Logger())
	}

	m.app.e.Use(middleware.Recover())
}

func InitMiddleware(app App) *Middleware {
	return &Middleware{app}
}
