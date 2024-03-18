package main

import (
	"context"
	"log"

	"github.com/deep-quality-dev/flight-tracker/pkg/configs"
	"github.com/deep-quality-dev/flight-tracker/pkg/handlers"
	"github.com/deep-quality-dev/flight-tracker/pkg/transport/http"
	"go.uber.org/fx"
)

// @title Flight Tracker
// @version 1.0
// @description Demo service for tracking flights.

// @host localhost:8080
// @BasePath /
func main() {
	errCh := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	app := fx.New(
		fx.Supply(errCh),
		fx.Provide(
			func() context.Context {
				return ctx
			},
		),
		configs.Module(),
		http.ServerModule(),
		handlers.Module(),
		fx.Invoke(http.RunServer),
	)
	if err := app.Start(ctx); err != nil {
		panic(err)
	}

	select {
	case <-ctx.Done():
		log.Println(ctx, "Context cancelled. Exiting...")
	case <-app.Done():
		log.Println(ctx, "Interrupt received. Exiting...")

		if err := app.Stop(ctx); err != nil {
			panic(err)
		}
	case err := <-errCh:
		log.Println(ctx, "App errored:", err)
	}
}
