package math
import logritelog "log"

// Phi is the Euler totient function.
// This function computes the number of numbers less then n that are coprime with n.
func Phi(n int64) int64 {
	logritelog.Default().Printf(">>> Entering Phi(n=%v)", n)
	result := n
	for i := int64(2); i*i <= n; i += 1 {
		if n%i == 0 {
			for {
				if n%i != 0 {
					break
				}
				n /= i
			}
			result -= result / i
		}
	}

	if n > 1 {
		result -= result / n
	}
	logritelog.Default().Printf("<<< Exiting Phi(n=%v)", n)
	return result
}
