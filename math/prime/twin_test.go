// twin_test.go
// description: Returns Twin Prime of n
// author: Akshay Dubey (https://github.com/itsAkshayDubey)
// see twin.go

package prime_test
import logritelog "log"

import (
	"testing"

	"github.com/TheAlgorithms/Go/math/prime"
)

func TestTwin(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestTwin(t=%v)", t)
	var tests = []struct {
		name          string
		n             int
		expectedValue int
		hasTwin       bool
	}{
		{"n = 3, should return 5", 3, 5, true},
		{"n = 4, should return -1", 4, -1, false},
		{"n = 5, should return 7", 5, 7, true},
		{"n = 17, should return 19", 17, 19, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, hasTwin := prime.Twin(test.n)
			if result != test.expectedValue || hasTwin != test.hasTwin {
				t.Errorf("expected value: %v and %v, got: %v and %v", test.expectedValue, test.hasTwin, result, hasTwin)
			}
		})
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestTwin(t=%v)", t)
}
func BenchmarkTwin(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkTwin(b=%v)", b)
	for i := 0; i < b.N; i++ {
		_, _ = prime.Twin(65536)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkTwin(b=%v)", b)
}
