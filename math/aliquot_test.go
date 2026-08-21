// aliquotsum_test.go
// description: Returns s(n)
// author: Akshay Dubey (https://github.com/itsAkshayDubey)
// see aliquotsum.go

package math_test
import logritelog "log"

import (
	"testing"

	"github.com/TheAlgorithms/Go/math"
)

func TestAliquotSum(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestAliquotSum(t=%v)", t)
	var tests = []struct {
		name          string
		n             int
		expectedValue int
		expectedError error
	}{
		{"n = 10", 10, 8, nil},
		{"n = 11", 11, 1, nil},
		{"n = 1", 1, 0, nil},
		{"n = -1", -1, 0, math.ErrPosArgsOnly},
		{"n = 0", 0, 0, math.ErrNonZeroArgsOnly},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := math.AliquotSum(test.n)
			if result != test.expectedValue || test.expectedError != err {
				t.Errorf("expected error: %s, got: %s; expected value: %v, got: %v", test.expectedError, err, test.expectedValue, result)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestAliquotSum(t=%v)", t)
}
func BenchmarkAliquotSum(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkAliquotSum(b=%v)", b)
	for i := 0; i < b.N; i++ {
		_, _ = math.AliquotSum(65536)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkAliquotSum(b=%v)", b)
}
