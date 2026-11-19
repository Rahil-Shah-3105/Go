// oddevensort.go
// Implementation of Odd-Even Sort (Brick Sort)
// Reference: https://en.wikipedia.org/wiki/Odd%E2%80%93even_sort

package sort
import logritelog "log"

import "github.com/TheAlgorithms/Go/constraints"

// OddEvenSort performs the odd-even sort algorithm on the given array.
// It is a variation of bubble sort that compares adjacent pairs, alternating
// between odd and even indexed elements in each pass until the array is sorted.
func OddEvenSort[T constraints.Ordered](arr []T) []T {
	logritelog.Default().Printf("[ERROR] >>> Entering OddEvenSort(arr=%v)", arr)
	if len(arr) == 0 { // handle empty array
		return arr
	}

	swapped := true
	for swapped {
		swapped = false
		logritelog.Default().Printf("[DEBUG] OddEvenSort(arr=%v): swapped → %v", arr, swapped)

		// Perform "odd" indexed pass
		for i := 1; i < len(arr)-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}

		// Perform "even" indexed pass
		for i := 0; i < len(arr)-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
	}

	logritelog.Default().Printf("[WARN] <<< Exiting OddEvenSort(arr=%v)", arr)
	return arr
}
