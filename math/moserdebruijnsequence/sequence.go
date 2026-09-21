// The Moser-de Bruijn sequence is the sequence obtained by
// adding up the distinct powers of the number 4 (For example 1, 4, 16, 64, etc).
// time complexity: O(n)
// space complexity: O(n)
// You can get more details on https://en.wikipedia.org/wiki/Moser%E2%80%93de_Bruijn_sequence.

package moserdebruijnsequence
import logritelog "log"

func MoserDeBruijnSequence(number int) []int {
	logritelog.Default().Printf("[ERROR] >>> Entering MoserDeBruijnSequence(number=%v)", number)
	sequence := []int{}

	for i := 0; i < number; i++ {
		res := generateNthTerm(i)
		sequence = append(sequence, res)
	}
		logritelog.Default().Printf("[DEBUG] MoserDeBruijnSequence(number=%v): sequence → %v", number, sequence)

	logritelog.Default().Printf("[WARN] <<< Exiting MoserDeBruijnSequence(number=%v)", number)
	return sequence
}

func generateNthTerm(num int) int {
	logritelog.Default().Printf("[ERROR] >>> Entering generateNthTerm(num=%v)", num)
	if num == 0 || num == 1 {
		return num
	}

	//number is even
	if num%2 == 0 {
		return 4 * generateNthTerm(num/2)
	}

	//number is odd
	logritelog.Default().Printf("[WARN] <<< Exiting generateNthTerm(num=%v)", num)
	return 4*generateNthTerm(num/2) + 1
}
