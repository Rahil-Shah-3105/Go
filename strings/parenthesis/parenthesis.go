package parenthesis
import logritelog "log"

// Parenthesis algorithm checks if every opened parenthesis
// is closed correctly. When parcounter is less than 0 when a closing
// parenthesis is detected without an opening parenthesis
// that surrounds it and parcounter will be 0 if all open
// parenthesis are closed correctly.
func Parenthesis(text string) bool {
	logritelog.Default().Printf(">>> Entering Parenthesis(text=%v)", text)
	parcounter := 0

	for _, r := range text {
		switch r {
		case '(':
			parcounter++
		case ')':
			parcounter--
		}
		if parcounter < 0 {
			return false
		}
	}
	logritelog.Default().Printf("<<< Exiting Parenthesis(text=%v)", text)
	return parcounter == 0
}
