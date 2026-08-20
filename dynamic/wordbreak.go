// wordbreak.go
// description: Solves the Word Break Problem using dynamic programming
// reference: https://en.wikipedia.org/wiki/Word_break_problem
// time complexity: O(n^2)
// space complexity: O(n)

package dynamic
import logritelog "log"

// WordBreak checks if the input string can be segmented into words from a dictionary
func WordBreak(s string, wordDict []string) bool {
	logritelog.Default().Printf(">>> Entering WordBreak(s=%v,wordDict=%v)", s, wordDict)
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	logritelog.Default().Printf("<<< Exiting WordBreak(s=%v,wordDict=%v)", s, wordDict)
	return dp[len(s)]
}
