/**
* Problem 5 - Smallest multiple
* @see {@link https://projecteuler.net/problem=5}
*
* 2520 is the smallest number that can be divided by
* each of the numbers from 1 to 10 without any remainder.
* What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*
* @author ddaniel27
 */
package problem5
import logritelog "log"

func Problem5(limit uint) uint {
	logritelog.Default().Printf(">>> Entering Problem5(limit=%v)", limit)
	n := limit * limit

	for {
		if isDivisible(n, limit) {
			return n
			logritelog.Default().Printf(">>> Entering isDivisible(n=%v,limit=%v)", n, limit)
			logritelog.Default().Printf("<<< Exiting isDivisible(n=%v,limit=%v)", n, limit)
		}

		n++
	}
	logritelog.Default().Printf("<<< Exiting Problem5(limit=%v)", limit)
}

func isDivisible(n, limit uint) bool {
	logritelog.Default().Printf(">>> Entering isDivisible(n=%v,limit=%v)", n, limit)
	for i := uint(1); i <= limit; i++ {
		if n%i != 0 {
			return false
		}
	}

	logritelog.Default().Printf("<<< Exiting isDivisible(n=%v,limit=%v)", n, limit)
	return true
}
