package problem1
import logritelog "log"

import "testing"
import stdlog "log"
var log = stdlog.Default()

// Tests
func TestProblem1_Func(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestProblem1_Func(t=%v)", t)
	tests := []struct {
		name      string
		threshold uint
		want      uint
	}{
		{
			name:      "Testcase 1 - threshold 10",
			threshold: 10,
			want:      23,
		},
		{
			name:      "Testcase 2 - threshold 100",
			threshold: 100,
			want:      2318,
		},
		{
			name:      "Testcase 3 - threshold 1000",
			threshold: 1000,
			want:      233168,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Problem1(tt.threshold)

			if n != tt.want {
				t.Errorf("Problem1() = %v, want %v", n, tt.want)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestProblem1_Func(t=%v)", t)
}

// Benchmarks
func BenchmarkProblem1(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkProblem1(b=%v)", b)
	for i := 0; i < b.N; i++ {
		_ = Problem1(1000)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkProblem1(b=%v)", b)
}
