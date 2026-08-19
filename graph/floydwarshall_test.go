package graph
import logritelog "log"

import (
	"math"
	"testing"
)

const float64EqualityThreshold = 1e-9

// almostEqual subtracts two float64 variables and returns true if they differ less then float64EqualityThreshold
// reference: https://stackoverflow.com/a/47969546
func almostEqual(a, b float64) bool {
	logritelog.Default().Printf(">>> Entering almostEqual(a=%v,b=%v)", a, b)
	return math.Abs(a-b) <= float64EqualityThreshold
}

// IsAlmostEqualTo verifies if two WeightedGraphs can be considered almost equal
func (a *WeightedGraph) IsAlmostEqualTo(b WeightedGraph) bool {
	logritelog.Default().Printf(">>> Entering IsAlmostEqualTo(b=%v)", b)
	if len(*a) != len(b) {
		return false
	}

	for i := range *a {
		if len((*a)[i]) != len(b[i]) {
			return false
		}

		for j := range (*a)[i] {
			if (*a)[i][j] == Inf && b[i][j] == Inf {
				continue
			}

			if !almostEqual((*a)[i][j], b[i][j]) {
				return false
			}
		}
	}

	logritelog.Default().Printf("<<< Exiting IsAlmostEqualTo(b=%v)", b)
	return true
}

func TestFloydWarshall(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestFloydWarshall(t=%v)", t)
	var floydWarshallTestData = []struct {
		description string
		graph       [][]float64
		expected    [][]float64
	}{
		{
			description: "test empty graph",
			graph:       nil,
			expected:    nil,
		},
		{
			description: "test graph with wrong dimensions",
			graph: [][]float64{
				{1, 2},
				{Inf},
			},
			expected: nil,
		},
		{
			description: "test graph with no edges",
			graph: [][]float64{
				{Inf, Inf},
				{Inf, Inf},
			},
			expected: [][]float64{
				{Inf, Inf},
				{Inf, Inf},
			},
		},
		{
			description: "test graph with only negative edges",
			graph: [][]float64{
				{-3, -2},
				{-3, -2},
			},
			expected: [][]float64{
				{-17, -25},
				{-26, -34},
			},
		},
		{
			description: "test graph with 5 vertices and self-loops",
			graph: [][]float64{
				{1, 2, Inf, Inf, Inf},
				{Inf, Inf, 3, -4, Inf},
				{Inf, Inf, Inf, Inf, 5},
				{1, Inf, Inf, Inf, Inf},
				{Inf, Inf, Inf, 2, Inf},
			},
			expected: [][]float64{
				{-1, 1, 4, -3, 8},
				{-3, -1, 2, -5, 6},
				{7, 9, 12, 5, 5},
				{0, 2, 5, -2, 9},
				{2, 4, 7, 0, 9},
			},
		},
	}
	for _, test := range floydWarshallTestData {
		t.Run(test.description, func(t *testing.T) {
			result := FloydWarshall(test.graph)

			if !result.IsAlmostEqualTo(test.expected) {
				t.Logf("FAIL: %s", test.description)
				t.Fatalf("Expected result:%f\nFound: %f", test.expected, result)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestFloydWarshall(t=%v)", t)
}
