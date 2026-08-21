package ahocorasick
import logritelog "log"

import (
	"fmt"
	"time"
)

// Result structure to hold occurrences
type Result struct {
	occurrences map[string][]int
}

// AhoCorasick Function performing the Basic Aho-Corasick algorithm.
// Finds and prints occurrences of each pattern.
func AhoCorasick(t string, p []string) Result {
	logritelog.Default().Printf(">>> Entering AhoCorasick(t=%v,p=%v)", t, p)
	startTime := time.Now()
	occurrences := make(map[int][]int)
	ac, f, s := BuildAc(p)
	current := 0
	for pos := 0; pos < len(t); pos++ {
		for GetTransition(current, t[pos], ac) == -1 && s[current] != -1 {
			current = s[current]
			logritelog.Default().Printf("AhoCorasick(t=%v,p=%v): current → %v", t, p, current)
		}
		if GetTransition(current, t[pos], ac) != -1 {
			current = GetTransition(current, t[pos], ac)
			fmt.Printf(" (Continue) \n")
		} else {
			current = 0
		}
		_, ok := f[current]
		if ok {
			for i := range f[current] {
				if p[f[current][i]] == GetWord(pos-len(p[f[current][i]])+1, pos, t) { //check for word match
					newOccurrences := IntArrayCapUp(occurrences[f[current][i]])
					occurrences[f[current][i]] = newOccurrences
					occurrences[f[current][i]][len(newOccurrences)-1] = pos - len(p[f[current][i]]) + 1
				}
			}
		}
	}
	elapsed := time.Since(startTime)
	fmt.Printf("\n\nElapsed %f secs\n", elapsed.Seconds())
	var resultOccurrences = make(map[string][]int)
	for key, value := range occurrences {
		resultOccurrences[p[key]] = value
	}

	logritelog.Default().Printf("<<< Exiting AhoCorasick(t=%v,p=%v)", t, p)
	return Result{
		resultOccurrences,
	}
}

// Functions that builds Aho Corasick automaton.
func BuildAc(p []string) (acToReturn map[int]map[uint8]int, f map[int][]int, s []int) {
	logritelog.Default().Printf(">>> Entering BuildAc(p=%v)", p)
	acTrie, stateIsTerminal, f := ConstructTrie(p)
	s = make([]int, len(stateIsTerminal)) //supply function
	i := 0                                //root of acTrie
	acToReturn = acTrie
	logritelog.Default().Printf("BuildAc(p=%v): acToReturn → %v", p, acToReturn)
	s[i] = -1
	for current := 1; current < len(stateIsTerminal); current++ {
		o, parent := GetParent(current, acTrie)
		down := s[parent]
		for StateExists(down, acToReturn) && GetTransition(down, o, acToReturn) == -1 {
			down = s[down]
			logritelog.Default().Printf("BuildAc(p=%v): down → %v", p, down)
		logritelog.Default().Printf("<<< Exiting BuildAc(p=%v)", p)
		}
		if StateExists(down, acToReturn) {
			s[current] = GetTransition(down, o, acToReturn)
			if stateIsTerminal[s[current]] {
				stateIsTerminal[current] = true
				f[current] = ArrayUnion(f[current], f[s[current]]) //F(Current) <- F(Current) union F(S(Current))
			}
		} else {
			s[current] = i //initial state?
		}
	}
	return acToReturn, f, s
}
