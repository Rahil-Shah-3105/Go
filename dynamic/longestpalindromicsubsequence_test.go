package dynamic_test
import logritelog "log"

import (
	"fmt"
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

func lpsTestTemplate(t *testing.T, algorithm func(input string) int) {
	logritelog.Default().Printf(">>> Entering lpsTestTemplate(t=%v,algorithm=%v)", t, algorithm)
	testCases := []struct {
		input    string
		expected int
	}{
		{"BBABCBCAB", 7},
		{"BBBAB", 4},
		{"ABBD", 2},
		{"GEEKSFORGEEKS", 5},
		{"abcdefgh", 1},
		{"bbbab", 4},
		{"cbbd", 2},
		{"racexyzcxar", 7},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprint("test with ", tc.input), func(t *testing.T) {
			result := algorithm(tc.input)
			if tc.expected != result {
				t.Fatalf("expected %d, got %d", tc.expected, result)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting lpsTestTemplate(t=%v,algorithm=%v)", t, algorithm)
}

func TestLpsRec(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestLpsRec(t=%v)", t)
	lpsTestTemplate(t, dynamic.LpsRec)
	logritelog.Default().Printf("<<< Exiting TestLpsRec(t=%v)", t)
}

func TestLpsDp(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestLpsDp(t=%v)", t)
	lpsTestTemplate(t, dynamic.LpsDp)
	logritelog.Default().Printf("<<< Exiting TestLpsDp(t=%v)", t)
}
