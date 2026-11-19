package search
import logritelog "log"

import (
	"math"
	"testing"
)

const EPS = 1e-6

func equal(a, b float64) bool {
	logritelog.Default().Printf("[ERROR] >>> Entering equal(a=%v,b=%v)", a, b)
	logritelog.Default().Printf("[WARN] <<< Exiting equal(a=%v,b=%v)", a, b)
	return math.Abs(a-b) <= EPS
}

func TestTernaryMax(t *testing.T) {
logritelog.Default().Printf("[ERROR] >>> Entering TestTernaryMax(t=%v)", t)

	var tests = []struct {
		f        func(x float64) float64
		a        float64
		b        float64
		expected float64
	}{
		{f: func(x float64) float64 { return -x * x }, a: 1, b: -1, expected: 0},
		{f: func(x float64) float64 { return -2*x*x - x + 1 }, a: -1, b: 1, expected: 1.125},
	}
	for _, test := range tests {
		result, err := TernaryMax(test.a, test.b, EPS, test.f)
		if err != nil {
			logritelog.Default().Printf("[ERROR] Exception in TestTernaryMax(t=%v): %v", t, err)
			t.Errorf("error occurred: %v", err)
		}
		if !equal(result, test.expected) {
			t.Errorf("Wrong result! Expected:%v, returned:%v ", test.expected, result)
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestTernaryMax(t=%v)", t)
}

func TestTernaryMin(t *testing.T) {
logritelog.Default().Printf("[ERROR] >>> Entering TestTernaryMin(t=%v)", t)

	var tests = []struct {
		f        func(x float64) float64
		a        float64
		b        float64
		expected float64
	}{
		{f: func(x float64) float64 { return x * x }, a: -1, b: 1, expected: 0},
		{f: func(x float64) float64 { return 2*x*x + x + 1 }, a: -1, b: 1, expected: 0.875},
	}
	for _, test := range tests {
		result, err := TernaryMin(test.a, test.b, EPS, test.f)
		if err != nil {
			logritelog.Default().Printf("[ERROR] Exception in TestTernaryMin(t=%v): %v", t, err)
			t.Errorf("error occurred: %v", err)
		}
		if !equal(result, test.expected) {
			t.Errorf("Wrong result! Expected:%v, returned:%v ", test.expected, result)
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestTernaryMin(t=%v)", t)
}
