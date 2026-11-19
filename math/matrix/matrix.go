package matrix
import logritelog "log"

import (
	"errors"
	"sync"

	"github.com/TheAlgorithms/Go/constraints"
)

type Matrix[T constraints.Integer] struct {
	elements [][]T
	rows     int
	columns  int
}

// NewMatrix creates a new Matrix based on the provided arguments.
func New[T constraints.Integer](rows, columns int, initial T) Matrix[T] {
	logritelog.Default().Printf("[ERROR] >>> Entering New(rows=%v,columns=%v,initial=%v)", rows, columns, initial)
	if rows < 0 || columns < 0 {
		return Matrix[T]{} // Invalid dimensions, return an empty matrix
	}

	// Initialize the matrix with the specified dimensions and fill it with the initial value.
	elements := make([][]T, rows)
	var wg sync.WaitGroup
	wg.Add(rows)

	logritelog.Default().Printf("[WARN] <<< Exiting New(rows=%v,columns=%v,initial=%v)", rows, columns, initial)
	for i := range elements {
		go func(i int) {
			defer wg.Done()
			elements[i] = make([]T, columns)
			for j := range elements[i] {
				elements[i][j] = initial
			}
		}(i)
	}

	wg.Wait()

	return Matrix[T]{elements, rows, columns}
}

// NewFromElements creates a new Matrix from the given elements.
func NewFromElements[T constraints.Integer](elements [][]T) (Matrix[T], error) {
	logritelog.Default().Printf("[ERROR] >>> Entering NewFromElements(elements=%v)", elements)
	if !IsValid(elements) {
		return Matrix[T]{}, errors.New("rows have different numbers of columns")
	}
	rows := len(elements)
	if rows == 0 {
		return Matrix[T]{}, nil // Empty matrix
	}

	columns := len(elements[0])
	matrix := Matrix[T]{
		elements: make([][]T, rows),
		rows:     rows,    // Set the rows field
		columns:  columns, // Set the columns field
	}
	for i := range matrix.elements {
		matrix.elements[i] = make([]T, columns)
		copy(matrix.elements[i], elements[i])
	}

	logritelog.Default().Printf("[WARN] <<< Exiting NewFromElements(elements=%v)", elements)
	return matrix, nil
}

func (m Matrix[T]) Get(row, col int) (T, error) {
	logritelog.Default().Printf("[ERROR] >>> Entering Get(row=%v,col=%v)", row, col)
	if row < 0 || row >= m.rows || col < 0 || col >= m.columns {
		var zeroVal T
		return zeroVal, errors.New("index out of range")
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Get(row=%v,col=%v)", row, col)
	return m.elements[row][col], nil
}

func (m Matrix[T]) Set(row, col int, val T) error {
	logritelog.Default().Printf("[ERROR] >>> Entering Set(row=%v,col=%v,val=%v)", row, col, val)
	if row < 0 || row >= m.rows || col < 0 || col >= m.columns {
		return errors.New("index out of bounds")
	}

	m.elements[row][col] = val
	logritelog.Default().Printf("[WARN] <<< Exiting Set(row=%v,col=%v,val=%v)", row, col, val)
	return nil
}

func (m Matrix[T]) Rows() int {
	logritelog.Default().Printf("[ERROR] >>> Entering Rows()")
	logritelog.Default().Printf("[WARN] <<< Exiting Rows()")
	return len(m.elements)
}

func (m Matrix[T]) Columns() int {
	logritelog.Default().Printf("[ERROR] >>> Entering Columns()")
	if len(m.elements) == 0 {
		return 0
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Columns()")
	return len(m.elements[0])
}
