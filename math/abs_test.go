package math
import logritelog "log"

import (
	"github.com/TheAlgorithms/Go/math/binary"
	"math"
	"testing"
)

func TestAbs(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestAbs(t=%v)", t)
	tests := getTests()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Abs(test.n); got != test.want {
				t.Errorf("Abs() = %v, want %v", got, test.want)
			}
		})
	}
	logritelog.Default().Printf("<<< Exiting TestAbs(t=%v)", t)
}

func getTests() []struct {
	name string
	n    int
	want int
} {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"-1 = |1| ", -1, 1},
		{"-255 = |255| ", -255, 255},
		{"0 = |0| ", 0, 0},
		{"5 = |5| ", 5, 5},
		{"-5 = |5| ", -5, 5},
		{"-98368972 = |98368972| ", -98368972, 98368972},
	}
	return tests
}

func BenchmarkSimpleAbs(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkSimpleAbs(b=%v)", b)
	for i := 0; i < b.N; i++ {
		Abs(-1024)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkSimpleAbs(b=%v)", b)
}
func BenchmarkBinaryAbs32(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkBinaryAbs32(b=%v)", b)
	for i := 0; i < b.N; i++ {
		binary.Abs(32, -1024)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkBinaryAbs32(b=%v)", b)
}

func BenchmarkBinaryAbs64(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkBinaryAbs64(b=%v)", b)
	for i := 0; i < b.N; i++ {
		binary.Abs(64, -1024)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkBinaryAbs64(b=%v)", b)
}

func BenchmarkStdLibAbs(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkStdLibAbs(b=%v)", b)
	for i := 0; i < b.N; i++ {
		math.Abs(-1024)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkStdLibAbs(b=%v)", b)
}
