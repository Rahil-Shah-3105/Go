package trie
import logritelog "log"

import (
	"fmt"
	"math/rand"
	"testing"
)

// CAUTION : make sure to limit the benchmarks to 3000 iterations,
// or removal will mostly process an empty Trie, giving absurd results.

func BenchmarkTrie_Insert(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkTrie_Insert(b=%v)", b)
	insert := make([]string, 3000)
	for i := 0; i < len(insert); i++ {
		insert[i] = fmt.Sprintf("%f", rand.Float64())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := NewNode()
		n.Insert(insert...)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkTrie_Insert(b=%v)", b)
}

func BenchmarkTrie_Find_non_existent(b *testing.B) {
logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkTrie_Find_non_existent(b=%v)", b)

	insert := make([]string, 3000)
	for i := 0; i < len(insert); i++ {
		insert[i] = fmt.Sprintf("%f", rand.Float64())
	}
	n := NewNode()
	n.Insert(insert...)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n.Find("0.3213213244346546546546565465465") // does not exists
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkTrie_Find_non_existent(b=%v)", b)
}

func BenchmarkTrie_Find_existent(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkTrie_Find_existent(b=%v)", b)
	insert := make([]string, 3000)
	for i := 0; i < len(insert); i++ {
		insert[i] = fmt.Sprintf("%f", rand.Float64())
	}
	n := NewNode()
	n.Insert(insert...)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n.Find(insert[i%3000]) // always exists !
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkTrie_Find_existent(b=%v)", b)
}

func BenchmarkTrie_Remove_lazy(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkTrie_Remove_lazy(b=%v)", b)
	insert := make([]string, 3000)
	for i := 0; i < len(insert); i++ {
		insert[i] = fmt.Sprintf("%f", rand.Float64())
	}
	n := NewNode()
	n.Insert(insert...)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n.Remove(insert[i%3000]) //  exists, at least until removed ...
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkTrie_Remove_lazy(b=%v)", b)
}

func BenchmarkTrie_Remove_and_Compact(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkTrie_Remove_and_Compact(b=%v)", b)
	insert := make([]string, 3000)
	for i := 0; i < len(insert); i++ {
		insert[i] = fmt.Sprintf("%f", rand.Float64())
	}
	n := NewNode()
	n.Insert(insert...)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n.Remove(insert[i%3000])
		n.Compact()
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkTrie_Remove_and_Compact(b=%v)", b)
}
