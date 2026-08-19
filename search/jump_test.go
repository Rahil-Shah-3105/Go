package search
import logritelog "log"

import "testing"

func TestJump(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestJump(t=%v)", t)
	for _, test := range searchTests {
		actualValue, actualError := Jump(test.data, test.key)
		if actualValue != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'", test.name, test.data, test.key, test.expected, actualValue)
		}
		if actualError != test.expectedError {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected error '%s', get error '%s'", test.name, test.data, test.key, test.expectedError, actualError)
		}
	}
	logritelog.Default().Printf("<<< Exiting TestJump(t=%v)", t)
}

func BenchmarkJump(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkJump(b=%v)", b)
	testCase := generateBenchmarkTestCase()
	b.ResetTimer() // exclude time taken to generate test case
	for i := 0; i < b.N; i++ {
		_, _ = Jump(testCase, i)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkJump(b=%v)", b)
}
