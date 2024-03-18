package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"go.uber.org/multierr"

	"github.com/deep-quality-dev/flight-tracker/pkg/flights"
)

// TrackerHandler represents an HTTP hander for flight tracking operations.
type TrackerHandler struct {
	FlightTracker *flights.Tracker
}

// NewTrackerHandler initializes a new instance of TrackerHandler.
func NewTrackerHandler(tracker *flights.Tracker) *TrackerHandler {
	return &TrackerHandler{
		FlightTracker: tracker,
	}
}

// TracePath godoc
// @Summary Trace start and end airport given a list of flight routes.
// @Description Trace start and end airport given a list of flight routes.
// @Tags flights
// @Accept json
// @Produce json
// @Param request body handlers.TracePath.request true "Flight Routes"
// @Router /calculate [post]
func (h *TrackerHandler) TracePath() http.HandlerFunc {
	type request [][]flights.Airport

	type response []flights.Airport

	return func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var reqBody request

		reqBytes, errReqBytes := io.ReadAll(r.Body)
		errReqUnmarshal := json.Unmarshal(reqBytes, &reqBody)

		errReq := multierr.Combine(errReqBytes, errReqUnmarshal)
		if errReq != nil {
			log.Println(ctx, "Could not unmarshal request params:", errReq)
			rw.WriteHeader(http.StatusBadRequest)
		}

		for _, route := range reqBody {
			h.FlightTracker.AddFlightRoute(route[0], route[1])
		}

		airports, err := h.FlightTracker.Trace()
		if err != nil {
			log.Println(ctx, "Could not calculate:", err)
			rw.WriteHeader(http.StatusBadRequest)
		}

		jsonResp, errRespMarshal := json.Marshal(response(airports))
		_, errRespWrite := rw.Write(jsonResp)

		errResp := multierr.Combine(errRespMarshal, errRespWrite)
		if errResp != nil {
			log.Println(ctx, "Could not write response:", errResp)
			rw.WriteHeader(http.StatusInternalServerError)
		}
	}
}
