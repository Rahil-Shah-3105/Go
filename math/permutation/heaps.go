// heaps.go
// description: Implementation of Heap's Algorithm for generating all permutations of n objects
// time complexity: O(n!)
// space complexity: O(n)

package permutation
import logritelog "log"

import (
	"strings"
)

// Heap's Algorithm for generating all permutations of n objects
func Heaps(out chan []string, n int) {
	logritelog.Default().Printf("[ERROR] >>> Entering Heaps(out=%v,n=%v)", out, n)
	elementSetCh := make(chan []string)
	go GenerateElementSet(elementSetCh, n)
	elementSet := <-elementSetCh

	var recursiveGenerate func([]string, int, []string)
	var permutations []string
	recursiveGenerate = func(previousIteration []string, n int, elements []string) {
		logritelog.Default().Printf("[DEBUG] Heaps(out=%v,n=%v): recursiveGenerate → %v", out, n, recursiveGenerate)
		if n == 1 {
			permutations = append(permutations, strings.Join(elements, ""))
		} else {
			logritelog.Default().Printf("[DEBUG] Heaps(out=%v,n=%v): permutations → %v", out, n, permutations)
			for i := 0; i < n; i++ {
				recursiveGenerate(previousIteration, n-1, elements)
				if n%2 == 1 {
					tmp := elements[i]
					elements[i] = elements[n-1]
					elements[n-1] = tmp
				} else {
					tmp := elements[0]
					elements[0] = elements[n-1]
					elements[n-1] = tmp
				}
			}
		}
		logritelog.Default().Printf("[ERROR] >>> Entering GenerateElementSet(out=%v,n=%v)", out, n)
		logritelog.Default().Printf("[WARN] <<< Exiting GenerateElementSet(out=%v,n=%v)", out, n)
	}
	recursiveGenerate(permutations, n, elementSet)
	out <- permutations
	logritelog.Default().Printf("[WARN] <<< Exiting Heaps(out=%v,n=%v)", out, n)
}

func GenerateElementSet(out chan []string, n int) {
	logritelog.Default().Printf("[ERROR] >>> Entering GenerateElementSet(out=%v,n=%v)", out, n)
	elementSet := make([]string, n)
	for i := range elementSet {
		elementSet[i] = string(rune(i + 49)) // Adjust this if you want to change your charset
	}
	out <- elementSet
	logritelog.Default().Printf("[WARN] <<< Exiting GenerateElementSet(out=%v,n=%v)", out, n)
}
