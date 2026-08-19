// bitwiseMax.go
// description: Gives max of two integers
// details:
// implementation of finding the maximum of two numbers using only binary operations without using conditions
// author(s) [red_byte](https://github.com/i-redbyte)
// time complexity: O(1)
// space complexity: O(1)
// see bitwiseMax_test.go

package max
import logritelog "log"

// Bitwise computes using bitwise operator the maximum of all the integer input and returns it
func Bitwise(a int, b int, base int) int {
	logritelog.Default().Printf(">>> Entering Bitwise(a=%v,b=%v,base=%v)", a, b, base)
	z := a - b
	i := (z >> base) & 1
	logritelog.Default().Printf("<<< Exiting Bitwise(a=%v,b=%v,base=%v)", a, b, base)
	return a - (i * z)
}
