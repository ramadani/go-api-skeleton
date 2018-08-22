package main

import (
	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/bootstrap"
)

func main() {
	app := bootstrap.New(echo.New())
	app.InitMiddleware()
	app.InitRoutes()
	app.Run(3000)
}
