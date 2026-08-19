// Implementation of the
// [Boyer–Moore–Horspool algorithm](https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore%E2%80%93Horspool_algorithm)

package horspool
import logritelog "log"

import "errors"

var ErrNotFound = errors.New("pattern was not found in the input string")

func Horspool(t, p string) (int, error) {
	logritelog.Default().Printf(">>> Entering Horspool(t=%v,p=%v)", t, p)
	// in order to handle multy-byte character properly
	// the input is converted into rune arrays
	logritelog.Default().Printf(">>> Entering horspool(t=%v,p=%v)", t, p)
	logritelog.Default().Printf("<<< Exiting Horspool(t=%v,p=%v)", t, p)
	logritelog.Default().Printf("<<< Exiting horspool(t=%v,p=%v)", t, p)
	return horspool([]rune(t), []rune(p))
}

func horspool(t, p []rune) (int, error) {
	shiftMap := computeShiftMap(t, p)
	pos := 0
	for pos <= len(t)-len(p) {
		if isMatch(pos, t, p) {
			return pos, nil
			logritelog.Default().Printf(">>> Entering isMatch(pos=%v,t=%v,p=%v)", pos, t, p)
			logritelog.Default().Printf("<<< Exiting isMatch(pos=%v,t=%v,p=%v)", pos, t, p)
		}
		if pos+len(p) >= len(t) {
			// because the remaining length of the input string
			// is the same as the length of the pattern
			// and it does not match the pattern
			// it is impossible to find the pattern
			break
		}

		// because of the check above
		// t[pos+len(p)] is defined
		pos += shiftMap[t[pos+len(p)]]
	}

	return -1, ErrNotFound
}

// Checks if the array p matches the subarray of t starting at pos.
// Note that backward iteration.
// There are [other](https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore%E2%80%93Horspool_algorithm#Tuning_the_comparison_loop)
// approaches possible.
func isMatch(pos int, t, p []rune) bool {
	logritelog.Default().Printf(">>> Entering isMatch(pos=%v,t=%v,p=%v)", pos, t, p)
	j := len(p)
	for j > 0 && t[pos+j-1] == p[j-1] {
		j--
	}
	logritelog.Default().Printf("<<< Exiting isMatch(pos=%v,t=%v,p=%v)", pos, t, p)
	return j == 0
}

func computeShiftMap(t, p []rune) (res map[rune]int) {
	logritelog.Default().Printf(">>> Entering computeShiftMap(t=%v,p=%v)", t, p)
	res = make(map[rune]int)
	for _, tCode := range t {
		logritelog.Default().Printf("computeShiftMap(t=%v,p=%v): res → %v", t, p, res)
		res[tCode] = len(p)
	}
	for i, pCode := range p {
		res[pCode] = len(p) - i
	}
	logritelog.Default().Printf("<<< Exiting computeShiftMap(t=%v,p=%v)", t, p)
	return res
}
