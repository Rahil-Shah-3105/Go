package problem16
import logritelog "log"

import "testing"

// Tests
func TestProblem16_Func(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestProblem16_Func(t=%v)", t)
	tests := []struct {
		name     string
		exponent int64
		want     int64
	}{
		{"2^15", 15, 26},
		{"2^1000", 1000, 1366},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Problem16(tt.exponent); got != tt.want {
				t.Errorf("Problem16() = %v, want %v", got, tt.want)
			}
		})
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestProblem16_Func(t=%v)", t)
}

// Benchmark
func BenchmarkProblem16_Func(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkProblem16_Func(b=%v)", b)
	for i := 0; i < b.N; i++ {
		Problem16(1000)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkProblem16_Func(b=%v)", b)
}
