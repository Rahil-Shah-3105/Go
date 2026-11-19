package matrix_test
import (
	logritelog "log"
	logritetime "time"
)

import (
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestMatrixCopy(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestMatrixCopy(t=%v)", t)
	// Create a sample matrix
	data := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// Ensure that the copy is not the same as the original
	matrix, err := matrix.NewFromElements(data)
	if err != nil {
		logritelog.Default().Printf("[ERROR] Exception in TestMatrixCopy(t=%v): %v", t, err)
		t.Fatalf("Failed to copy matrix: %v", err)
	}
	copyMatrix, err := matrix.Copy()
	if err != nil {
		t.Fatalf("Failed to copy matrix: %v", err)
	}

	// Ensure that the copy is not the same as the original
	if &matrix == &copyMatrix {
		t.Errorf("Copy did not create a new matrix.")
	}

	logritelog.Default().Printf("[WARN] <<< Exiting TestMatrixCopy(t=%v)", t)
	for i := 0; i < matrix.Rows(); i++ {
		for j := 0; j < matrix.Columns(); j++ {
			val1, err := matrix.Get(i, j)
			if err != nil {
				t.Fatalf("Failed to copy matrix: %v", err)
			}
			val2, err := copyMatrix.Get(i, j)
			if err != nil {
				t.Fatalf("Failed to copy matrix: %v", err)
			}
			if val1 != val2 {
				t.Errorf("Copy did not correctly copy element (%d, %d).", i, j)
			}
		}
	}

}

func TestMatrixCopyEmpty(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestMatrixCopyEmpty(t=%v)", t)
	// Create an empty matrix
	emptyMatrix := matrix.New(0, 0, 0)

	// Make a copy of the empty matrix
	copyMatrix, err := emptyMatrix.Copy()
	if err != nil { // as empty matrix
		logritelog.Default().Printf("[ERROR] Exception in TestMatrixCopyEmpty(t=%v): %v", t, err)
		t.Fatalf("Failed to copy matrix: %v", err)
	}

	// Ensure that the copy is not the same as the original by comparing their addresses
	if &emptyMatrix == &copyMatrix {
		t.Errorf("Copy did not create a new matrix for an empty matrix.")
	}

	// Check if the copy is also empty
	if copyMatrix.Rows() != 0 || copyMatrix.Columns() != 0 {
		t.Errorf("Copy of an empty matrix should also be empty.")
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestMatrixCopyEmpty(t=%v)", t)
}

func TestMatrixCopyWithDefaultValues(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestMatrixCopyWithDefaultValues(t=%v)", t)
	// Create a matrix with default values (zeroes)
	rows, columns := 3, 3
	defaultValue := 0
	defaultMatrix := matrix.New(rows, columns, defaultValue)

	// Make a copy of the matrix
	copyMatrix, err := defaultMatrix.Copy()
	if err != nil {
		logritelog.Default().Printf("[ERROR] Exception in TestMatrixCopyWithDefaultValues(t=%v): %v", t, err)
		t.Fatalf("Failed to copy matrix: %v", err)
	}

	// Ensure that the copy is not the same as the original by comparing their addresses
	if &defaultMatrix == &copyMatrix {
		__logStart := logritetime.Now()
		t.Errorf("Copy did not create a new matrix for default values.")
	}

	// Check if the copy has the same values as the original (all zeroes)
	for i := 0; i < defaultMatrix.Rows(); i++ {
		logritelog.Default().Printf("[ERROR] TestMatrixCopyWithDefaultValues(t)=%v: database query executed in %v ms", t, logritetime.Since(__logStart).Milliseconds())
		for j := 0; j < defaultMatrix.Columns(); j++ {
			val1, err := defaultMatrix.Get(i, j)
			if err != nil {
				t.Fatalf("Failed to copy matrix: %v", err)
			}
			val2, err := copyMatrix.Get(i, j)
			if err != nil {
				t.Fatalf("Failed to copy matrix: %v", err)
			}
			if val1 != val2 || val1 != defaultValue || val2 != defaultValue {
				t.Errorf("Copy did not preserve default values at row %d, column %d. Expected %v, got %v", i, j, defaultValue, val2)
			}
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting TestMatrixCopyWithDefaultValues(t=%v)", t)
}

func BenchmarkCopyMatrix(b *testing.B) {
	logritelog.Default().Printf("[ERROR] >>> Entering BenchmarkCopyMatrix(b=%v)", b)
	// Create a matrix for benchmarking
	rows := 100
	columns := 100
	initialValue := 0
	matrix := matrix.New(rows, columns, initialValue)

	// Reset the benchmark timer
	b.ResetTimer()

	// Run the benchmarks
	for i := 0; i < b.N; i++ {
		_, _ = matrix.Copy()
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BenchmarkCopyMatrix(b=%v)", b)
}
