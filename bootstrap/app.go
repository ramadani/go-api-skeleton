package bootstrap

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/ramadani/go-api-skeleton/config"
	"github.com/ramadani/go-api-skeleton/db"
	"github.com/ramadani/go-api-skeleton/middleware"
	"github.com/ramadani/go-api-skeleton/providers"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// App contains the libraries that can be used in the app.
type App struct {
	e   *echo.Echo
	cog *config.Config
	db  *db.Database
	md  *middleware.Middleware
}

// boot is to use execute the bootables code before their run.
func (app *App) boot() {
	bootables := []Bootable{}

	if app.cog.Config.GetBool("db.auto_migration") {
		bootables = append(bootables, providers.NewDbMigration(app.db))
	}

	bootables = append(bootables, providers.NewHTTP(app.e, app.cog, app.md))

	for _, bootable := range bootables {
		bootable.Boot()
	}
}

func (app *App) serve() {
	port := app.cog.Config.GetInt("port")
	app.e.Logger.SetLevel(log.INFO)

	go func() {
		if err := app.e.Start(fmt.Sprintf(":%d", port)); err != nil {
			app.e.Logger.Info("Shutting down the server")
		}
	}()
}

func (app *App) shutdown() {
	defer app.db.Close()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.e.Shutdown(ctx); err != nil {
		app.e.Logger.Fatal(err)
	}
}

// Run and serve the app.
func (app *App) Run() {
	app.boot()
	app.serve()
	app.shutdown()
}

// New returns app.
func New(
	e *echo.Echo,
	cog *config.Config,
	db *db.Database,
	md *middleware.Middleware,
) *App {
	return &App{e, cog, db, md}
}
