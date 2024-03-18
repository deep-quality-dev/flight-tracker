package handlers

import (
	"github.com/deep-quality-dev/flight-tracker/pkg/configs"
	"github.com/gorilla/mux"
)

func registerHttpRoutes(config *configs.ServerConfig, muxer *mux.Router, handler *TrackerHandler) *mux.Router {
	muxer.HandleFunc("/calculate", handler.TracePath()).Methods("POST")

	return muxer
}
