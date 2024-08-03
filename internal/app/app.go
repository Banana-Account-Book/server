package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type App struct {
	*fiber.App
}

func New() *App {
	fiber := fiber.New()
	return &App{fiber}
}

func (app *App) Start(port string) error {
	return app.Listen(":" + port)
}

func (app *App) Stop() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		app.Shutdown()
	}()
}

func NewServer(lc fx.Lifecycle) *App {
	server := New()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				port := os.Getenv("PORT")
				fmt.Println("Server started on port:", port)
				server.Start(port)
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.Stop()
			return nil
		},
	})
	return server
}
