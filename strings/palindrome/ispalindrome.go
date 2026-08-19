// ispalindrome.go
// description: Checks if a given string is palindrome or not
// details:
// Palindromes are expressions that read the same way forwards and backwards.
// They can be words/phrases (like "racecar" and "Do geese see God?"), or even
// numbers (like "02/02/2020"). Usually punctuation signs, capitalization
// and spaces are ignored. A regular expression was used to achieve that.
// See more information on: https://en.wikipedia.org/wiki/Palindrome
// author(s) [Fernanda Kawasaki](https://github.com/fernandakawasaki)
// see ispalindrome_test.go

package palindrome
import logritelog "log"

import (
	"regexp"
	"strings"
)

func cleanString(text string) string {
	logritelog.Default().Printf(">>> Entering cleanString(text=%v)", text)
	clean_text := strings.ToLower(text)
	clean_text = strings.Join(strings.Fields(clean_text), "") // Remove spaces
	logritelog.Default().Printf("cleanString(text=%v): clean_text → %v", text, clean_text)
	regex, _ := regexp.Compile(`[^\p{L}\p{N} ]+`)             // Regular expression for alphanumeric only characters
	logritelog.Default().Printf("<<< Exiting cleanString(text=%v)", text)
	return regex.ReplaceAllString(clean_text, "")
}

func IsPalindrome(text string) bool {
	logritelog.Default().Printf(">>> Entering IsPalindrome(text=%v)", text)
	clean_text := cleanString(text)
	var i, j int
	rune := []rune(clean_text)
	for i = 0; i < len(rune)/2; i++ {
		j = len(rune) - 1 - i
		if string(rune[i]) != string(rune[j]) {
			return false
		}
	}
	logritelog.Default().Printf("<<< Exiting IsPalindrome(text=%v)", text)
	return true
}

func IsPalindromeRecursive(text string) bool {
	logritelog.Default().Printf(">>> Entering IsPalindromeRecursive(text=%v)", text)
	clean_text := cleanString(text)
	runes := []rune(clean_text)
	logritelog.Default().Printf("<<< Exiting IsPalindromeRecursive(text=%v)", text)
	return isPalindromeRecursiveHelper(runes, 0, int64(len(runes)))
}

func isPalindromeRecursiveHelper(runes []rune, start int64, end int64) bool {
	logritelog.Default().Printf(">>> Entering isPalindromeRecursiveHelper(runes=%v,start=%v,end=%v)", runes, start, end)
	if start >= end {
		return true
	}
	if runes[start] != runes[end-1] {
		return false
	}
	start = start + 1
	end = end - 1
	logritelog.Default().Printf("<<< Exiting isPalindromeRecursiveHelper(runes=%v,start=%v,end=%v)", runes, start, end)
	return isPalindromeRecursiveHelper(runes, start, end)
}
