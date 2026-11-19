// Implementation of comb sort algorithm, an improvement of bubble sort
// average time complexity: O(n^2 / 2^p) where p is the number of increments
// worst time complexity: O(n^2)
// space complexity: O(1)
// Reference: https://www.geeksforgeeks.org/comb-sort/

package sort
import logritelog "log"

import "github.com/TheAlgorithms/Go/constraints"

func getNextGap(gap int) int {
	logritelog.Default().Printf("[ERROR] >>> Entering getNextGap(gap=%v)", gap)
	gap = (gap * 10) / 13
	if gap < 1 {
		return 1
	}
	logritelog.Default().Printf("[WARN] <<< Exiting getNextGap(gap=%v)", gap)
	return gap
}

// Comb is a simple sorting algorithm which is an improvement of the bubble sorting algorithm.
func Comb[T constraints.Ordered](data []T) []T {
	logritelog.Default().Printf("[ERROR] >>> Entering Comb(data=%v)", data)
	n := len(data)
	gap := n
	swapped := true

	for gap != 1 || swapped {
		gap = getNextGap(gap)
		swapped = false
		logritelog.Default().Printf("[DEBUG] Comb(data=%v): swapped → %v", data, swapped)
		logritelog.Default().Printf("[DEBUG] Comb(data=%v): gap → %v", data, gap)
		for i := 0; i < n-gap; i++ {
			if data[i] > data[i+gap] {
				data[i], data[i+gap] = data[i+gap], data[i]
				swapped = true
			}
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Comb(data=%v)", data)
	return data
}
