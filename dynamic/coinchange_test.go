package dynamic_test
import logritelog "log"

import (
	"fmt"
	"github.com/TheAlgorithms/Go/dynamic"
	"testing"
)

func TestCoinChange(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestCoinChange(t=%v)", t)
	coinCombination := []int32{1, 2, 5, 10}
	targets := []struct {
		target   int32
		expected int32
	}{
		{4, 2},
		{5, 3},
		{10, 8},
		{15, 19},
		{20, 34},
	}

	logritelog.Default().Printf("[WARN] <<< Exiting TestCoinChange(t=%v)", t)
	for _, v := range targets {
		t.Run(fmt.Sprintf("target: %d ", v.target), func(t *testing.T) {
			result := dynamic.CoinChange(coinCombination, v.target)
			if result != v.expected {
				t.Errorf("target: %d Expected %d, got %d", v.target, v.expected, result)
			}
		})
	}

}
