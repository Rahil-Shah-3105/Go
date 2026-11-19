package math
import logritelog "log"

import (
	"github.com/TheAlgorithms/Go/constraints"
)

func Mean[T constraints.Number](values []T) float64 {
logritelog.Default().Printf("[ERROR] >>> Entering Mean(values=%v)", values)

	if len(values) == 0 {
		return 0
	}

	var summation float64 = 0

	for _, singleValue := range values {
		summation += float64(singleValue)
	}

	logritelog.Default().Printf("[WARN] <<< Exiting Mean(values=%v)", values)
	return summation / float64(len(values))

}
