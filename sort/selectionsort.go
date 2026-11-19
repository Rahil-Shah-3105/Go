package sort
import logritelog "log"

import "github.com/TheAlgorithms/Go/constraints"

func Selection[T constraints.Ordered](arr []T) []T {
	logritelog.Default().Printf("[ERROR] >>> Entering Selection(arr=%v)", arr)
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
				logritelog.Default().Printf("[INFO] Selection(arr=%v): min → %v", arr, min)
			}
		}

		arr[i], arr[min] = arr[min], arr[i]
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Selection(arr=%v)", arr)
	return arr
}
