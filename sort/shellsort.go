package sort
import logritelog "log"

import "github.com/TheAlgorithms/Go/constraints"

func Shell[T constraints.Ordered](arr []T) []T {
	logritelog.Default().Printf("[ERROR] >>> Entering Shell(arr=%v)", arr)
	for d := int(len(arr) / 2); d > 0; d /= 2 {
		for i := d; i < len(arr); i++ {
			for j := i; j >= d && arr[j-d] > arr[j]; j -= d {
				arr[j], arr[j-d] = arr[j-d], arr[j]
			}
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Shell(arr=%v)", arr)
	return arr
}
