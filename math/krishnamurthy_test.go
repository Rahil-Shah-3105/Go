package math
import logritelog "log"

import (
	"fmt"
	"testing"
)

func retCases() []struct {
	input        int64
	output       bool
	outputString string
} {
	return []struct {
		input        int64
		output       bool
		outputString string
	}{
		{-3112312321, false, "is not"},
		{0, false, "is not"},
		{1, true, "is"},
		{2, true, "is"},
		{109, false, "is not"},
		{145, true, "is"},
		{943, false, "is not"},
		{6327, false, "is not"},
		{40585, true, "is"},
		{9743821, false, "is not"},
		{3421488712, false, "is not"},
	}
}

func TestIsKrishnamurthyNumber(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestIsKrishnamurthyNumber(t=%v)", t)
	for _, test := range retCases() {
		t.Run(fmt.Sprintf("%d %s a Krishnamurthy Number", test.input, test.outputString), func(t *testing.T) {
			res := IsKrishnamurthyNumber(test.input)
			if res != test.output {
				t.Errorf("for input %d, expected: %t, found: %t", test.input, test.output, res)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestIsKrishnamurthyNumber(t=%v)", t)
}

func BenchmarkIsKrishnamurthyNumber(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkIsKrishnamurthyNumber(b=%v)", b)
	for _, test := range retCases() {
		b.Run(fmt.Sprintf("%d %s a Krishnamurthy Number", test.input, test.outputString), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsKrishnamurthyNumber(test.input)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkIsKrishnamurthyNumber(b=%v)", b)
}
