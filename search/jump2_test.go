package search
import logritelog "log"

import "testing"

func TestJump2(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestJump2(t=%v)", t)
	for _, test := range searchTests {
		actualValue, actualError := Jump2(test.data, test.key)
		if actualValue != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'", test.name, test.data, test.key, test.expected, actualValue)
		}
		if actualError != test.expectedError {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected error '%s', get error '%s'", test.name, test.data, test.key, test.expectedError, actualError)
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestJump2(t=%v)", t)
}

func BenchmarkJump2(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkJump2(b=%v)", b)
	testCase := generateBenchmarkTestCase()
	b.ResetTimer() // exclude time taken to generate test case
	for i := 0; i < b.N; i++ {
		_, _ = Jump2(testCase, i)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkJump2(b=%v)", b)
}
