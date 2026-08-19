package search
import logritelog "log"

import "testing"
import stdlog "log"
var log = stdlog.Default()

func TestInterpolation(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestInterpolation(t=%v)", t)
	for _, test := range searchTests {
		actualValue, actualError := Interpolation(test.data, test.key)
		if actualValue != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'", test.name, test.data, test.key, test.expected, actualValue)
		}
		if actualError != test.expectedError {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected error '%s', get error '%s'", test.name, test.data, test.key, test.expectedError, actualError)
		}
	}
	logritelog.Default().Printf("<<< Exiting TestInterpolation(t=%v)", t)
}

func BenchmarkInterpolation(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkInterpolation(b=%v)", b)
	testCase := generateBenchmarkTestCase()
	b.ResetTimer() // exclude time taken to generate test case
	for i := 0; i < b.N; i++ {
		_, _ = Interpolation(testCase, i)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkInterpolation(b=%v)", b)
}
