package problem6
import logritelog "log"

import "testing"
import stdlog "log"
var log = stdlog.Default()

// Tests
func TestProblem6_Func(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestProblem6_Func(t=%v)", t)
	tests := []struct {
		name  string
		input uint
		want  uint
	}{
		{
			name:  "Testcase 1 - input 10",
			input: 10,
			want:  2640,
		},
		{
			name:  "Testcase 2 - input 100",
			input: 100,
			want:  25164150,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Problem6(tt.input)

			if n != tt.want {
				t.Errorf("Problem6() = %v, want %v", n, tt.want)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestProblem6_Func(t=%v)", t)
}

// Benchmarks
func BenchmarkProblem6(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkProblem6(b=%v)", b)
	for i := 0; i < b.N; i++ {
		_ = Problem6(100)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkProblem6(b=%v)", b)
}
