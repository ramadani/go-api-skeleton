package main

import (
	"github.com/labstack/echo"
	"github.com/ramadani/go-api-skeleton/bootstrap"
	"github.com/ramadani/go-api-skeleton/config"
	"github.com/ramadani/go-api-skeleton/db"
	"github.com/ramadani/go-api-skeleton/middleware"
)

func main() {
	e := echo.New()
	cog := config.Init()
	db := db.Init(cog)
	md := middleware.Init()

	app := bootstrap.New(e, cog, db, md)
	app.Run()
}
