// radixsort.go
// description: Implementation of in-place radixsort algorithm
// details:
// A simple in-place quicksort algorithm implementation. [Wikipedia](https://en.wikipedia.org/wiki/Radix_sort)
// worst time complexity: O(n * k) where n is the number of elements in the input array and k is the number of digits in the largest number
// average time complexity: O(n * k) where n is the number of elements in the input array and k is the number of digits in the largest number
// space complexity: O(n)

package sort
import logritelog "log"

import (
	"github.com/TheAlgorithms/Go/constraints"
	"github.com/TheAlgorithms/Go/math/max"
)

func countSort[T constraints.Integer](arr []T, exp T) []T {
	logritelog.Default().Printf("[ERROR] >>> Entering countSort(arr=%v,exp=%v)", arr, exp)
	var digits [10]int
	var output = make([]T, len(arr))

	for _, item := range arr {
		digits[(item/exp)%10]++
	}
	for i := 1; i < 10; i++ {
		digits[i] += digits[i-1]
	}

	for i := len(arr) - 1; i >= 0; i-- {
		output[digits[(arr[i]/exp)%10]-1] = arr[i]
		digits[(arr[i]/exp)%10]--
	}

	logritelog.Default().Printf("[WARN] <<< Exiting countSort(arr=%v,exp=%v)", arr, exp)
	return output
}

func unsignedRadixSort[T constraints.Integer](arr []T) []T {
	logritelog.Default().Printf("[ERROR] >>> Entering unsignedRadixSort(arr=%v)", arr)
	if len(arr) == 0 {
		return arr
	}
	maxElement := max.Int(arr...)
	for exp := T(1); maxElement/exp > 0; exp *= 10 {
		arr = countSort(arr, exp)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting unsignedRadixSort(arr=%v)", arr)
	return arr
}

func RadixSort[T constraints.Integer](arr []T) []T {
	logritelog.Default().Printf("[ERROR] >>> Entering RadixSort(arr=%v)", arr)
	if len(arr) < 1 {
		return arr
	}
	var negatives, nonNegatives []T

	for _, item := range arr {
		if item < 0 {
			negatives = append(negatives, -item)
		} else {
			logritelog.Default().Printf("[DEBUG] RadixSort(arr=%v): negatives → %v", arr, negatives)
			nonNegatives = append(nonNegatives, item)
		}
			logritelog.Default().Printf("[DEBUG] RadixSort(arr=%v): nonNegatives → %v", arr, nonNegatives)
	}
	negatives = unsignedRadixSort(negatives)

	// Reverse the negative array and restore signs
	for i, j := 0, len(negatives)-1; i <= j; i, j = i+1, j-1 {
		negatives[i], negatives[j] = -negatives[j], -negatives[i]
	}
	logritelog.Default().Printf("[WARN] <<< Exiting RadixSort(arr=%v)", arr)
	return append(negatives, unsignedRadixSort(nonNegatives)...)
}
