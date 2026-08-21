package power
import logritelog "log"

import "testing"

var testCases = []struct {
	name     string
	base     uint
	power    uint
	expected uint
}{
	{"0^2", 0, 2, 0},
	{"2^0", 2, 0, 1},
	{"2^3", 2, 3, 8},
	{"8^3", 8, 3, 512},
	{"10^5", 10, 5, 100000},
}

func TestIterativePower(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestIterativePower(t=%v)", t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := IterativePower(tc.base, tc.power)
			if actual != tc.expected {
				t.Errorf("Expected %d to the power of %d to be: %d, but got: %d", tc.base, tc.power, tc.expected, actual)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestIterativePower(t=%v)", t)
}

func TestRecursivePower(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestRecursivePower(t=%v)", t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := RecursivePower(tc.base, tc.power)
			if actual != tc.expected {
				t.Errorf("Expected %d to the power of %d to be: %d, but got: %d", tc.base, tc.power, tc.expected, actual)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestRecursivePower(t=%v)", t)
}

func TestRecursivePower1(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestRecursivePower1(t=%v)", t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := RecursivePower1(tc.base, tc.power)
			if actual != tc.expected {
				t.Errorf("Expected %d to the power of %d to be: %d, but got: %d", tc.base, tc.power, tc.expected, actual)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestRecursivePower1(t=%v)", t)
}

func BenchmarkIterativePower(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkIterativePower(b=%v)", b)
	for i := 0; i < b.N; i++ {
		IterativePower(10, 5)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkIterativePower(b=%v)", b)
}

func BenchmarkRecursivePower(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkRecursivePower(b=%v)", b)
	for i := 0; i < b.N; i++ {
		RecursivePower(10, 5)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkRecursivePower(b=%v)", b)
}

func BenchmarkRecursivePower1(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkRecursivePower1(b=%v)", b)
	for i := 0; i < b.N; i++ {
		RecursivePower1(10, 5)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkRecursivePower1(b=%v)", b)
}
