package main

import (
	"banana-account-book.com/internal/app"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(app.NewServer),
		fx.Invoke(func(*app.App) {}),
	).Run()
}
