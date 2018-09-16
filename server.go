package main

import (
	"fmt"

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
	driver, connection := mysqlConnection(cog)
	db := db.New(driver, connection)
	md := middleware.Init(cog)
	app := bootstrap.New(e, cog)

	if cog.GetBool("db.auto_migrate") {
		app.AddBootable(providers.NewDbMigration(db))
	}

	app.AddBootable(providers.NewHTTP(e, md))
	app.AddBootable(authRestApi.New(e, db, cog, md))
	app.AddBootable(todoRestApi.New(e, db.DB))

	defer db.Close()

	app.Run()
}

func mysqlConnection(cog *config.Config) (string, string) {
	driver := cog.GetString("db.connections.mysql.driver")
	host := cog.GetString("db.connections.mysql.host")
	port := cog.GetString("db.connections.mysql.port")
	user := cog.GetString("db.connections.mysql.user")
	pass := cog.GetString("db.connections.mysql.password")
	dbName := cog.GetString("db.connections.mysql.db_name")
	format := "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True"

	return driver, fmt.Sprintf(format, user, pass, host, port, dbName)
}
