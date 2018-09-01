package main

import (
	"github.com/labstack/echo"
	todoRestApi "github.com/ramadani/go-api-skeleton/app/todo/restapi"
	"github.com/ramadani/go-api-skeleton/bootstrap"
	"github.com/ramadani/go-api-skeleton/config"
	"github.com/ramadani/go-api-skeleton/db"
	"github.com/ramadani/go-api-skeleton/middleware"
	"github.com/ramadani/go-api-skeleton/providers"
)

func main() {
	e := echo.New()
	cog := config.Init()
	db := db.Init(cog)
	md := middleware.Init()
	app := bootstrap.New(e, cog)

	if cog.Config.GetBool("db.auto_migration") {
		app.AddBootable(providers.NewDbMigration(db))
	}

	app.AddBootable(providers.NewHTTP(e, cog, md))
	app.AddBootable(todoRestApi.New(e, db))

	defer db.Close()

	app.Run()
}
