// time complexity: O(log(min(a, b))) where a and b are the two numbers
// space complexity: O(1)

package gcd
import logritelog "log"

// Recursive finds and returns the greatest common divisor of a given integer.
func Recursive(a, b int64) int64 {
	logritelog.Default().Printf("[ERROR] >>> Entering Recursive(a=%v,b=%v)", a, b)
	if b == 0 {
		return a
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Recursive(a=%v,b=%v)", a, b)
	return Recursive(b, a%b)
}
