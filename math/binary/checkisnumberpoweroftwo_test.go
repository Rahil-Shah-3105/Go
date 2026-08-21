// checkisnumberpoweroftwo_test.go
// description: Test for Is the number a power of two
// author(s) [red_byte](https://github.com/i-redbyte)
// see checkisnumberpoweroftwo.go

package binary
import logritelog "log"

import (
	math2 "github.com/TheAlgorithms/Go/math"
	"math"
	"testing"
)

func getTestsForPowerOfTwo() []struct {
	name    string
	a       int
	missing bool
} {
	var tests = []struct {
		name    string
		a       int
		missing bool
	}{
		{"Is 64 a power of 2? - YES", 64, true},
		{"Is 1 a power of 2? - YES", 1, true},
		{"Is 2 a power of 2? - YES", 2, true},
		{"Is 5 a power of 2? - NO", 5, false},
		{"Is 1023 a power of 2? - NO", 1023, false},
		{"Is 1024 a power of 2? - YES", 1024, true},
		{"Is 0 a power of 2? - NO", 0, false},
		{"Is 9223372036854775807 a power of 2? - NO", math.MaxInt64, false},
		{"Is 9223372036854775806 a power of 2? - NO", math.MaxInt64, false},
		{"Is 4611686018427387904 a power of 2? - YES", 4611686018427387904, true},
	}
	return tests
}

func TestIsPowerOfTwo(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestIsPowerOfTwo(t=%v)", t)
	tests := getTestsForPowerOfTwo()
	for _, tv := range tests {
		t.Run(tv.name, func(t *testing.T) {
			result := IsPowerOfTwo(tv.a)
			t.Log(tv.a, " ", result)
			if result != tv.missing {
				t.Errorf("Wrong result! Expected:%v, returned:%v ", tv.missing, result)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestIsPowerOfTwo(t=%v)", t)
}

func TestIsPowerOfTwoLeftShift(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestIsPowerOfTwoLeftShift(t=%v)", t)
	tests := getTestsForPowerOfTwo()
	for _, tv := range tests {
		t.Run(tv.name, func(t *testing.T) {
			result := IsPowerOfTwoLeftShift(uint(tv.a))
			t.Log(tv.a, " ", result)
			if result != tv.missing {
				t.Errorf("Wrong result! Expected:%v, returned:%v ", tv.missing, result)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestIsPowerOfTwoLeftShift(t=%v)", t)
}

func TestIsPowOfTwoUseLog(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestIsPowOfTwoUseLog(t=%v)", t)
	tests := getTestsForPowerOfTwo()
	for _, tv := range tests {
		t.Run(tv.name, func(t *testing.T) {
			result := math2.IsPowOfTwoUseLog(float64(tv.a))
			t.Log(tv.a, " ", result)
			if result != tv.missing {
				t.Errorf("Wrong result! Expected:%v, returned:%v ", tv.missing, result)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestIsPowOfTwoUseLog(t=%v)", t)
}

func BenchmarkIsPowerOfTwoBinaryMethod(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkIsPowerOfTwoBinaryMethod(b=%v)", b)
	for i := 0; i < b.N; i++ {
		IsPowerOfTwo(1024)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkIsPowerOfTwoBinaryMethod(b=%v)", b)
}

func BenchmarkIsPowerOfTwoUseCycleAndLeftShift(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkIsPowerOfTwoUseCycleAndLeftShift(b=%v)", b)
	for i := 0; i < b.N; i++ {
		IsPowerOfTwoLeftShift(1024)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkIsPowerOfTwoUseCycleAndLeftShift(b=%v)", b)
}

func BenchmarkIsPowerOfTwoUseLog(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkIsPowerOfTwoUseLog(b=%v)", b)
	for i := 0; i < b.N; i++ {
		math2.IsPowOfTwoUseLog(1024)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkIsPowerOfTwoUseLog(b=%v)", b)
}
