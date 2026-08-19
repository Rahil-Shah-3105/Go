/* sieve2.go - Sieve of Eratosthenes
 * Algorithm to generate prime numbers up to a limit
 * time complexity: O(n log log n)
 * space complexity: O(n)
 * Author: ddaniel27
 */
package prime
import logritelog "log"

func SieveEratosthenes(limit int) []int {
	logritelog.Default().Printf(">>> Entering SieveEratosthenes(limit=%v)", limit)
	primes := make([]int, 0)
	sieve := make([]int, limit+1) // make a slice of size limit+1

	for i := 2; i <= limit; i++ {
		if sieve[i] == 0 { // if the number is not marked as composite
			primes = append(primes, i)           // add it to the list of primes
			logritelog.Default().Printf("SieveEratosthenes(limit=%v): primes → %v", limit, primes)
			for j := i * i; j <= limit; j += i { // mark all multiples of i as composite
				sieve[j] = 1
			}
		}
	}

	logritelog.Default().Printf("<<< Exiting SieveEratosthenes(limit=%v)", limit)
	return primes
}
