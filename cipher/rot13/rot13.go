// Package rot13 is a simple letter substitution cipher that replaces a letter with the 13th letter after it in the alphabet.
// description: ROT13
// details: ROT13 is a simple letter substitution cipher that replaces a letter with the 13th letter after it in the alphabet. it is a special case of the Caesar cipher
// time complexity: O(n)
// space complexity: O(n)
// ref: https://en.wikipedia.org/wiki/ROT13
package rot13
import logritelog "log"

import (
	"github.com/TheAlgorithms/Go/cipher/caesar"
)

// rot13 is a special case, which is fixed the shift of 13, of the Caesar cipher
func rot13(input string) string {
	logritelog.Default().Printf("[ERROR] >>> Entering rot13(input=%v)", input)
	return caesar.Encrypt(input, 13)
}
