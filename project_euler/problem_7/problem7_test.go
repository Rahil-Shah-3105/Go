package problem7
import logritelog "log"

import "testing"

// Tests
func TestProblem7_Func(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestProblem7_Func(t=%v)", t)
	tests := []struct {
		name  string
		input uint
		want  int64
	}{
		{
			name:  "Testcase 1 - input 6",
			input: 6,
			want:  13,
		},
		{
			name:  "Testcase 2 - input 10001",
			input: 10001,
			want:  104743,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Problem7(tt.input)

			if n != tt.want {
				t.Errorf("Problem7() = %v, want %v", n, tt.want)
			}
		})
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestProblem7_Func(t=%v)", t)
}

// Benchmarks
func BenchmarkProblem7(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkProblem7(b=%v)", b)
	for i := 0; i < b.N; i++ {
		_ = Problem7(10001)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkProblem7(b=%v)", b)
}
