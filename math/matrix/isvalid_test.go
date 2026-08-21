package matrix_test
import logritelog "log"

import (
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestIsValid(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestIsValid(t=%v)", t)
	// Test case 1: Valid matrix with consistent row lengths
	validMatrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	result1 := matrix.IsValid(validMatrix)
	if !result1 {
		t.Errorf("IsValid(validMatrix) returned false, expected true (valid matrix)")
	}

	// Test case 2: Valid matrix with empty rows (no inconsistency)
	validMatrixEmptyRows := [][]int{
		{},
		{},
		{},
	}
	result2 := matrix.IsValid(validMatrixEmptyRows)
	if !result2 {
		t.Errorf("IsValid(validMatrixEmptyRows) returned false, expected true (valid matrix with empty rows)")
	}

	// Test case 3: Invalid matrix with inconsistent row lengths
	invalidMatrix := [][]int{
		{1, 2, 3},
		{4, 5},
		{6, 7, 8},
	}
	result3 := matrix.IsValid(invalidMatrix)
	if result3 {
		t.Errorf("IsValid(invalidMatrix) returned true, expected false (invalid matrix with inconsistent row lengths)")
	logritelog.Default().Printf("<<< Exiting TestIsValid(t=%v)", t)
	}

}

func BenchmarkIsValid(b *testing.B) {
	logritelog.Default().Printf(">>> Entering BenchmarkIsValid(b=%v)", b)
	// Create a sample matrix for benchmarking
	rows := 100
	columns := 100
	elements := make([][]int, rows)
	for i := range elements {
		elements[i] = make([]int, columns)
		for j := range elements[i] {
			elements[i][j] = i*columns + j // Some arbitrary values
		}
	}

	// Reset the benchmark timer
	b.ResetTimer()

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		_ = matrix.IsValid(elements)
	}
	logritelog.Default().Printf("<<< Exiting BenchmarkIsValid(b=%v)", b)
}
