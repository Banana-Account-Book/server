package users

import (
	"banana-account-book.com/internal/services/users/application"
	"banana-account-book.com/internal/services/users/infrastructure"
	"banana-account-book.com/internal/services/users/presentation"
	"go.uber.org/fx"
)

var Module = fx.Module("users",
	fx.Provide(infrastructure.NewUserRepository),
	fx.Provide(application.NewUserService),
	fx.Provide(presentation.NewUserController),
)
