package matrix_test
import logritelog "log"

import (
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestCheckEqual(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestCheckEqual(t=%v)", t)
	// Create two matrices with the same dimensions and equal values
	m1 := matrix.New(2, 2, 0)
	m2 := matrix.New(2, 2, 0)

	// Test case 1: Matrices are equal
	equal := m1.CheckEqual(m2)

	if !equal {
		t.Errorf("CheckEqual(m1, m2) returned false, expected true (matrices are equal)")
	}

	// Create two matrices with the same dimensions but different values
	m3 := matrix.New(2, 2, 1)
	m4 := matrix.New(2, 2, 0)

	// Test case 2: Matrices are not equal
	equal2 := m3.CheckEqual(m4)
	if equal2 {
		t.Errorf("CheckEqual(m3, m4) returned true, expected false (matrices are not equal)")
	}

	// Create two matrices with different dimensions
	m5 := matrix.New(2, 2, 0)
	m6 := matrix.New(2, 3, 0)

	// Test case 3: Matrices have different dimensions
	equal3 := m5.CheckEqual(m6)

	if equal3 {
		t.Errorf("CheckEqual(m5, m6) returned true, expected false (matrices are not equal)")
	}
	logritelog.Default().Printf("<<< Exiting TestCheckEqual(t=%v)", t)
}

func BenchmarkCheckEqualSmallMatrix(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkCheckEqualSmallMatrix(b=%v)", b)
	m1 := matrix.New(10, 10, 0) // Create a 10x10 matrix with all zeros
	m2 := matrix.New(10, 10, 0) // Create another 10x10 matrix with all zeros

	for i := 0; i < b.N; i++ {
		_ = m1.CheckEqual(m2)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkCheckEqualSmallMatrix(b=%v)", b)
}

func BenchmarkCheckEqualLargeMatrix(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkCheckEqualLargeMatrix(b=%v)", b)
	size := 1000 // Choose an appropriate size for your large matrix
	m1 := MakeRandomMatrix[int](size, size)
	m2 := MakeRandomMatrix[int](size, size)

	b.ResetTimer() // Reset the timer to exclude setup time

	for i := 0; i < b.N; i++ {
		_ = m1.CheckEqual(m2)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkCheckEqualLargeMatrix(b=%v)", b)
}
