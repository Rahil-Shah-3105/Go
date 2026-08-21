package matrix
import logritelog "log"

import "fmt"

// String implements the fmt.Stringer interface for Matrix.
func (m Matrix[T]) String() string {
	logritelog.Default().Printf(">>> Entering String()")
	var result string
	for i := range m.elements {
		for j := range m.elements[i] {
			result += fmt.Sprintf("%v ", m.elements[i][j])
		}
		result += "\n"
	}
	logritelog.Default().Printf("<<< Exiting String()")
	return result
}
