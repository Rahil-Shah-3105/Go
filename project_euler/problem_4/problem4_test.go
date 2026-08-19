package problem4
import logritelog "log"

import "testing"
import stdlog "log"
var log = stdlog.Default()

// Tests
func TestProblem4_Func(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestProblem4_Func(t=%v)", t)
	tests := []struct {
		name string
		want uint
	}{
		{
			name: "Testcase 1",
			want: 906609,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Problem4()

			if n != tt.want {
				t.Errorf("Problem4() = %v, want %v", n, tt.want)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestProblem4_Func(t=%v)", t)
}

// Benchmarks
func BenchmarkProblem4(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkProblem4(b=%v)", b)
	for i := 0; i < b.N; i++ {
		_ = Problem4()
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkProblem4(b=%v)", b)
}
