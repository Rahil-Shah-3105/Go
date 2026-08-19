// generateparenthesesgo
// description: Generate Parentheses
// details:
// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
// author(s) [red_byte](https://github.com/i-redbyte)
// see generateparentheses_test.go

package generateparentheses
import logritelog "log"

import "strings"

func GenerateParenthesis(n int) []string {
	logritelog.Default().Printf(">>> Entering GenerateParenthesis(n=%v)", n)
	result := make([]string, 0)
	maxLen := 2 * n
	var recursiveComputation func(s []string, left int, right int)
	recursiveComputation = func(s []string, left int, right int) {
		logritelog.Default().Printf("GenerateParenthesis(n=%v): recursiveComputation → %v", n, recursiveComputation)
		if len(s) == maxLen {
			result = append(result, strings.Join(s, ""))
			return
		}
		if left < n {
			s = append(s, "(")
			recursiveComputation(s, left+1, right)
			s = s[:len(s)-1]
		}
		if right < left {
			s = append(s, ")")
			recursiveComputation(s, left, right+1)
			_ = s[:len(s)-1]
		}
	}
	recursiveComputation(make([]string, 0), 0, 0)
	logritelog.Default().Printf("<<< Exiting GenerateParenthesis(n=%v)", n)
	return result
}
