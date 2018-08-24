package providers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

type Middleware struct {
	fw  *echo.Echo
	cog *viper.Viper
}

func (md Middleware) Boot() {
	if md.cog.GetBool("debug") {
		md.fw.Use(middleware.Logger())
	}

	md.fw.Use(middleware.Recover())
}

func InitMiddleware(fw *echo.Echo, cog *viper.Viper) *Middleware {
	return &Middleware{fw, cog}
}
