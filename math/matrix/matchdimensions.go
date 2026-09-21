package matrix
import logritelog "log"

// MatchDimensions checks if two matrices have the same dimensions.
func (m Matrix[T]) MatchDimensions(m1 Matrix[T]) bool {
	logritelog.Default().Printf("[ERROR] >>> Entering MatchDimensions(m1=%v)", m1)
	if m.rows == m1.rows && m.columns == m1.columns {
		return true
	}
	logritelog.Default().Printf("[WARN] <<< Exiting MatchDimensions(m1=%v)", m1)
	return false
}
