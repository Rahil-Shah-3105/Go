package sort_test
import logritelog "log"

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/TheAlgorithms/Go/sort"
)

func testFramework(t *testing.T, sortingFunction func([]int) []int) {
	logritelog.Default().Printf("[ERROR] >>> Entering testFramework(t=%v,sortingFunction=%v)", t, sortingFunction)
	sortTests := []struct {
		input    []int
		expected []int
		name     string
	}{
		//Sorted slice
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			name:     "Sorted Unsigned",
		},
		//Reversed slice
		{
			input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			name:     "Reversed Unsigned",
		},
		//Sorted slice
		{
			input:    []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			name:     "Sorted Signed",
		},
		//Reversed slice
		{
			input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
			expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			name:     "Reversed Signed",
		},
		//Reversed slice, even length
		{
			input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
			expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			name:     "Reversed Signed #2",
		},
		//Random order with repetitions
		{
			input:    []int{-5, 7, 4, -2, 6, 5, 8, 3, 2, -7, -1, 0, -3, 9, -6, -4, 10, 9, 1, -8, -9, -10},
			expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 10},
			name:     "Random order Signed",
		},
		//Single-entry slice
		{
			input:    []int{1},
			expected: []int{1},
			name:     "Singleton",
		},
		// Empty slice
		{
			input:    []int{},
			expected: []int{},
			name:     "Empty Slice",
		},
	}
	for _, test := range sortTests {
		t.Run(test.name, func(t *testing.T) {
			actual := sortingFunction(test.input)
			sorted := reflect.DeepEqual(actual, test.expected)
			if !sorted {
				t.Errorf("test %s failed", test.name)
				t.Errorf("actual %v expected %v", actual, test.expected)
			}
		})
	}
	logritelog.Default().Printf("[WARN] <<< Exiting testFramework(t=%v,sortingFunction=%v)", t, sortingFunction)
}

// BEGIN TESTS
func TestBinaryInsertion(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestBinaryInsertion(t=%v)", t)
	testFramework(t, sort.BinaryInsertion[int])
	logritelog.Default().Printf("[WARN] <<< Exiting TestBinaryInsertion(t=%v)", t)
}

func TestBubble(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestBubble(t=%v)", t)
	testFramework(t, sort.Bubble[int])
	logritelog.Default().Printf("[WARN] <<< Exiting TestBubble(t=%v)", t)
}

func TestBogo(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestBogo(t=%v)", t)
	t.Skip("Skipping test for Bogo Sort, as it uses a lot of resource.")
	testFramework(t, sort.Bogo[int])
	logritelog.Default().Printf("[WARN] <<< Exiting TestBogo(t=%v)", t)
}

func TestBucketSort(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestBucketSort(t=%v)", t)
	testFramework(t, sort.Bucket[int])
	logritelog.Default().Printf("[WARN] <<< Exiting TestBucketSort(t=%v)", t)
}

func TestCocktailSort(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestCocktailSort(t=%v)", t)
	testFramework(t, sort.Cocktail[int])
	logritelog.Default().Printf("[WARN] <<< Exiting TestCocktailSort(t=%v)", t)
}

func TestExchange(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestExchange(t=%v)", t)
	testFramework(t, sort.Exchange[int])
	logritelog.Default().Printf("[WARN] <<< Exiting TestExchange(t=%v)", t)
}

func TestInsertion(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestInsertion(t=%v)", t)
	testFramework(t, sort.Insertion[int])
	logritelog.Default().Printf("[WARN] <<< Exiting TestInsertion(t=%v)", t)
}

func TestMerge(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestMerge(t=%v)", t)
	testFramework(t, sort.Merge[int])
	logritelog.Default().Printf("[WARN] <<< Exiting TestMerge(t=%v)", t)
}

func TestMergeIter(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestMergeIter(t=%v)", t)
	testFramework(t, sort.MergeIter[int])
	logritelog.Default().Printf("[WARN] <<< Exiting TestMergeIter(t=%v)", t)
}

func TestMergeParallel(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestMergeParallel(t=%v)", t)
	testFramework(t, sort.ParallelMerge[int])

	// Test parallel merge sort with a large slice
	t.Run("ParallelMerge on large slice", func(t *testing.T) {
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		size := 100000
		randomLargeSlice := make([]int, size)
		for i := range randomLargeSlice {
			randomLargeSlice[i] = rnd.Intn(size)
		}
		sortedSlice := sort.ParallelMerge[int](randomLargeSlice)
		for i := 0; i < len(sortedSlice)-1; i++ {
			if sortedSlice[i] > sortedSlice[i+1] {
				t.Errorf("ParallelMerge failed")
			}
		}
	})
	logritelog.Default().Printf("[WARN] <<< Exiting TestMergeParallel(t=%v)", t)
}

