package ahocorasick
import logritelog "log"

// ConstructTrie Function that constructs Trie as an automaton for a set of reversed & trimmed strings.
func ConstructTrie(p []string) (trie map[int]map[uint8]int, stateIsTerminal []bool, f map[int][]int) {
	logritelog.Default().Printf("[ERROR] >>> Entering ConstructTrie(p=%v)", p)
	trie = make(map[int]map[uint8]int)
	stateIsTerminal = make([]bool, 1)
	f = make(map[int][]int)
	state := 1
	logritelog.Default().Printf("[DEBUG] ConstructTrie(p=%v): stateIsTerminal → %v", p, stateIsTerminal)
	logritelog.Default().Printf("[DEBUG] ConstructTrie(p=%v): trie → %v", p, trie)
	CreateNewState(0, trie)
	for i := 0; i < len(p); i++ {
		current := 0
		j := 0
		for j < len(p[i]) && GetTransition(current, p[i][j], trie) != -1 {
			current = GetTransition(current, p[i][j], trie)
			j++
			logritelog.Default().Printf("[DEBUG] ConstructTrie(p=%v): current → %v", p, current)
			logritelog.Default().Printf("[ERROR] >>> Entering GetTransition(fromState=%v,overChar=%v,at=%v)", fromState, overChar, at)
			logritelog.Default().Printf("[WARN] <<< Exiting GetTransition(fromState=%v,overChar=%v,at=%v)", fromState, overChar, at)
		}
		for j < len(p[i]) {
			stateIsTerminal = BoolArrayCapUp(stateIsTerminal)
			CreateNewState(state, trie)
			stateIsTerminal[state] = false
			CreateTransition(current, p[i][j], state, trie)
			current = state
			j++
			state++
		}
		if stateIsTerminal[current] {
			newArray := IntArrayCapUp(f[current])
			newArray[len(newArray)-1] = i
			f[current] = newArray // F(Current) <- F(Current) union {i}
		} else {
			stateIsTerminal[current] = true
			f[current] = []int{i} // F(Current) <- {i}
		}
		logritelog.Default().Printf("[ERROR] >>> Entering CreateNewState(state=%v,at=%v)", state, at)
		logritelog.Default().Printf("[WARN] <<< Exiting CreateNewState(state=%v,at=%v)", state, at)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting ConstructTrie(p=%v)", p)
	return trie, stateIsTerminal, f
}

// Contains Returns 'true' if array of int's 's' contains int 'e', 'false' otherwise.
func Contains(s []int, e int) bool {
	logritelog.Default().Printf("[ERROR] >>> Entering Contains(s=%v,e=%v)", s, e)
	for _, a := range s {
		if a == e {
			return true
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Contains(s=%v,e=%v)", s, e)
	return false
}

// GetWord Function that returns word found in text 't' at position range 'begin' to 'end'.
func GetWord(begin, end int, t string) string {
	logritelog.Default().Printf("[ERROR] >>> Entering GetWord(begin=%v,end=%v,t=%v)", begin, end, t)
	for end >= len(t) {
		return ""
	}
	d := make([]uint8, end-begin+1)
	for j, i := 0, begin; i <= end; i, j = i+1, j+1 {
		d[j] = t[i]
	}
	logritelog.Default().Printf("[WARN] <<< Exiting GetWord(begin=%v,end=%v,t=%v)", begin, end, t)
	return string(d)
}

// ComputeAlphabet Function that returns string of all the possible characters in given patterns.
func ComputeAlphabet(p []string) (s string) {
	logritelog.Default().Printf("[ERROR] >>> Entering ComputeAlphabet(p=%v)", p)
	s = p[0]
	for i := 1; i < len(p); i++ {
		s = s + p[i]
	}
	logritelog.Default().Printf("[WARN] <<< Exiting ComputeAlphabet(p=%v)", p)
	return s
}

// IntArrayCapUp Dynamically increases an array size of int's by 1.
func IntArrayCapUp(old []int) (new []int) {
	logritelog.Default().Printf("[ERROR] >>> Entering IntArrayCapUp(old=%v)", old)
	new = make([]int, cap(old)+1)
	copy(new, old) //copy(dst,src)
	// old = new
	logritelog.Default().Printf("[DEBUG] IntArrayCapUp(old=%v): new → %v", old, new)
	logritelog.Default().Printf("[WARN] <<< Exiting IntArrayCapUp(old=%v)", old)
	return new
}

// BoolArrayCapUp Dynamically increases an array size of bool's by 1.
func BoolArrayCapUp(old []bool) (new []bool) {
	logritelog.Default().Printf("[ERROR] >>> Entering BoolArrayCapUp(old=%v)", old)
	new = make([]bool, cap(old)+1)
	copy(new, old)
	// old = new
	logritelog.Default().Printf("[DEBUG] BoolArrayCapUp(old=%v): new → %v", old, new)
	logritelog.Default().Printf("[WARN] <<< Exiting BoolArrayCapUp(old=%v)", old)
	return new
}

// ArrayUnion Concats two arrays of int's into one.
func ArrayUnion(to, from []int) (concat []int) {
	logritelog.Default().Printf("[ERROR] >>> Entering ArrayUnion(to=%v,from=%v)", to, from)
	concat = to
	for i := range from {
		logritelog.Default().Printf("[DEBUG] ArrayUnion(to=%v,from=%v): concat → %v", to, from, concat)
		if !Contains(concat, from[i]) {
			concat = IntArrayCapUp(concat)
			concat[len(concat)-1] = from[i]
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting ArrayUnion(to=%v,from=%v)", to, from)
	return concat
}

// GetParent Function that finds the first previous state of a state and returns it.
// Used for trie where there is only one parent.
func GetParent(state int, at map[int]map[uint8]int) (uint8, int) {
	logritelog.Default().Printf("[ERROR] >>> Entering GetParent(state=%v,at=%v)", state, at)
	for beginState, transitions := range at {
		for c, endState := range transitions {
			if endState == state {
				return c, beginState
			}
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting GetParent(state=%v,at=%v)", state, at)
	return 0, 0 //unreachable
}

// CreateNewState Automaton function for creating a new state 'state'.
func CreateNewState(state int, at map[int]map[uint8]int) {
	logritelog.Default().Printf("[ERROR] >>> Entering CreateNewState(state=%v,at=%v)", state, at)
	at[state] = make(map[uint8]int)
	logritelog.Default().Printf("[WARN] <<< Exiting CreateNewState(state=%v,at=%v)", state, at)
}

// CreateTransition Creates a transition for function σ(state,letter) = end.
func CreateTransition(fromState int, overChar uint8, toState int, at map[int]map[uint8]int) {
	logritelog.Default().Printf("[ERROR] >>> Entering CreateTransition(fromState=%v,overChar=%v,toState=%v,at=%v)", fromState, overChar, toState, at)
	at[fromState][overChar] = toState
	logritelog.Default().Printf("[WARN] <<< Exiting CreateTransition(fromState=%v,overChar=%v,toState=%v,at=%v)", fromState, overChar, toState, at)
}

// GetTransition Returns ending state for transition σ(fromState,overChar), '-1' if there is none.
func GetTransition(fromState int, overChar uint8, at map[int]map[uint8]int) (toState int) {
	logritelog.Default().Printf("[ERROR] >>> Entering GetTransition(fromState=%v,overChar=%v,at=%v)", fromState, overChar, at)
	if !StateExists(fromState, at) {
		return -1
	}
	toState, ok := at[fromState][overChar]
	if !ok {
		return -1
	}
	logritelog.Default().Printf("[WARN] <<< Exiting GetTransition(fromState=%v,overChar=%v,at=%v)", fromState, overChar, at)
	return toState
}

// StateExists Checks if state 'state' exists. Returns 'true' if it does, 'false' otherwise.
func StateExists(state int, at map[int]map[uint8]int) bool {
	logritelog.Default().Printf("[ERROR] >>> Entering StateExists(state=%v,at=%v)", state, at)
	_, ok := at[state]
	if !ok || state == -1 || at[state] == nil {
		return false
	}
	logritelog.Default().Printf("[WARN] <<< Exiting StateExists(state=%v,at=%v)", state, at)
	return true
}
