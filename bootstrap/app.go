package bootstrap

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/viper"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type App struct {
	e *echo.Echo
}

func (app App) Run() {
	bootables := []Bootable{
		NewConfigBoot(),
		NewMiddlewareBoot(app),
		NewRouteBoot(app),
	}

	for _, bootable := range bootables {
		bootable.Boot()
	}

	port := viper.GetInt("port")
	app.e.Logger.SetLevel(log.INFO)

	go func() {
		if err := app.e.Start(fmt.Sprintf(":%d", port)); err != nil {
			app.e.Logger.Info("shutting down the server")
		}
	}()

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

func New(e *echo.Echo) *App {
	return &App{e}
}
