package problem11
import logritelog "log"

import "testing"

// Tests
func TestProblem11_Func(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestProblem11_Func(t=%v)", t)
	testCases := []struct {
		name     string
		expected uint
	}{
		{"Test Case 1", 70600674},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Problem11()
			if actual != tc.expected {
				t.Errorf("Expected: %v, but got %v", tc.expected, actual)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestProblem11_Func(t=%v)", t)
}

// Benchmark
func BenchmarkProblem11_Func(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkProblem11_Func(b=%v)", b)
	for i := 0; i < b.N; i++ {
		Problem11()
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkProblem11_Func(b=%v)", b)
}
