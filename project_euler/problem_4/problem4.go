/**
* Problem 4 - Largest palindrome product
* @see {@link https://projecteuler.net/problem=4}
*
* A palindromic number reads the same both ways.
* The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
* Find the largest palindrome made from the product of two 3-digit numbers.
*
* @author ddaniel27
 */
package problem4
import logritelog "log"

import (
	"fmt"

	"github.com/TheAlgorithms/Go/strings/palindrome"
)

func Problem4() uint {
	logritelog.Default().Printf("[ERROR] >>> Entering Problem4()")
	max := uint(0)

	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			n := uint(i * j)

			if palindrome.IsPalindrome(fmt.Sprintf("%d", n)) && n > max {
				max = n
				logritelog.Default().Printf("[INFO] Problem4(): max → %v", max)
			}
		}
	}

	logritelog.Default().Printf("[WARN] <<< Exiting Problem4()")
	return max
}
