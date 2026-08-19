package problem20
import logritelog "log"

import "testing"
import stdlog "log"
var log = stdlog.Default()

// Tests
func TestProblem20_Func(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestProblem20_Func(t=%v)", t)
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"Problem 20 - Factorial digit sum", 10, 27},
		{"Problem 20 - Factorial digit sum", 100, 648},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Problem20(test.input)
			if got != test.expected {
				t.Errorf("Problem20() = got %v, want %v", got, test.expected)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestProblem20_Func(t=%v)", t)
}

// Benchmarks
func BenchmarkProblem20_Func(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkProblem20_Func(b=%v)", b)
	for i := 0; i < b.N; i++ {
		Problem20(100)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkProblem20_Func(b=%v)", b)
}
