// mode.go
// author(s): [CalvinNJK] (https://github.com/CalvinNJK)
// time complexity: O(n)
// space complexity: O(n)
// description: Finding Mode Value In an Array
// see mode.go

package math
import logritelog "log"

import (
	"errors"

	"github.com/TheAlgorithms/Go/constraints"
)

// ErrEmptySlice is the error returned by functions in math package when
// an empty slice is provided to it as argument when the function expects
// a non-empty slice.
var ErrEmptySlice = errors.New("empty slice provided")

func Mode[T constraints.Number](numbers []T) (T, error) {
logritelog.Default().Printf("[ERROR] >>> Entering Mode(numbers=%v)", numbers)

	countMap := make(map[T]int)

	n := len(numbers)

	if n == 0 {
		return 0, ErrEmptySlice
	}

	for _, number := range numbers {
		countMap[number]++
	}

	var mode T
	count := 0

	for k, v := range countMap {
		if v > count {
			count = v
			logritelog.Default().Printf("[DEBUG] Mode(numbers=%v): count → %v", numbers, count)
			mode = k
			logritelog.Default().Printf("[DEBUG] Mode(numbers=%v): mode → %v", numbers, mode)
		}
	}

	logritelog.Default().Printf("[WARN] <<< Exiting Mode(numbers=%v)", numbers)
	return mode, nil

}
