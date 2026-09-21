package problem19
import logritelog "log"

import "testing"

// Tests
func TestProblem19_Func(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestProblem19_Func(t=%v)", t)
	tests := []struct {
		name     string
		expected int
	}{
		{"Problem 19 - Counting Sundays", 171},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Problem19()
			if got != test.expected {
				t.Errorf("Problem19() = got %v, want %v", got, test.expected)
			}
		})
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestProblem19_Func(t=%v)", t)
}

// Benchmarks
func BenchmarkProblem19_Func(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkProblem19_Func(b=%v)", b)
	for i := 0; i < b.N; i++ {
		Problem19()
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkProblem19_Func(b=%v)", b)
}
