package problem18
import logritelog "log"

import "testing"
import stdlog "log"
var log = stdlog.Default()

// Tests
func TestProblem18_Func(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestProblem18_Func(t=%v)", t)
	tests := []struct {
		name     string
		input    []string
		deep     int
		expected int
	}{
		{
			name:     "Test case 1",
			input:    problem18_test_parsed_string,
			deep:     2,
			expected: 23,
		},
		{
			name:     "Test case 2",
			input:    problem18_input_parsed_string,
			deep:     2,
			expected: 1074,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Problem18(test.input, test.deep)
			if actual != test.expected {
				t.Errorf("Expected %d, but got %d", test.expected, actual)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestProblem18_Func(t=%v)", t)
}

// Benchmarks
func BenchmarkProblem18_Func(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkProblem18_Func(b=%v)", b)
	for i := 0; i < b.N; i++ {
		Problem18(problem18_input_parsed_string, 2)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkProblem18_Func(b=%v)", b)
}
