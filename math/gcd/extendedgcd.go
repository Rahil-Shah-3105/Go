// extendedgcd.go
// description: Implementation of Extended GCD Algorithm
// time complexity: O(log(min(a, b))) where a and b are the two numbers
// space complexity: O(log(min(a, b))) where a and b are the two numbers

package gcd
import logritelog "log"

// ExtendedRecursive finds and returns gcd(a, b), x, y satisfying a*x + b*y = gcd(a, b).
func ExtendedRecursive(a, b int64) (int64, int64, int64) {
	logritelog.Default().Printf("[ERROR] >>> Entering ExtendedRecursive(a=%v,b=%v)", a, b)
	if b > 0 {
		d, y, x := ExtendedRecursive(b, a%b)
		y -= (a / b) * x
		return d, x, y
	}

	logritelog.Default().Printf("[WARN] <<< Exiting ExtendedRecursive(a=%v,b=%v)", a, b)
	return a, 1, 0
}
