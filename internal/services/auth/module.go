package auth

import (
	"banana-account-book.com/internal/libs/oauth"
	"banana-account-book.com/internal/services/auth/application"
	"banana-account-book.com/internal/services/auth/presentation"
	"go.uber.org/fx"
)

var Module = fx.Module("auth",
	fx.Provide(oauth.NewOAuthProvider),
	fx.Provide(application.NewAuthService),
	fx.Provide(presentation.NewAuthController),
)
