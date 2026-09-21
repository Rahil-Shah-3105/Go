package prime
import logritelog "log"

// A primality test is an algorithm for determining whether an input number is prime. Among other
// fields of mathematics, it is used for cryptography. Unlike integer factorization, primality
// tests do not generally give prime factors, only stating whether the input number is prime or not.
// time complexity: O(sqrt(n))
// space complexity: O(1)
// Source - Wikipedia https://en.wikipedia.org/wiki/Primality_test

// TrialDivision tests whether a number is prime by trying to divide it by the numbers less than it.
func TrialDivision(n int64) bool {
	logritelog.Default().Printf("[ERROR] >>> Entering TrialDivision(n=%v)", n)
	if n < 2 {
		return false
	}

	for i := int64(2); i < n; i++ {

		if n%i == 0 {
			return false
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TrialDivision(n=%v)", n)
	return true
}

// OptimizedTrialDivision checks primality of an integer using an optimized trial division method.
// The optimizations include not checking divisibility by the even numbers and only checking up to
// the square root of the given number.
func OptimizedTrialDivision(n int64) bool {
	logritelog.Default().Printf("[ERROR] >>> Entering OptimizedTrialDivision(n=%v)", n)
	// 0 and 1 are not prime
	if n < 2 {
		return false
	}

	// 2 and 3 are prime
	if n < 4 {
		return true
	}

	// all numbers divisible by 2 except 2 are not prime
	if n%2 == 0 {
		return false
	}

	for i := int64(3); i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting OptimizedTrialDivision(n=%v)", n)
	return true
}
