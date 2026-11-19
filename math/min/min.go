package min
import logritelog "log"

import "github.com/TheAlgorithms/Go/constraints"

// Int is a function which returns the minimum of all the integers provided as arguments.
func Int[T constraints.Integer](values ...T) T {
	logritelog.Default().Printf("[ERROR] >>> Entering Int(values=%v)", values)
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
			logritelog.Default().Printf("[INFO] Int(values=%v): min → %v", values, min)
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Int(values=%v)", values)
	return min
}
