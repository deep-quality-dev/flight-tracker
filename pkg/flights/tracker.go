package flights

import "github.com/deep-quality-dev/flight-tracker/pkg/models"

// Airport represents the code of an airport.
type Airport string

type history struct {
	ledger *models.Graph[Airport]
}

func newHistory() *history {
	return &history{
		ledger: models.NewGraph[Airport](),
	}
}

func (h *history) addFlightRoute(start, end Airport) {
	h.ledger.AddEdge(start, end)
}

func (h *history) getLongestFlightPath() ([]Airport, error) {
	return h.ledger.LongestPath()
}

// Tracker represents a flight tracker.
type Tracker struct {
	flightHistory *history
}

// NewTracker constructs a new instance of Tracker.
func NewTracker() *Tracker {
	return &Tracker{
		flightHistory: newHistory(),
	}
}

// AddFlightRoute adds a new flight route to the flight tracker.
func (t *Tracker) AddFlightRoute(start, end Airport) {
	t.flightHistory.addFlightRoute(start, end)
}

// Trace determines the start and end airport of complicated flight
// path constructed out of several flight routes.
func (t *Tracker) Trace() ([]Airport, error) {
	longestPath, err := t.flightHistory.getLongestFlightPath()
	if err != nil {
		return nil, err
	}

	fromAirport := longestPath[len(longestPath)-1]
	toAirport := longestPath[0]

	return []Airport{
		fromAirport,
		toAirport,
	}, nil
}
