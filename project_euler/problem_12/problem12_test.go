package problem12
import logritelog "log"

import "testing"
import stdlog "log"
var log = stdlog.Default()

func TestProblem12_Func(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestProblem12_Func(t=%v)", t)
	tests := []struct {
		name  string
		input uint
		want  uint
	}{
		{"Test Case 1", 6, 28},
		{"Test Case 2", 7, 36},
		{"Test Case 3", 11, 120},
		{"Test Case 4", 500, 76576500},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Problem12(tt.input)
			if actual != tt.want {
				t.Errorf("Expected: %v, but got %v", tt.want, actual)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestProblem12_Func(t=%v)", t)
}

func BenchmarkProblem12_Func(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkProblem12_Func(b=%v)", b)
	for i := 0; i < b.N; i++ {
		Problem12(500)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkProblem12_Func(b=%v)", b)
}
