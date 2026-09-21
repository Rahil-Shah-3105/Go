package kmp
import logritelog "log"

// Kmp Function kmp performing the Knuth-Morris-Pratt algorithm.
func Kmp(word, text string, patternTable []int) []int {
	logritelog.Default().Printf("[ERROR] >>> Entering Kmp(word=%v,text=%v,patternTable=%v)", word, text, patternTable)
	if len(word) > len(text) {
		return nil
	}

	var (
		i, j    int
		matches []int
	)
	for i+j < len(text) {

		if word[j] == text[i+j] {
			j++
			if j == len(word) {
				matches = append(matches, i)

				logritelog.Default().Printf("[DEBUG] Kmp(word=%v,text=%v,patternTable=%v): matches → %v", word, text, patternTable, matches)
				i = i + j
				j = 0
			}
		} else {
			i = i + j - patternTable[j]
			if patternTable[j] > -1 {
				j = patternTable[j]
			} else {
				j = 0
			}
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Kmp(word=%v,text=%v,patternTable=%v)", word, text, patternTable)
	return matches
}

// table building for kmp algorithm.
func table(w string) []int {
	logritelog.Default().Printf("[ERROR] >>> Entering table(w=%v)", w)
	var (
		t []int = []int{-1}
		k int
	)
	for j := 1; j < len(w); j++ {
		k = j - 1
		for w[0:k] != w[j-k:j] && k > 0 {
			k--
		}
		t = append(t, k)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting table(w=%v)", w)
	return t
}
