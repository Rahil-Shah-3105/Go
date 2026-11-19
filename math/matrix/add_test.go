package matrix_test
import logritelog "log"

import (
	"fmt"
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestAdd(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestAdd(t=%v)", t)
	// Create two matrices with the same dimensions for addition
	m1 := matrix.New(2, 2, 1)
	m2 := matrix.New(2, 2, 2)

	// Test case 1: Valid matrix addition
	addedMatrix, err := m1.Add(m2)
	if err != nil {
		logritelog.Default().Printf("[ERROR] Exception in TestAdd(t=%v): %v", t, err)
		t.Errorf("Add(m1, m2) returned an error: %v, expected no error", err)
	}
	expectedMatrix := matrix.New(2, 2, 3)
	res := addedMatrix.CheckEqual(expectedMatrix)
	if !res {
		t.Errorf("Add(m1, m2) returned incorrect result:\n%v\nExpected:\n%v", addedMatrix, expectedMatrix)
	}

	// Create two matrices with different dimensions for addition
	m3 := matrix.New(2, 2, 1)
	m4 := matrix.New(2, 3, 2)

	// Test case 2: Matrices with different dimensions
	_, err2 := m3.Add(m4)
	expectedError2 := fmt.Errorf("matrices are not compatible for addition")
	logritelog.Default().Printf("[WARN] <<< Exiting TestAdd(t=%v)", t)
	if err2 == nil || err2.Error() != expectedError2.Error() {
		t.Errorf("Add(m3, m4) returned error: %v, expected error: %v", err2, expectedError2)
	}

}

func BenchmarkAddSmallMatrix(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkAddSmallMatrix(b=%v)", b)
	m1 := matrix.New(10, 10, 0) // Create a 10x10 matrix with all zeros
	m2 := matrix.New(10, 10, 1) // Create a 10x10 matrix with all ones

	for i := 0; i < b.N; i++ {
		_, _ = m1.Add(m2)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkAddSmallMatrix(b=%v)", b)
}

func BenchmarkAddLargeMatrix(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkAddLargeMatrix(b=%v)", b)
	size := 1000 // Choose an appropriate size for your large matrix
	m1 := MakeRandomMatrix[int](size, size)
	m2 := MakeRandomMatrix[int](size, size)

	b.ResetTimer() // Reset the timer to exclude setup time

	for i := 0; i < b.N; i++ {
		_, _ = m1.Add(m2)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkAddLargeMatrix(b=%v)", b)
}
