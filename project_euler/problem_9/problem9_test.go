package problem9
import logritelog "log"

import "testing"

// Tests
func TestProblem9_Func(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestProblem9_Func(t=%v)", t)
	tests := []struct {
		name string
		want uint
	}{
		{
			name: "Testcase 1",
			want: 31875000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Problem9()

			if n != tt.want {
				t.Errorf("Problem9() = %v, want %v", n, tt.want)
			}
		})
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestProblem9_Func(t=%v)", t)
}

// Benchmarks
func BenchmarkProblem9(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkProblem9(b=%v)", b)
	for i := 0; i < b.N; i++ {
		_ = Problem9()
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkProblem9(b=%v)", b)
}
