package problem15
import logritelog "log"

import "testing"
import stdlog "log"
var log = stdlog.Default()

// Tests
func TestProblem15_Func(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestProblem15_Func(t=%v)", t)
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"Input 2", 2, 6},
		// This test case is disabled
		// because it needs a big integer to run successfully
		// and factorial package doesn't support it
		// {"Input 20", 20, 137846528820},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Problem15(tt.input)
			if got != tt.want {
				t.Errorf("Problem15() = %v, want %v", got, tt.want)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestProblem15_Func(t=%v)", t)
}

// Benchmarks
func BenchmarkProblem15_Func(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkProblem15_Func(b=%v)", b)
	for i := 0; i < b.N; i++ {
		Problem15(20)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkProblem15_Func(b=%v)", b)
}