func TestHeap(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestHeap(t=%v)", t)
	testFramework(t, sort.HeapSort[int])
	logritelog.Default().Printf("[WARN] <<< Exiting TestHeap(t=%v)", t)
}

func TestCount(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestCount(t=%v)", t)
	testFramework(t, sort.Count[int])
	logritelog.Default().Printf("[WARN] <<< Exiting TestCount(t=%v)", t)
}

func TestQuick(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestQuick(t=%v)", t)
	testFramework(t, sort.Quicksort[int])
	logritelog.Default().Printf("[WARN] <<< Exiting TestQuick(t=%v)", t)
}

func TestShell(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestShell(t=%v)", t)
	testFramework(t, sort.Shell[int])
	logritelog.Default().Printf("[WARN] <<< Exiting TestShell(t=%v)", t)
}

func TestRadix(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestRadix(t=%v)", t)
	testFramework(t, sort.RadixSort[int])
	logritelog.Default().Printf("[WARN] <<< Exiting TestRadix(t=%v)", t)
}

func TestSimple(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestSimple(t=%v)", t)
	testFramework(t, sort.Simple[int])
	logritelog.Default().Printf("[WARN] <<< Exiting TestSimple(t=%v)", t)
}

func TestImprovedSimple(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestImprovedSimple(t=%v)", t)
	testFramework(t, sort.ImprovedSimple[int])
	logritelog.Default().Printf("[WARN] <<< Exiting TestImprovedSimple(t=%v)", t)
}

func TestSelection(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestSelection(t=%v)", t)
	testFramework(t, sort.Selection[int])
	logritelog.Default().Printf("[WARN] <<< Exiting TestSelection(t=%v)", t)
}

func TestComb(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestComb(t=%v)", t)
	testFramework(t, sort.Comb[int])
	logritelog.Default().Printf("[WARN] <<< Exiting TestComb(t=%v)", t)
}

func TestPancakeSort(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestPancakeSort(t=%v)", t)
	testFramework(t, sort.Pancake[int])
	logritelog.Default().Printf("[WARN] <<< Exiting TestPancakeSort(t=%v)", t)
}

func TestPigeonhole(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestPigeonhole(t=%v)", t)
	testFramework(t, sort.Pigeonhole[int])
	logritelog.Default().Printf("[WARN] <<< Exiting TestPigeonhole(t=%v)", t)
}

func TestPatience(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestPatience(t=%v)", t)
	testFramework(t, sort.Patience[int])
	logritelog.Default().Printf("[WARN] <<< Exiting TestPatience(t=%v)", t)
}

func TestCycle(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestCycle(t=%v)", t)
	testFramework(t, sort.Cycle[int])
	logritelog.Default().Printf("[WARN] <<< Exiting TestCycle(t=%v)", t)
}

func TestTimsort(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestTimsort(t=%v)", t)
	testFramework(t, sort.Timsort[int])
	logritelog.Default().Printf("[WARN] <<< Exiting TestTimsort(t=%v)", t)
}

func TestCircle(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestCircle(t=%v)", t)
	testFramework(t, sort.Circle[int])
	logritelog.Default().Printf("[WARN] <<< Exiting TestCircle(t=%v)", t)
}

func TestOddEvenSort(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestOddEvenSort(t=%v)", t)
	testFramework(t, sort.OddEvenSort[int])
	logritelog.Default().Printf("[WARN] <<< Exiting TestOddEvenSort(t=%v)", t)
}

func TestStooge(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestStooge(t=%v)", t)
	testFramework(t, sort.Stooge[int])
	logritelog.Default().Printf("[WARN] <<< Exiting TestStooge(t=%v)", t)
}

// END TESTS

