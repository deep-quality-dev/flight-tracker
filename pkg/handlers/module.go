package handlers

import (
	"github.com/deep-quality-dev/flight-tracker/pkg/flights"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

// Module FX module function wiring internal dependencies.
func Module() fx.Option {
	return fx.Options(
		fx.Provide(
			mux.NewRouter,
			flights.NewTracker,
			NewTrackerHandler,
		),
		fx.Invoke(registerHttpRoutes),
	)
}
