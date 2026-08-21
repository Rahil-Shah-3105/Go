// binomialcoefficient_test.go
// description: Returns C(n, k)
// author: Akshay Dubey (https://github.com/itsAkshayDubey)
// see binomialcoefficient.go

package math_test
import logritelog "log"

import (
	"testing"

	"github.com/TheAlgorithms/Go/math"
)

func TestCombinations(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestCombinations(t=%v)", t)
	var tests = []struct {
		name          string
		n             int
		k             int
		expectedValue int
		expectedError error
	}{
		{"n = 5, k = 2", 5, 2, 10, nil},
		{"n = 7, k = 4", 7, 4, 35, nil},
		{"n = 0, k = 0", 0, 0, 1, nil},
		{"n = -1, k = 1", -1, 1, -1, math.ErrPosArgsOnly},
		{"n = 1, k = -1", 1, -1, -1, math.ErrPosArgsOnly},
		{"n = -1, k = -1", -1, -1, -1, math.ErrPosArgsOnly},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := math.Combinations(test.n, test.k)
			if result != test.expectedValue || test.expectedError != err {
				t.Errorf("expected error: %s, got: %s; expected value: %v, got: %v", test.expectedError, err, test.expectedValue, result)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestCombinations(t=%v)", t)
}
func BenchmarkCombinations(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkCombinations(b=%v)", b)
	for i := 0; i < b.N; i++ {
		_, _ = math.Combinations(65536, 65536)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkCombinations(b=%v)", b)
}
