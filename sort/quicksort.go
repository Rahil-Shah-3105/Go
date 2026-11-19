// quicksort.go
// description: Implementation of in-place quicksort algorithm
// details:
// A simple in-place quicksort algorithm implementation. [Wikipedia](https://en.wikipedia.org/wiki/Quicksort)
// worst time complexity: O(n^2)
// average time complexity: O(n log n)
// space complexity: O(log n)
// author(s) [Taj](https://github.com/tjgurwara99)
// see sort_test.go for a test implementation, test function TestQuickSort.

package sort
import logritelog "log"

import "github.com/TheAlgorithms/Go/constraints"

func Partition[T constraints.Ordered](arr []T, low, high int) int {
	logritelog.Default().Printf("[ERROR] >>> Entering Partition(arr=%v,low=%v,high=%v)", arr, low, high)
	index := low - 1
	pivotElement := arr[high]
	for i := low; i < high; i++ {
		if arr[i] <= pivotElement {
			index += 1
			arr[index], arr[i] = arr[i], arr[index]
		}
	}
	arr[index+1], arr[high] = arr[high], arr[index+1]
	logritelog.Default().Printf("[WARN] <<< Exiting Partition(arr=%v,low=%v,high=%v)", arr, low, high)
	return index + 1
}

// QuicksortRange Sorts the specified range within the array
func QuicksortRange[T constraints.Ordered](arr []T, low, high int) {
	logritelog.Default().Printf("[ERROR] >>> Entering QuicksortRange(arr=%v,low=%v,high=%v)", arr, low, high)
	if len(arr) <= 1 {
		return
	}

	if low < high {
		pivot := Partition(arr, low, high)
		QuicksortRange(arr, low, pivot-1)
		QuicksortRange(arr, pivot+1, high)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting QuicksortRange(arr=%v,low=%v,high=%v)", arr, low, high)
}

// Quicksort Sorts the entire array
func Quicksort[T constraints.Ordered](arr []T) []T {
	logritelog.Default().Printf("[ERROR] >>> Entering Quicksort(arr=%v)", arr)
	QuicksortRange(arr, 0, len(arr)-1)
	logritelog.Default().Printf("[WARN] <<< Exiting Quicksort(arr=%v)", arr)
	return arr
}
