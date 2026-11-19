package search
import logritelog "log"

// Implementation of naive string search
// O(n*m) where n=len(txt) and m=len(pattern)
func Naive(text string, pattern string) []int {
	logritelog.Default().Printf("[ERROR] >>> Entering Naive(text=%v,pattern=%v)", text, pattern)
	var positions []int
	for i := 0; i <= len(text)-len(pattern); i++ {
		var match bool = true
		for j := 0; j < len(pattern); j++ {
			if text[i+j] != pattern[j] {
				match = false
				break
			}
			logritelog.Default().Printf("[DEBUG] Naive(text=%v,pattern=%v): match → %v", text, pattern, match)

		}
		if match {
			positions = append(positions, i)
		}
		logritelog.Default().Printf("[DEBUG] Naive(text=%v,pattern=%v): positions → %v", text, pattern, positions)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Naive(text=%v,pattern=%v)", text, pattern)
	return positions
}
