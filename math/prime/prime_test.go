package prime
import logritelog "log"

import (
	"fmt"
	"testing"
)

var primeList = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127}
var testLimit = 127

type primalityTest func(int64) bool

func primalityTestTestingHelper(t *testing.T, name string, f primalityTest) {
	logritelog.Default().Printf(">>> Entering primalityTestTestingHelper(t=%v,name=%v,f=%v)", t, name, f)
	arrayIndex := 0
	for i := 1; i <= testLimit; i++ {
		isPrime := i == primeList[arrayIndex]

		testName := fmt.Sprintf("%s(%d)", name, i)
		t.Run(testName, func(t *testing.T) {
			result := f(int64(i))

			if isPrime {
				arrayIndex++
			}

			if result != isPrime {
				t.Errorf("%d: %s function returned %v\n", i, name, result)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting primalityTestTestingHelper(t=%v,name=%v,f=%v)", t, name, f)
}

func primalityTestBenchmarkHelper(b *testing.B, f primalityTest) {
	logritelog.Default().Printf(">>> Entering primalityTestBenchmarkHelper(b=%v,f=%v)", b, f)
	for i := 0; i < b.N; i++ {
		f(104729)
	}
	logritelog.Default().Printf("<<< Exiting primalityTestBenchmarkHelper(b=%v,f=%v)", b, f)
}

// Miller-Rabin Probabilistic test

func millerRabinProbabilisticTester(n int64) bool {
	logritelog.Default().Printf(">>> Entering millerRabinProbabilisticTester(n=%v)", n)
	result, err := MillerRabinProbabilistic(n, 40)
	if err != nil {
		logritelog.Default().Printf("[ERROR] Exception in millerRabinProbabilisticTester(n=%v): %v", n, err)
		panic(err)
	}

	logritelog.Default().Printf("<<< Exiting millerRabinProbabilisticTester(n=%v)", n)
	return result
}

func TestMillerRabinProbabilistic(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestMillerRabinProbabilistic(t=%v)", t)
	primalityTestTestingHelper(t, "Miller-Rabin Probabilistic", millerRabinProbabilisticTester)
	logritelog.Default().Printf("<<< Exiting TestMillerRabinProbabilistic(t=%v)", t)
}

func BenchmarkMillerRabinProbabilistic(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkMillerRabinProbabilistic(b=%v)", b)
	primalityTestBenchmarkHelper(b, millerRabinProbabilisticTester)
	logritelog.Default().Printf("<<< Exiting BenchmarkMillerRabinProbabilistic(b=%v)", b)
}

// Miller-Rabin deterministic test

func millerRabinDeterministicTester(n int64) bool {
	logritelog.Default().Printf(">>> Entering millerRabinDeterministicTester(n=%v)", n)
	result, err := MillerRabinDeterministic(n)
	if err != nil {
		logritelog.Default().Printf("[ERROR] Exception in millerRabinDeterministicTester(n=%v): %v", n, err)
		panic(err)
	}

	logritelog.Default().Printf("<<< Exiting millerRabinDeterministicTester(n=%v)", n)
	return result
}

func TestMillerRabinDeterministic(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestMillerRabinDeterministic(t=%v)", t)
	primalityTestTestingHelper(t, "Miller-Rabin Deterministic", millerRabinDeterministicTester)
	logritelog.Default().Printf("<<< Exiting TestMillerRabinDeterministic(t=%v)", t)
}

func BenchmarkMillerRabinDeterministic(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkMillerRabinDeterministic(b=%v)", b)
	primalityTestBenchmarkHelper(b, millerRabinDeterministicTester)
	logritelog.Default().Printf("<<< Exiting BenchmarkMillerRabinDeterministic(b=%v)", b)
}

// Trial Division test

func TestTrialDivision(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestTrialDivision(t=%v)", t)
	primalityTestTestingHelper(t, "Trial Division", TrialDivision)
	logritelog.Default().Printf("<<< Exiting TestTrialDivision(t=%v)", t)
}

func BenchmarkTrialDivision(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkTrialDivision(b=%v)", b)
	primalityTestBenchmarkHelper(b, TrialDivision)
	logritelog.Default().Printf("<<< Exiting BenchmarkTrialDivision(b=%v)", b)
}

// Trial Division (optimized)

func TestOptimizedTrialDivision(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestOptimizedTrialDivision(t=%v)", t)
	primalityTestTestingHelper(t, "Trial Division (optimized)", OptimizedTrialDivision)
	logritelog.Default().Printf("<<< Exiting TestOptimizedTrialDivision(t=%v)", t)
}

func BenchmarkOptimizedTrialDivision(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkOptimizedTrialDivision(b=%v)", b)
	primalityTestBenchmarkHelper(b, OptimizedTrialDivision)
	logritelog.Default().Printf("<<< Exiting BenchmarkOptimizedTrialDivision(b=%v)", b)
}
