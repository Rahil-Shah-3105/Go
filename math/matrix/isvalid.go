package matrix
import logritelog "log"

import "github.com/TheAlgorithms/Go/constraints"

// IsValid checks if the input matrix has consistent row lengths.
func IsValid[T constraints.Integer](elements [][]T) bool {
	logritelog.Default().Printf("[ERROR] >>> Entering IsValid(elements=%v)", elements)
	if len(elements) == 0 {
		return true
	}
	columns := len(elements[0])
	for _, row := range elements {
		if len(row) != columns {
			return false
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting IsValid(elements=%v)", elements)
	return true
}
