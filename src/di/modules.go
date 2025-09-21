package di

import (
	"github.com/STommydx/VibeFolio/src/services"
	"go.uber.org/fx"
)

// Module is the Fx module definition for dependency injection
var Module = fx.Module("app",
	fx.Provide(
		// Services
		services.NewHealthService,
		services.NewConfigService,
	),
)

// HealthServiceParams defines the parameters required by the health service constructor
type HealthServiceParams struct {
	fx.In
}

// ConfigServiceParams defines the parameters required by the config service constructor
type ConfigServiceParams struct {
	fx.In
}
