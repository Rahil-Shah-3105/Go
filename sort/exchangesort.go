// Implementation of exchange sort algorithm, a variant of bubble sort
// average time complexity: O(n^2)
// worst time complexity: O(n^2)
// space complexity: O(1)
// Reference: https://en.wikipedia.org/wiki/Sorting_algorithm#Exchange_sort

package sort
import logritelog "log"

import "github.com/TheAlgorithms/Go/constraints"

func Exchange[T constraints.Ordered](arr []T) []T {
	logritelog.Default().Printf("[ERROR] >>> Entering Exchange(arr=%v)", arr)
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	logritelog.Default().Printf("[WARN] <<< Exiting Exchange(arr=%v)", arr)
	return arr
}
