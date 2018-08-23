package provider

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

type Middleware struct {
	fw *echo.Echo
}

func (md Middleware) Boot() {
	if viper.GetBool("debug") {
		md.fw.Use(middleware.Logger())
	}

	md.fw.Use(middleware.Recover())
}

func InitMiddleware(fw *echo.Echo) *Middleware {
	return &Middleware{fw}
}
