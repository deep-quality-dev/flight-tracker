package flights

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlightTracker(t *testing.T) {
	tests := []struct {
		input    [][]Airport
		expected []Airport
	}{
		{
			input:    [][]Airport{{"A", "B"}},
			expected: []Airport{"A", "B"},
		},
		{
			input:    [][]Airport{{"C", "B"}, {"A", "C"}},
			expected: []Airport{"A", "B"},
		},
		{
			input:    [][]Airport{{"D", "B"}, {"A", "C"}, {"E", "D"}, {"C", "E"}},
			expected: []Airport{"A", "B"},
		},
		{
			input:    [][]Airport{{"SFO", "EWR"}},
			expected: []Airport{"SFO", "EWR"},
		},
		{
			input:    [][]Airport{{"ATL", "EWR"}, {"SFO", "ATL"}},
			expected: []Airport{"SFO", "EWR"},
		},
		{
			input:    [][]Airport{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}},
			expected: []Airport{"SFO", "EWR"},
		},
	}
	for _, test := range tests {
		tracker := NewTracker()
		for _, row := range test.input {
			tracker.AddFlightRoute(row[0], row[1])
		}

		actual, err := tracker.Trace()

		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, actual, test.expected)
	}
}

func TestFlightTrackerCyclicDependencies(t *testing.T) {
	tests := []struct {
		input                  [][]Airport
		expectedPossibleCycles []string
	}{
		{
			input: [][]Airport{
				{"A", "B"},
				{"B", "A"},
			},
			expectedPossibleCycles: []string{
				"A -> B -> A",
				"B -> A -> B",
			},
		},
		{
			input: [][]Airport{
				{"A", "B"},
				{"B", "C"},
				{"C", "B"},
			},
			expectedPossibleCycles: []string{
				"B -> C -> B",
				"C -> B -> C",
				"B -> A -> B",
			},
		},
		{
			input: [][]Airport{
				{"A", "B"},
				{"B", "C"},
				{"C", "A"},
			},
			expectedPossibleCycles: []string{
				"A -> B -> C -> A",
				"B -> C -> A -> B",
				"C -> A -> B -> C",
			},
		},
		{
			input: [][]Airport{
				{"SFO", "EWR"},
				{"EWR", "SFO"},
			},
			expectedPossibleCycles: []string{
				"SFO -> EWR -> SFO",
				"EWR -> SFO -> EWR",
			},
		},
		{
			input: [][]Airport{
				{"SFO", "EWR"},
				{"EWR", "ATL"},
				{"ATL", "EWR"},
			},
			expectedPossibleCycles: []string{
				"EWR -> ATL -> EWR",
				"EWR -> SFO -> EWR",
				"ATL -> EWR -> ATL",
			},
		},
		{
			input: [][]Airport{
				{"SFO", "EWR"},
				{"EWR", "ATL"},
				{"ATL", "SFO"},
			},
			expectedPossibleCycles: []string{
				"SFO -> EWR -> ATL -> SFO",
				"EWR -> ATL -> SFO -> EWR",
				"ATL -> SFO -> EWR -> ATL",
			},
		},
	}

	for _, test := range tests {
		tracker := NewTracker()

		for _, row := range test.input {
			tracker.AddFlightRoute(row[0], row[1])
		}

		_, err := tracker.Trace()
		if err == nil {
			t.Error("Expected cycle error did not occur!")
		}

		if !errorContainsAnyOf(test.expectedPossibleCycles, err) {
			t.Errorf("Error does not print cycle: %q", err)
		}
	}
}

func errorContainsAnyOf(s []string, err error) bool {
	for _, v := range s {
		if strings.Contains(err.Error(), v) {
			return true
		}
	}

	return false
}
