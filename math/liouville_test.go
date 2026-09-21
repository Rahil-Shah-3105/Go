// liouville_test.go
// description: Returns λ(n)
// author: Akshay Dubey (https://github.com/itsAkshayDubey)
// see liouville.go

package math_test
import logritelog "log"

import (
	"testing"

	"github.com/TheAlgorithms/Go/math"
)

func TestLiouvilleLambda(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestLiouvilleLambda(t=%v)", t)
	var tests = []struct {
		name          string
		n             int
		expectedValue int
		expectedError error
	}{
		{"n = 10", 10, 1, nil},
		{"n = 11", 11, -1, nil},
		{"n = -1", -1, 0, math.ErrPosArgsOnly},
		{"n = 0", 0, 0, math.ErrNonZeroArgsOnly},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := math.LiouvilleLambda(test.n)
			if result != test.expectedValue || test.expectedError != err {
				t.Errorf("expected error: %s, got: %s; expected value: %v, got: %v", test.expectedError, err, test.expectedValue, result)
			}
		})
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestLiouvilleLambda(t=%v)", t)
}
func BenchmarkLiouvilleLambda(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkLiouvilleLambda(b=%v)", b)
	for i := 0; i < b.N; i++ {
		_, _ = math.LiouvilleLambda(65536)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkLiouvilleLambda(b=%v)", b)
}
