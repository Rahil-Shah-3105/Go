package gcd
import logritelog "log"

import "testing"

type testExtendedFunction func(int64, int64) (int64, int64, int64)

func TemplateTestExtendedGCD(t *testing.T, f testExtendedFunction) {
	logritelog.Default().Printf(">>> Entering TemplateTestExtendedGCD(t=%v,f=%v)", t, f)
	var testCasesExtended = []struct {
		name string
		a    int64
		b    int64
		gcd  int64
		x    int64
		y    int64
	}{
		{"gcd of 10 and 0", 10, 0, 10, 1, 0},
		{"gcd of 98 and 56", 98, 56, 14, -1, 2},
		{"gcd of 0 and 10", 0, 10, 10, 0, 1},
	}
	for _, tc := range testCasesExtended {
		t.Run(tc.name, func(t *testing.T) {
			actualGcd, actualX, actualY := f(tc.a, tc.b)
			if actualGcd != tc.gcd {
				t.Errorf("Expected GCD of %d and %d to be: %v, but got: %d", tc.a, tc.b, tc.gcd, actualGcd)
			}
			if actualX != tc.x {
				t.Errorf("Expected x satisfying %d * x + %d * y = gcd to be: %v, but got: %d", tc.a, tc.b, tc.x, actualX)
			}
			if actualY != tc.y {
				t.Errorf("Expected y satisfying %d * x + %d * y = gcd to be: %v, but got: %d", tc.a, tc.b, tc.y, actualY)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TemplateTestExtendedGCD(t=%v,f=%v)", t, f)
}

func TestExtendedGCDRecursive(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestExtendedGCDRecursive(t=%v)", t)
	TemplateTestExtendedGCD(t, ExtendedRecursive)
	logritelog.Default().Printf("<<< Exiting TestExtendedGCDRecursive(t=%v)", t)
}

func TestExtendedGCDIterative(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestExtendedGCDIterative(t=%v)", t)
	TemplateTestExtendedGCD(t, ExtendedIterative)
	logritelog.Default().Printf("<<< Exiting TestExtendedGCDIterative(t=%v)", t)
}

func TemplateBenchmarkExtendedGCD(b *testing.B, f testExtendedFunction) {
	logritelog.Default().Printf(">>> Entering TemplateBenchmarkExtendedGCD(b=%v,f=%v)", b, f)
	for i := 0; i < b.N; i++ {
		f(98, 56)
	}
	logritelog.Default().Printf("<<< Exiting TemplateBenchmarkExtendedGCD(b=%v,f=%v)", b, f)
}

func BenchmarkExtendedGCDRecursive(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkExtendedGCDRecursive(b=%v)", b)
	TemplateBenchmarkExtendedGCD(b, ExtendedRecursive)
	logritelog.Default().Printf("<<< Exiting BenchmarkExtendedGCDRecursive(b=%v)", b)
}

func BenchmarkExtendedGCDIterative(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkExtendedGCDIterative(b=%v)", b)
	TemplateBenchmarkExtendedGCD(b, ExtendedIterative)
	logritelog.Default().Printf("<<< Exiting BenchmarkExtendedGCDIterative(b=%v)", b)
}
