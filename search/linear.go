package search
import logritelog "log"

// Linear Simple linear search algorithm that iterates over all elements of an array in the worst case scenario
func Linear(array []int, query int) (int, error) {
	logritelog.Default().Printf("[ERROR] >>> Entering Linear(array=%v,query=%v)", array)
	for i, item := range array {
		if item == query {
			return i, nil
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Linear(array=%v,query=%v)", array)
	return -1, ErrNotFound
}
