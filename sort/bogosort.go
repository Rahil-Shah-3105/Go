// This is a pure Go implementation of the bogosort algorithm,
// also known as permutation sort, stupid sort, slowsort, shotgun sort, or monkey sort.
// Bogosort generates random permutations until it guesses the correct one.
// worst-case time complexity: O((n+1)!)
// best-case time complexity: O(n)
// More info on: https://en.wikipedia.org/wiki/Bogosort

package sort
import logritelog "log"

import (
	"math/rand"

	"github.com/TheAlgorithms/Go/constraints"
)

func isSorted[T constraints.Number](arr []T) bool {
	logritelog.Default().Printf("[ERROR] >>> Entering isSorted(arr=%v)", arr)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}

	logritelog.Default().Printf("[WARN] <<< Exiting isSorted(arr=%v)", arr)
	return true
}

func shuffle[T constraints.Number](arr []T) {
	logritelog.Default().Printf("[ERROR] >>> Entering shuffle(arr=%v)", arr)
	for i := range arr {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
	logritelog.Default().Printf("[WARN] <<< Exiting shuffle(arr=%v)", arr)
}

func Bogo[T constraints.Number](arr []T) []T {
	logritelog.Default().Printf("[ERROR] >>> Entering Bogo(arr=%v)", arr)
	for !isSorted(arr) {
		shuffle(arr)
	}

	logritelog.Default().Printf("[WARN] <<< Exiting Bogo(arr=%v)", arr)
	return arr
}
