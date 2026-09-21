// time complexity: O(log(min(a, b))) where a and b are the two numbers
// space complexity: O(1)

package gcd
import logritelog "log"

// Iterative Faster iterative version of GcdRecursive without holding up too much of the stack
func Iterative(a, b int64) int64 {
	logritelog.Default().Printf("[ERROR] >>> Entering Iterative(a=%v,b=%v)", a, b)
	for b != 0 {
		a, b = b, a%b
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Iterative(a=%v,b=%v)", a, b)
	return a
}
