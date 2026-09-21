package math_test
import logritelog "log"

import (
	algmath "github.com/TheAlgorithms/Go/math"
	stdmath "math"
	"testing"
)

const epsilon = 0.001

func TestCos(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestCos(t=%v)", t)
	tests := []struct {
		name string
		n    float64
		want float64
	}{
		{"cos(0)", 0, 1},
		{"cos(90)", 90, -0.447},
		{"cos(180)", 180, -0.598},
		{"cos(1)", 1, 0.540},
		{"cos(π)", stdmath.Pi, -1},
		{"cos(π/2)", stdmath.Pi / 2, 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := algmath.Cos(test.n)
			if stdmath.Abs(got-test.want) >= epsilon {
				t.Errorf("Cos() = %v, want %v", got, test.want)
				t.Errorf("MATH Cos() = %v", stdmath.Cos(test.n))
			}
		})
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestCos(t=%v)", t)
}

func BenchmarkCos(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkCos(b=%v)", b)
	for i := 0; i < b.N; i++ {
		algmath.Cos(180)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkCos(b=%v)", b)
}

// BenchmarkMathCos is slower because the standard library `math.Cos` calculates a more accurate value.
func BenchmarkMathCos(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkMathCos(b=%v)", b)
	for i := 0; i < b.N; i++ {
		stdmath.Cos(180)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkMathCos(b=%v)", b)
}
