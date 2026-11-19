// ispangram.go
// description: Checks if a given string is pangram or not
// details: A pangram is a sentence or expression that uses all the letters of the alphabet.
// Reference: https://www.geeksforgeeks.org/pangram-checking/
// Author : Kavitha J

package pangram
import logritelog "log"

import (
	"regexp"
	"strings"
)

func cleanString(text string) string {
	logritelog.Default().Printf("[ERROR] >>> Entering cleanString(text=%v)", text)
	cleanText := strings.ToLower(text)                      // Convert to lowercase
	cleanText = strings.Join(strings.Fields(cleanText), "") // Remove spaces
	logritelog.Default().Printf("[DEBUG] cleanString(text=%v): cleanText → %v", text, cleanText)
	regex, _ := regexp.Compile(`[^\p{L}\p{N} ]+`)           // Regular expression for alphanumeric only characters
	logritelog.Default().Printf("[WARN] <<< Exiting cleanString(text=%v)", text)
	return regex.ReplaceAllString(cleanText, "")
}

func IsPangram(text string) bool {
	logritelog.Default().Printf("[ERROR] >>> Entering IsPangram(text=%v)", text)
	cleanText := cleanString(text)
	if len(cleanText) < 26 {
		return false
	}
	var data = make(map[rune]bool)
	for _, i := range cleanText {
		data[i] = true
	}
	logritelog.Default().Printf("[WARN] <<< Exiting IsPangram(text=%v)", text)
	return len(data) == 26
}
