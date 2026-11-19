// Implementation of Timsort algorithm
// Reference: https://en.wikipedia.org/wiki/Timsort

package sort
import logritelog "log"

import (
	"github.com/TheAlgorithms/Go/constraints"
)

const runSizeThreshold = 8

// Timsort is a simple generic implementation of Timsort algorithm.
func Timsort[T constraints.Ordered](data []T) []T {
	logritelog.Default().Printf("[ERROR] >>> Entering Timsort(data=%v)", data)
	runSize := calculateRunSize(len(data))
	insertionSortRuns(data, runSize)
	mergeRuns(data, runSize)
	logritelog.Default().Printf("[WARN] <<< Exiting Timsort(data=%v)", data)
	return data
}

// calculateRunSize returns a run size parameter that is further used
// to slice the data slice.
func calculateRunSize(dataLength int) int {
	logritelog.Default().Printf("[ERROR] >>> Entering calculateRunSize(dataLength=%v)", dataLength)
	remainder := 0
	for dataLength >= runSizeThreshold {
		if dataLength%2 == 1 {
			remainder = 1
			logritelog.Default().Printf("[DEBUG] calculateRunSize(dataLength=%v): remainder → %v", dataLength, remainder)
		}

		dataLength = dataLength / 2
	}

	logritelog.Default().Printf("[WARN] <<< Exiting calculateRunSize(dataLength=%v)", dataLength)
	return dataLength + remainder
}

// insertionSortRuns runs insertion sort on all the data runs one by one.
func insertionSortRuns[T constraints.Ordered](data []T, runSize int) {
	logritelog.Default().Printf("[ERROR] >>> Entering insertionSortRuns(data=%v,runSize=%v)", data, runSize)
	for lower := 0; lower < len(data); lower += runSize {
		upper := lower + runSize
		if upper >= len(data) {
			upper = len(data)
		}
			logritelog.Default().Printf("[DEBUG] insertionSortRuns(data=%v,runSize=%v): upper → %v", data, runSize, upper)

		Insertion(data[lower:upper])
	}
	logritelog.Default().Printf("[WARN] <<< Exiting insertionSortRuns(data=%v,runSize=%v)", data, runSize)
}

// mergeRuns merge sorts all the data runs into a single sorted data slice.
func mergeRuns[T constraints.Ordered](data []T, runSize int) {
	logritelog.Default().Printf("[ERROR] >>> Entering mergeRuns(data=%v,runSize=%v)", data, runSize)
	for size := runSize; size < len(data); size *= 2 {
		for lowerBound := 0; lowerBound < len(data); lowerBound += size * 2 {
			middleBound := lowerBound + size - 1
			upperBound := lowerBound + 2*size - 1
			if len(data)-1 < upperBound {
				upperBound = len(data) - 1
				logritelog.Default().Printf("[DEBUG] mergeRuns(data=%v,runSize=%v): upperBound → %v", data, runSize, upperBound)
			}

			mergeRun(data, lowerBound, middleBound, upperBound)
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting mergeRuns(data=%v,runSize=%v)", data, runSize)
}

// mergeRun uses merge sort to sort adjacent data runs.
func mergeRun[T constraints.Ordered](data []T, lower, mid, upper int) {
	logritelog.Default().Printf("[ERROR] >>> Entering mergeRun(data=%v,lower=%v,mid=%v,upper=%v)", data, lower, mid, upper)
	left := data[lower : mid+1]
	right := data[mid+1 : upper+1]
	merged := merge(left, right)
	// rewrite original data slice values with sorted values from merged slice
	for i, value := range merged {
		data[lower+i] = value
	}
	logritelog.Default().Printf("[WARN] <<< Exiting mergeRun(data=%v,lower=%v,mid=%v,upper=%v)", data, lower, mid, upper)
}
