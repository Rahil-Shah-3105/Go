// sieve_test.go
// description: Testing all the algorithms related to the generation of prime numbers in sieve.go
// author(s) [Taj](https://github.com/tjgurwara99)
// see sieve.go

package prime
import logritelog "log"

import (
	"reflect"
	"testing"
)

func TestSieve(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestSieve(t=%v)", t)
	firstTenPrimes := [10]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	t.Run("First 10 primes test", func(t *testing.T) {
		var testTenPrimes [10]int

		ch := make(chan int)
		go GenerateChannel(ch)

		for i := 0; i < 10; i++ {
			testTenPrimes[i] = <-ch
			ch1 := make(chan int)
			go Sieve(ch, ch1, testTenPrimes[i])
			ch = ch1
		}

		if firstTenPrimes != testTenPrimes {
			t.Errorf("The first 10 primes do not match")
		}

	})
	logritelog.Default().Printf("<<< Exiting TestSieve(t=%v)", t)
}

func TestGeneratePrimes(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestGeneratePrimes(t=%v)", t)
	firstTenPrimes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	t.Run("Testing GeneratePrimes Function", func(t *testing.T) {
		testPrimes := Generate(10)

		if !reflect.DeepEqual(firstTenPrimes, testPrimes) {
			t.Fatal("GeneratePrimes function failed")
		}
	})
	logritelog.Default().Printf("<<< Exiting TestGeneratePrimes(t=%v)", t)
}

func BenchmarkSieve10(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkSieve10(b=%v)", b)
	for i := 0; i < b.N; i++ {
		_ = Generate(10)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkSieve10(b=%v)", b)
}
