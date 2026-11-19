package heap
import logritelog "log"

import (
	"errors"
	"github.com/TheAlgorithms/Go/constraints"
)

// Heap heap implementation using generic.
type Heap[T any] struct {
	heaps    []T
	lessFunc func(a, b T) bool
}

// New gives a new heap object.
func New[T constraints.Ordered]() *Heap[T] {
	logritelog.Default().Printf("[ERROR] >>> Entering New()")
	less := func(a, b T) bool {
		return a < b
	}
	h, _ := NewAny[T](less)
	logritelog.Default().Printf("[WARN] <<< Exiting New()")
	return h
}

// NewAny gives a new heap object. element can be anything, but must provide less function.
func NewAny[T any](less func(a, b T) bool) (*Heap[T], error) {
	logritelog.Default().Printf("[ERROR] >>> Entering NewAny(less=%v)", less)
	if less == nil {
		return nil, errors.New("less func is necessary")
	}
	logritelog.Default().Printf("[WARN] <<< Exiting NewAny(less=%v)", less)
	return &Heap[T]{
		lessFunc: less,
	}, nil
}

// Push pushes the element t onto the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap[T]) Push(t T) {
	logritelog.Default().Printf("[ERROR] >>> Entering Push(t=%v)", t)
	h.heaps = append(h.heaps, t)
	h.up(len(h.heaps) - 1)
	logritelog.Default().Printf("[WARN] <<< Exiting Push(t=%v)", t)
}

// Top returns the minimum element (according to Less) from the heap.
// Top panics if the heap is empty.
func (h *Heap[T]) Top() T {
	logritelog.Default().Printf("[ERROR] >>> Entering Top()")
	logritelog.Default().Printf("[WARN] <<< Exiting Top()")
	return h.heaps[0]
}

// Pop removes the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap[T]) Pop() {
	logritelog.Default().Printf("[ERROR] >>> Entering Pop()")
	if len(h.heaps) <= 1 {
		h.heaps = nil
		logritelog.Default().Printf("[DEBUG] Pop(): h.heaps → %v", h.heaps)
		return
	}
	h.swap(0, len(h.heaps)-1)
	h.heaps = h.heaps[:len(h.heaps)-1]
	h.down(0)
	logritelog.Default().Printf("[WARN] <<< Exiting Pop()")
}

// Empty returns the heap is empty or not.
func (h *Heap[T]) Empty() bool {
	logritelog.Default().Printf("[ERROR] >>> Entering Empty()")
	logritelog.Default().Printf("[WARN] <<< Exiting Empty()")
	return len(h.heaps) == 0
}

// Size returns the size of the heap
func (h *Heap[T]) Size() int {
	logritelog.Default().Printf("[ERROR] >>> Entering Size()")
	logritelog.Default().Printf("[WARN] <<< Exiting Size()")
	return len(h.heaps)
}

func (h *Heap[T]) swap(i, j int) {
	logritelog.Default().Printf("[ERROR] >>> Entering swap(i=%v,j=%v)", i, j)
	h.heaps[i], h.heaps[j] = h.heaps[j], h.heaps[i]
	logritelog.Default().Printf("[WARN] <<< Exiting swap(i=%v,j=%v)", i, j)
}

func (h *Heap[T]) up(child int) {
	logritelog.Default().Printf("[ERROR] >>> Entering up(child=%v)", child)
	if child <= 0 {
		return
	}
	parent := (child - 1) >> 1
	if !h.lessFunc(h.heaps[child], h.heaps[parent]) {
		return
	}
	h.swap(child, parent)
	h.up(parent)
	logritelog.Default().Printf("[WARN] <<< Exiting up(child=%v)", child)
}

func (h *Heap[T]) down(parent int) {
	logritelog.Default().Printf("[ERROR] >>> Entering down(parent=%v)", parent)
	lessIdx := parent
	lChild, rChild := (parent<<1)+1, (parent<<1)+2
	if lChild < len(h.heaps) && h.lessFunc(h.heaps[lChild], h.heaps[lessIdx]) {
		lessIdx = lChild
	}
	if rChild < len(h.heaps) && h.lessFunc(h.heaps[rChild], h.heaps[lessIdx]) {
		lessIdx = rChild
	}
	if lessIdx == parent {
		return
	}
	h.swap(lessIdx, parent)
	h.down(lessIdx)
	logritelog.Default().Printf("[WARN] <<< Exiting down(parent=%v)", parent)
}