func benchmarkFramework(b *testing.B, f func(arr []int) []int) {
	logritelog.Default().Printf("[ERROR] >>> Entering benchmarkFramework(b=%v,f=%v)", b, f)
	var sortTests = []struct {
		input    []int
		expected []int
		name     string
	}{
		//Sorted slice
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Sorted Unsigned"},
		//Reversed slice
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Reversed Unsigned"},
		//Sorted slice
		{[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Sorted Signed"},
		//Reversed slice
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
			[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Reversed Signed"},
		//Reversed slice, even length
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
			[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "Reversed Signed #2"},
		//Random order with repetitions
		{[]int{-5, 7, 4, -2, 6, 5, 8, 3, 2, -7, -1, 0, -3, 9, -6, -4, 10, 9, 1, -8, -9, -10},
			[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 10}, "Random order Signed"},
		//Empty slice
		{[]int{}, []int{}, "Empty"},
		//Single-entry slice
		{[]int{1}, []int{1}, "Singleton"},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range sortTests {
			f(test.input)
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting benchmarkFramework(b=%v,f=%v)", b, f)
}

//BEGIN BENCHMARKS

func BenchmarkBinaryInsertion(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkBinaryInsertion(b=%v)", b)
	benchmarkFramework(b, sort.BinaryInsertion[int])
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkBinaryInsertion(b=%v)", b)
}

func BenchmarkBubble(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkBubble(b=%v)", b)
	benchmarkFramework(b, sort.Bubble[int])
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkBubble(b=%v)", b)
}

func BenchmarkBogo(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkBogo(b=%v)", b)
	b.Skip("Skipping benchmark for Bogo Sort, as it uses a lot of resource.")
	benchmarkFramework(b, sort.Bogo[int])
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkBogo(b=%v)", b)
}

func BenchmarkBucketSort(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkBucketSort(b=%v)", b)
	benchmarkFramework(b, sort.Bucket[int])
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkBucketSort(b=%v)", b)
}

func BenchmarkCocktailSort(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkCocktailSort(b=%v)", b)
	benchmarkFramework(b, sort.Cocktail[int])
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkCocktailSort(b=%v)", b)
}

func BenchmarkExchange(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkExchange(b=%v)", b)
	benchmarkFramework(b, sort.Exchange[int])
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkExchange(b=%v)", b)
}

func BenchmarkInsertion(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkInsertion(b=%v)", b)
	benchmarkFramework(b, sort.Insertion[int])
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkInsertion(b=%v)", b)
}

func BenchmarkMerge(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkMerge(b=%v)", b)
	benchmarkFramework(b, sort.Merge[int])
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkMerge(b=%v)", b)
}

func BenchmarkMergeIter(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkMergeIter(b=%v)", b)
	benchmarkFramework(b, sort.MergeIter[int])
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkMergeIter(b=%v)", b)
}

func BenchmarkMergeParallel(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkMergeParallel(b=%v)", b)
	benchmarkFramework(b, sort.ParallelMerge[int])
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkMergeParallel(b=%v)", b)
}

func BenchmarkHeap(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkHeap(b=%v)", b)
	benchmarkFramework(b, sort.HeapSort[int])
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkHeap(b=%v)", b)
}

func BenchmarkCount(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkCount(b=%v)", b)
	benchmarkFramework(b, sort.Count[int])
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkCount(b=%v)", b)
}

func BenchmarkQuick(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkQuick(b=%v)", b)
	benchmarkFramework(b, sort.Quicksort[int])
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkQuick(b=%v)", b)
}

func BenchmarkShell(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkShell(b=%v)", b)
	benchmarkFramework(b, sort.Shell[int])
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkShell(b=%v)", b)
}

func BenchmarkRadix(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkRadix(b=%v)", b)
	benchmarkFramework(b, sort.RadixSort[int])
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkRadix(b=%v)", b)
}

func BenchmarkSimple(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkSimple(b=%v)", b)
	benchmarkFramework(b, sort.Simple[int])
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkSimple(b=%v)", b)
}

func BenchmarkImprovedSimple(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkImprovedSimple(b=%v)", b)
	benchmarkFramework(b, sort.ImprovedSimple[int])
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkImprovedSimple(b=%v)", b)
}

// Very Slow, consider commenting
func BenchmarkSelection(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkSelection(b=%v)", b)
	benchmarkFramework(b, sort.Selection[int])
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkSelection(b=%v)", b)
}

func BenchmarkComb(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkComb(b=%v)", b)
	benchmarkFramework(b, sort.Comb[int])
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkComb(b=%v)", b)
}

func BenchmarkPancakeSort(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkPancakeSort(b=%v)", b)
	benchmarkFramework(b, sort.Pancake[int])
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkPancakeSort(b=%v)", b)
}

func BenchmarkPigeonhole(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkPigeonhole(b=%v)", b)
	benchmarkFramework(b, sort.Pigeonhole[int])
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkPigeonhole(b=%v)", b)
}

func BenchmarkPatience(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkPatience(b=%v)", b)
	benchmarkFramework(b, sort.Patience[int])
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkPatience(b=%v)", b)
}

func BenchmarkCycle(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkCycle(b=%v)", b)
	benchmarkFramework(b, sort.Cycle[int])
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkCycle(b=%v)", b)
}

func BenchmarkTimsort(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkTimsort(b=%v)", b)
	benchmarkFramework(b, sort.Timsort[int])
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkTimsort(b=%v)", b)
}

func BenchmarkCircle(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkCircle(b=%v)", b)
	benchmarkFramework(b, sort.Circle[int])
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkCircle(b=%v)", b)
}

func BenchmarkStooge(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkStooge(b=%v)", b)
	benchmarkFramework(b, sort.Stooge[int])
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkStooge(b=%v)", b)
}
