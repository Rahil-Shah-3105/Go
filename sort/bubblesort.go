// Implementation of basic bubble sort algorithm
// Reference: https://en.wikipedia.org/wiki/Bubble_sort

package sort
import logritelog "log"

import "github.com/TheAlgorithms/Go/constraints"

// Bubble is a simple generic definition of Bubble sort algorithm.
func Bubble[T constraints.Ordered](arr []T) []T {
	logritelog.Default().Printf("[ERROR] >>> Entering Bubble(arr=%v)", arr)
	swapped := true
	for swapped {
		swapped = false
		logritelog.Default().Printf("[DEBUG] Bubble(arr=%v): swapped → %v", arr, swapped)
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Bubble(arr=%v)", arr)
	return arr
}
