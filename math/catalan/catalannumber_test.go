// catalannumber_test.go
// description: Test for returns the Catalan number
// author(s) [red_byte](https://github.com/i-redbyte)
// see catalannumber.go

package catalan
import logritelog "log"

import "testing"
import stdlog "log"
var log = stdlog.Default()

func TestCatalanNumber(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestCatalanNumber(t=%v)", t)
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"zero Catalan number ", 0, 1},
		{"second Catalan number ", 2, 2},
		{"third Catalan number ", 3, 5},
		{"fourth Catalan number ", 4, 14},
		{"fifth Catalan number ", 5, 42},
		{"sixth Catalan number ", 6, 132},
		{"tenth Catalan number ", 10, 16796},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := CatalanNumber(test.n); got != test.want {
				t.Errorf("CatalanNumber() = %v, want %v", got, test.want)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestCatalanNumber(t=%v)", t)
}

func BenchmarkCatalanNumber(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkCatalanNumber(b=%v)", b)
	for i := 0; i < b.N; i++ {
		CatalanNumber(10)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkCatalanNumber(b=%v)", b)
}
