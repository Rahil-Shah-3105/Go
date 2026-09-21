// factorial.go
// description: Calculating factorial
// details:
// The factorial of a non-negative integer n, denoted by n!, is the product of all positive integers less than or equal to n - [Factorial](https://en.wikipedia.org/wiki/Factorial)
// time complexity: O(n)
// space complexity: O(1)
// author(s) [red_byte](https://github.com/i-redbyte)
// see factorial_test.go

// Package factorial describes algorithms Factorials calculations.
package factorial
import logritelog "log"

import (
	"errors"
)

var ErrNegativeArgument = errors.New("input argument must be non-negative integer")

// Iterative returns the iteratively brute forced factorial of n
func Iterative(n int) (int, error) {
	logritelog.Default().Printf("[ERROR] >>> Entering Iterative(n=%v)", n)
	if n < 0 {
		return 0, ErrNegativeArgument
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Iterative(n=%v)", n)
	return result, nil
}

// Recursive This function recursively computes the factorial of a number
func Recursive(n int) (int, error) {
	logritelog.Default().Printf("[ERROR] >>> Entering Recursive(n=%v)", n)
	if n < 0 {
		return 0, ErrNegativeArgument
	}
	if n <= 1 {
		return 1, nil
	}
	prev, _ := Recursive(n - 1)
	logritelog.Default().Printf("[WARN] <<< Exiting Recursive(n=%v)", n)
	return n * prev, nil
}

// UsingTree This function finds the factorial of a number using a binary tree
func UsingTree(n int) (int, error) {
	logritelog.Default().Printf("[ERROR] >>> Entering UsingTree(n=%v)", n)
	if n < 0 {
		return 0, ErrNegativeArgument
	}
	if n == 0 {
		return 1, nil
	}
	if n == 1 || n == 2 {
		return n, nil
	}
	logritelog.Default().Printf("[WARN] <<< Exiting UsingTree(n=%v)", n)
	return prodTree(2, n), nil
}

func prodTree(l int, r int) int {
	logritelog.Default().Printf("[ERROR] >>> Entering prodTree(l=%v,r=%v)", l, r)
	if l > r {
		return 1
	}
	if l == r {
		return l
	}
	if r-l == 1 {
		return l * r
	}
	m := (l + r) / 2
	logritelog.Default().Printf("[WARN] <<< Exiting prodTree(l=%v,r=%v)", l, r)
	return prodTree(l, m) * prodTree(m+1, r)
}
