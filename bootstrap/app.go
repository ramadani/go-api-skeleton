package bootstrap

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/ramadani/go-api-skeleton/config"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// App contains the libraries that can be used in the app.
type App struct {
	e         *echo.Echo
	cog       *config.Config
	bootables []Bootable
}

// AddBootable to run the boot on startup
func (app *App) AddBootable(bootable Bootable) {
	app.bootables = append(app.bootables, bootable)
}

// Run and serve the app.
func (app *App) Run() {
	app.boot()
	app.serve()
	app.shutdown()
}

// boot is to use execute the bootables code before their run.
func (app *App) boot() {
	for _, bootable := range app.bootables {
		bootable.Boot()
	}
}

func (app *App) serve() {
	port := app.cog.Config.GetInt("app.port")
	isDebug := app.cog.Config.GetBool("app.debug")

	app.e.Logger.SetLevel(log.DEBUG)
	app.e.HideBanner = !isDebug
	app.e.Debug = isDebug

	go func() {
		if err := app.e.Start(fmt.Sprintf(":%d", port)); err != nil {
			app.e.Logger.Info("Shutting down the server")
		}
	}()
}

func (app *App) shutdown() {
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

// New returns app.
func New(e *echo.Echo, cog *config.Config) *App {
	return &App{e, cog, []Bootable{}}
}
