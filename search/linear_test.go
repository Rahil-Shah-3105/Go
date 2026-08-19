package search
import logritelog "log"

import (
	"testing"
)

func TestLinear(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestLinear(t=%v)", t)
	for _, test := range searchTests {
		actualValue, actualError := Linear(test.data, test.key)
		if actualValue != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'", test.name, test.data, test.key, test.expected, actualValue)
		}
		if actualError != test.expectedError {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected error '%s', get error '%s'", test.name, test.data, test.key, test.expectedError, actualError)
		}
	}
	logritelog.Default().Printf("<<< Exiting TestLinear(t=%v)", t)
}

func BenchmarkLinear(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkLinear(b=%v)", b)
	testCase := generateBenchmarkTestCase()
	b.ResetTimer() // exclude time taken to generate test case
	for i := 0; i < b.N; i++ {
		_, _ = Linear(testCase, i)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkLinear(b=%v)", b)
}
