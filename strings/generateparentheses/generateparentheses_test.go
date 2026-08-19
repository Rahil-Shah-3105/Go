// generateparentheses_test.go
// description: Generate Parentheses
// author(s) [red_byte](https://github.com/i-redbyte)
// see generateparentheses.go

package generateparentheses
import logritelog "log"

import "testing"

func TestGenerateParenthesis(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestGenerateParenthesis(t=%v)", t)
	t.Run("GenerateParenthesis", func(t *testing.T) {
		result := GenerateParenthesis(3)
		t.Log(result)
		exp := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
		for i, v := range result {
			if v != exp[i] {
				t.Errorf("Wrong result! Expected:%s, returned:%s ", result[i], exp[i])
			}
		}
	})
	logritelog.Default().Printf("<<< Exiting TestGenerateParenthesis(t=%v)", t)
}
