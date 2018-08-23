package main

import (
	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/bootstrap"
	"github.com/ramadani/go-api-skeleton/config"
)

func main() {
	fw := echo.New()
	cog := config.Init()

	app := bootstrap.New(fw, cog)
	app.Run()
}
