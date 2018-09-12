package main

import (
	"github.com/labstack/echo"
	authRestApi "github.com/ramadani/go-api-skeleton/app/auth/restapi"
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
	md := middleware.Init(cog)
	app := bootstrap.New(e, cog)

	if cog.Config.GetBool("db.auto_migrate") {
		app.AddBootable(providers.NewDbMigration(db))
	}

	app.AddBootable(providers.NewHTTP(e, md))
	app.AddBootable(authRestApi.New(e, db, cog, md))
	app.AddBootable(todoRestApi.New(e, db))

	defer db.Close()

	app.Run()
}
