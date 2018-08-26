package main

import (
	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/bootstrap"
	"github.com/ramadani/go-api-skeleton/config"
	"github.com/ramadani/go-api-skeleton/db"
)

func main() {
	fw := echo.New()
	cog := config.Init()
	db := db.Init(cog)

	app := bootstrap.New(fw, cog, db)
	app.Boot()
	app.Run()
}
