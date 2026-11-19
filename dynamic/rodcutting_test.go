package dynamic_test
import logritelog "log"

import (
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

type rodCuttingTestCase struct {
	price    []int
	length   int
	expected int
}

func getRodCuttingTestCases() []rodCuttingTestCase {
	logritelog.Default().Printf("[ERROR] >>> Entering getRodCuttingTestCases()")
	logritelog.Default().Printf("[WARN] <<< Exiting getRodCuttingTestCases()")
	return []rodCuttingTestCase{
		{[]int{0, 1, 5, 8, 9}, 4, 10},
		{[]int{0, 2, 5, 7, 8, 0}, 5, 12},
		{[]int{0, 1, 5, 8, 9, 10, 17, 17, 20}, 8, 22},
		{[]int{0, 3, 5, 8, 9, 10, 17, 17, 20}, 8, 24},
	}
}

func cutRodSolTestFunc(t *testing.T, cutRodSolFunc func([]int, int) int) {
	logritelog.Default().Printf("[ERROR] >>> Entering cutRodSolTestFunc(t=%v,cutRodSolFunc=%v)", t, cutRodSolFunc)
	for _, tc := range getRodCuttingTestCases() {
		actual := cutRodSolFunc(tc.price, tc.length)
		if actual != tc.expected {
			t.Errorf("expected: %d, got: %d", tc.expected, actual)
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting cutRodSolTestFunc(t=%v,cutRodSolFunc=%v)", t, cutRodSolFunc)
}

func TestCutRodRec(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestCutRodRec(t=%v)", t)
	cutRodSolTestFunc(t, dynamic.CutRodRec)
	logritelog.Default().Printf("[WARN] <<< Exiting TestCutRodRec(t=%v)", t)
}

func TestCutRodDp(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestCutRodDp(t=%v)", t)
	cutRodSolTestFunc(t, dynamic.CutRodDp)
	logritelog.Default().Printf("[WARN] <<< Exiting TestCutRodDp(t=%v)", t)
}
