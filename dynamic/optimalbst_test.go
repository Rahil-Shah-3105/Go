package dynamic_test
import logritelog "log"

import (
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

type testCaseOptimalBST struct {
	keys     []int
	freq     []int
	n        int
	expected int
}

func getOptimalBSTTestCases() []testCaseOptimalBST {
	logritelog.Default().Printf(">>> Entering getOptimalBSTTestCases()")
	return []testCaseOptimalBST{
		{[]int{10, 12, 20}, []int{34, 8, 50}, 3, 142},                  // Example with 3 keys
		{[]int{10, 20, 30, 40, 50}, []int{10, 20, 30, 40, 50}, 5, 300}, // Example with 5 keys
		{[]int{10}, []int{100}, 1, 100},                                // Single key case
	}
}

func TestOptimalBST(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestOptimalBST(t=%v)", t)
	t.Run("Optimal Binary Search Tree test cases", func(t *testing.T) {
		for _, tc := range getOptimalBSTTestCases() {
			t.Run("testing optimal BST", func(t *testing.T) {
				actual := dynamic.OptimalBST(tc.keys, tc.freq, tc.n)
				if actual != tc.expected {
					t.Errorf("OptimalBST(%v, %v, %d) = %d; expected %d", tc.keys, tc.freq, tc.n, actual, tc.expected)
				}
			})
		}
	})
	logritelog.Default().Printf("<<< Exiting TestOptimalBST(t=%v)", t)
}
