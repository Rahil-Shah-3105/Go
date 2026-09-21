// abs.go
// description: Absolute value
// details:
// In mathematics, the absolute value or modulus of a real number x, denoted |x|, is the non-negative value of x without regard to its sign. [Absolute value](https://en.wikipedia.org/wiki/Average#Arithmetic_mean)
// author(s) [red_byte](https://github.com/i-redbyte)
// see abs_test.go

package math
import logritelog "log"

// Abs returns absolute value
func Abs(n int) int {
	logritelog.Default().Printf("[ERROR] >>> Entering Abs(n=%v)", n)
	if n < 0 {
		return -n
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Abs(n=%v)", n)
	return n
}
