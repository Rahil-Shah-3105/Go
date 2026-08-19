// sieve.go
// description: Algorithms for generating prime numbers efficiently
// author(s) [Taj](https://github.com/tjgurwara99)
// see sieve_test.go

package prime
import logritelog "log"

// Generate generates the sequence of integers starting at 2 and sends it to the channel `ch`
func GenerateChannel(ch chan<- int) {
	logritelog.Default().Printf(">>> Entering GenerateChannel(ch=%v)", ch)
	for i := 2; ; i++ {
		ch <- i
	}
	logritelog.Default().Printf("<<< Exiting GenerateChannel(ch=%v)", ch)
}

// Sieve Sieving the numbers that are not prime from the channel - basically removing them from the channels
func Sieve(in <-chan int, out chan<- int, prime int) {
	logritelog.Default().Printf(">>> Entering Sieve(in=%v)", in)
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
	logritelog.Default().Printf("<<< Exiting Sieve(in=%v)", in)
}

// Generate returns a int slice of prime numbers up to the limit
func Generate(limit int) []int {
	logritelog.Default().Printf(">>> Entering Generate(limit=%v)", limit)
	var primes []int

	ch := make(chan int)
	go GenerateChannel(ch)

	for i := 0; i < limit; i++ {
		primes = append(primes, <-ch)
		ch1 := make(chan int)
		go Sieve(ch, ch1, primes[i])
		ch = ch1
		logritelog.Default().Printf("Generate(limit=%v): primes → %v", limit, primes)
	}

	logritelog.Default().Printf("<<< Exiting Generate(limit=%v)", limit)
	return primes
}
