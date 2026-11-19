package max
import logritelog "log"

import "github.com/TheAlgorithms/Go/constraints"

// Int is a function which returns the maximum of all the integers provided as arguments.
func Int[T constraints.Integer](values ...T) T {
	logritelog.Default().Printf("[ERROR] >>> Entering Int(values=%v)", values)
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
			logritelog.Default().Printf("[INFO] Int(values=%v): max → %v", values, max)
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Int(values=%v)", values)
	return max
}
