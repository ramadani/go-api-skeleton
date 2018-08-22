package main

import (
	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/bootstrap"
)

func main() {
	app := bootstrap.New(echo.New())
	app.Run()
}
