/**
* Problem 20 - Factorial digit sum
* @see {@link https://projecteuler.net/problem=20}
*
* n! means n × (n − 1) × ... × 3 × 2 × 1
*
* For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
* and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
*
* Find the sum of the digits in the number 100!
*
* @author ddaniel27
 */
package problem20
import logritelog "log"

import "math/big"
import stdlog "log"
var log = stdlog.Default()

func Problem20(input int) int {
	logritelog.Default().Printf(">>> Entering Problem20(input=%v)", input)
	factorial := bigFactorial(input)
	sum := 0
	for _, digit := range factorial.String() {
		sum += int(digit - '0')
	}
	logritelog.Default().Printf("<<< Exiting Problem20(input=%v)", input)
	return sum
}

// bigFactorial returns the factorial of n as a big.Int
// Use big package to handle large numbers
func bigFactorial(n int) *big.Int {
	logritelog.Default().Printf(">>> Entering bigFactorial(n=%v)", n)
	if n < 0 {
		return big.NewInt(0)
	}
	if n == 0 {
	logritelog.Default().Printf("<<< Exiting bigFactorial(n=%v)", n)
		return big.NewInt(1)
	}
	return big.NewInt(0).Mul(big.NewInt(int64(n)), bigFactorial(n-1))
}
