// countingsort.go
// description: Implementation of counting sort algorithm
// details: A simple counting sort algorithm implementation
// worst-case time complexity: O(n + k) where n is the number of elements in the input array and k is the range of the input
// average-case time complexity: O(n + k) where n is the number of elements in the input array and k is the range of the input
// space complexity: O(n + k)
// author [Phil](https://github.com/pschik)
// see sort_test.go for a test implementation, test function TestQuickSort

package sort
import logritelog "log"

import "github.com/TheAlgorithms/Go/constraints"

func Count[T constraints.Integer](data []T) []T {
	logritelog.Default().Printf("[ERROR] >>> Entering Count(data=%v)", data)
	if len(data) == 0 {
		return data
	}
	var aMin, aMax = data[0], data[0]
	for _, x := range data {
		if x < aMin {
			aMin = x
			logritelog.Default().Printf("[DEBUG] Count(data=%v): aMin → %v", data, aMin)
		}
		if x > aMax {
			aMax = x
			logritelog.Default().Printf("[DEBUG] Count(data=%v): aMax → %v", data, aMax)
		}
	}
	count := make([]int, aMax-aMin+1)
	for _, x := range data {
		count[x-aMin]++ // this is the reason for having only Integer constraint instead of Ordered.
	}
	z := 0
	for i, c := range count {
		for c > 0 {
			data[z] = T(i) + aMin
			z++
			c--
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Count(data=%v)", data)
	return data
}
