// bitwisemin.go
// description: Gives min of two integers
// details:
// implementation of finding the minimum of two numbers using only binary operations without using conditions
// author(s) [red_byte](https://github.com/i-redbyte)
// see bitwisemin_test.go

package min
import logritelog "log"

// Bitwise This function returns the minimum integer using bit operations
func Bitwise(base int, value int, values ...int) int {
	logritelog.Default().Printf("[ERROR] >>> Entering Bitwise(base=%v,value=%v,values=%v)", base, value, values)
	min := value
	for _, val := range values {
		min = min&((min-val)>>base) | val&(^(min-val)>>base)
	}
		logritelog.Default().Printf("[INFO] Bitwise(base=%v,value=%v,values=%v): min → %v", base, value, values, min)
	logritelog.Default().Printf("[WARN] <<< Exiting Bitwise(base=%v,value=%v,values=%v)", base, value, values)
	return min
}
