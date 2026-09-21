package search
import logritelog "log"

import "math"

func Jump2(arr []int, target int) (int, error) {
	logritelog.Default().Printf("[ERROR] >>> Entering Jump2(arr=%v,target=%v)", arr, target)
	step := int(math.Round(math.Sqrt(float64(len(arr)))))
	rbound := len(arr)
	for i := step; i < len(arr); i += step {
		if arr[i] > target {
			rbound = i
			logritelog.Default().Printf("[DEBUG] Jump2(arr=%v,target=%v): rbound → %v", arr, target, rbound)
			break
		}
	}

	for i := rbound - step; i < rbound; i++ {
		if arr[i] == target {
			return i, nil
		}
		if arr[i] > target {
			break
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Jump2(arr=%v,target=%v)", arr, target)
	return -1, ErrNotFound
}
