package main

import (
	"banana-account-book.com/internal/app"
	"banana-account-book.com/internal/libs/db"
	"banana-account-book.com/internal/services/auth"
	"banana-account-book.com/internal/services/users"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func main() {
	fx.New(
		fx.Provide(db.Init),
		fx.Provide(app.NewServer),
		users.Module,
		auth.Module,
		fx.Invoke(func(*app.App) {}),
		fx.Invoke(func(*gorm.DB) {}),
	).Run()
}
