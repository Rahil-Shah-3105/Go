// time complexity: O(log(min(a, b))) where a and b are the two numbers
// space complexity: O(1)

package gcd
import logritelog "log"

// Recursive finds and returns the greatest common divisor of a given integer.
func Recursive(a, b int64) int64 {
	logritelog.Default().Printf(">>> Entering Recursive(a=%v,b=%v)", a, b)
	if b == 0 {
	logritelog.Default().Printf("<<< Exiting Recursive(a=%v,b=%v)", a, b)
		return a
	}
	return Recursive(b, a%b)
}
