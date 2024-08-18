package roles

import (
	"banana-account-book.com/internal/services/roles/infrastructure"
	"go.uber.org/fx"
)

var Module = fx.Module("role",
	fx.Provide(infrastructure.NewRoleRepository),
)
