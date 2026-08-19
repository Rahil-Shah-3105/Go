package gcd
import logritelog "log"

import "testing"

type testFunction func(int64, int64) int64

var testCases = []struct {
	name   string
	a      int64
	b      int64
	output int64
}{
	{"gcd of 10 and 0", 10, 0, 10},
	{"gcd of 98 and 56", 98, 56, 14},
	{"gcd of 0 and 10", 0, 10, 10},
}

func TemplateTestGCD(t *testing.T, f testFunction) {
	logritelog.Default().Printf(">>> Entering TemplateTestGCD(t=%v,f=%v)", t, f)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := f(tc.a, tc.b)
			if actual != tc.output {
				t.Errorf("Expected GCD of %d and %d to be: %v, but got: %d", tc.a, tc.b, tc.output, actual)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TemplateTestGCD(t=%v,f=%v)", t, f)
}

func TestGCDRecursive(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestGCDRecursive(t=%v)", t)
	TemplateTestGCD(t, Recursive)
	logritelog.Default().Printf("<<< Exiting TestGCDRecursive(t=%v)", t)
}

func TestGCDIterative(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestGCDIterative(t=%v)", t)
	TemplateTestGCD(t, Iterative)
	logritelog.Default().Printf("<<< Exiting TestGCDIterative(t=%v)", t)
}

func TemplateBenchmarkGCD(b *testing.B, f testFunction) {
	logritelog.Default().Printf(">>> Entering TemplateBenchmarkGCD(b=%v,f=%v)", b, f)
	for i := 0; i < b.N; i++ {
		f(98, 56)
	}
	logritelog.Default().Printf("<<< Exiting TemplateBenchmarkGCD(b=%v,f=%v)", b, f)
}

func BenchmarkGCDRecursive(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkGCDRecursive(b=%v)", b)
	TemplateBenchmarkGCD(b, Recursive)
	logritelog.Default().Printf("<<< Exiting BenchmarkGCDRecursive(b=%v)", b)
}

func BenchmarkGCDIterative(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkGCDIterative(b=%v)", b)
	TemplateBenchmarkGCD(b, Iterative)
	logritelog.Default().Printf("<<< Exiting BenchmarkGCDIterative(b=%v)", b)
}
