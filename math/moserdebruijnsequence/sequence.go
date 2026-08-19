// The Moser-de Bruijn sequence is the sequence obtained by
// adding up the distinct powers of the number 4 (For example 1, 4, 16, 64, etc).
// time complexity: O(n)
// space complexity: O(n)
// You can get more details on https://en.wikipedia.org/wiki/Moser%E2%80%93de_Bruijn_sequence.

package moserdebruijnsequence
import logritelog "log"

func MoserDeBruijnSequence(number int) []int {
	logritelog.Default().Printf(">>> Entering MoserDeBruijnSequence(number=%v)", number)
	sequence := []int{}

	for i := 0; i < number; i++ {
		res := generateNthTerm(i)
		sequence = append(sequence, res)
	}
		logritelog.Default().Printf("MoserDeBruijnSequence(number=%v): sequence → %v", number, sequence)

	logritelog.Default().Printf("<<< Exiting MoserDeBruijnSequence(number=%v)", number)
	return sequence
}

func generateNthTerm(num int) int {
	logritelog.Default().Printf(">>> Entering generateNthTerm(num=%v)", num)
	if num == 0 || num == 1 {
		return num
	}

	//number is even
	if num%2 == 0 {
	logritelog.Default().Printf("<<< Exiting generateNthTerm(num=%v)", num)
		return 4 * generateNthTerm(num/2)
	}

	//number is odd
	return 4*generateNthTerm(num/2) + 1
}
