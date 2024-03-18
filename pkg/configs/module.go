package configs

import "go.uber.org/fx"

// Module FX module function wiring internal dependencies.
func Module() fx.Option {
	return fx.Options(
		fx.Provide(NewServerConfig),
	)
}
