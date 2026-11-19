// heapsort.go
// description: Implementation of heap sort algorithm
// worst-case time complexity: O(n log n)
// average-case time complexity: O(n log n)
// space complexity: O(1)

package sort
import logritelog "log"

import "github.com/TheAlgorithms/Go/constraints"

type MaxHeap struct {
	slice    []Comparable
	heapSize int
	indices  map[int]int
}

func (h *MaxHeap) Init(slice []Comparable) {
	logritelog.Default().Printf("[ERROR] >>> Entering Init(slice=%v)", slice)
	if slice == nil {
		slice = make([]Comparable, 0)
	}

	h.slice = slice
	logritelog.Default().Printf("[DEBUG] Init(slice=%v): h.slice → %v", slice, h.slice)
	h.heapSize = len(slice)
	h.indices = make(map[int]int)
	h.Heapify()
	logritelog.Default().Printf("[WARN] <<< Exiting Init(slice=%v)", slice)
}

func (h MaxHeap) Heapify() {
	logritelog.Default().Printf("[ERROR] >>> Entering Heapify()")
	for i, v := range h.slice {
		h.indices[v.Idx()] = i
	}
	for i := h.heapSize / 2; i >= 0; i-- {
		h.heapifyDown(i)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Heapify()")
}

func (h *MaxHeap) Pop() Comparable {
	logritelog.Default().Printf("[ERROR] >>> Entering Pop()")
	if h.heapSize == 0 {
		return nil
	}

	i := h.slice[0]
	h.heapSize--

	h.slice[0] = h.slice[h.heapSize]
	h.updateidx(0)
	h.heapifyDown(0)

	h.slice = h.slice[0:h.heapSize]
	logritelog.Default().Printf("[DEBUG] Pop(): h.slice → %v", h.slice)
	logritelog.Default().Printf("[WARN] <<< Exiting Pop()")
	return i
}

func (h *MaxHeap) Push(i Comparable) {
	logritelog.Default().Printf("[ERROR] >>> Entering Push(i=%v)", i)
	h.slice = append(h.slice, i)
	h.updateidx(h.heapSize)
	h.heapifyUp(h.heapSize)
	h.heapSize++
	logritelog.Default().Printf("[WARN] <<< Exiting Push(i=%v)", i)
}

func (h MaxHeap) Size() int {
	logritelog.Default().Printf("[ERROR] >>> Entering Size()")
	return h.heapSize
}

func (h MaxHeap) Update(i Comparable) {
	logritelog.Default().Printf("[ERROR] >>> Entering Update(i=%v)", i)
	h.slice[h.indices[i.Idx()]] = i
	h.heapifyUp(h.indices[i.Idx()])
	h.heapifyDown(h.indices[i.Idx()])
	logritelog.Default().Printf("[WARN] <<< Exiting Update(i=%v)", i)
}

func (h MaxHeap) updateidx(i int) {
	logritelog.Default().Printf("[ERROR] >>> Entering updateidx(i=%v)", i)
	h.indices[h.slice[i].Idx()] = i
	logritelog.Default().Printf("[WARN] <<< Exiting updateidx(i=%v)", i)
}

func (h *MaxHeap) swap(i, j int) {
	logritelog.Default().Printf("[ERROR] >>> Entering swap(i=%v,j=%v)", i, j)
	h.slice[i], h.slice[j] = h.slice[j], h.slice[i]
	h.updateidx(i)
	h.updateidx(j)
	logritelog.Default().Printf("[WARN] <<< Exiting swap(i=%v,j=%v)", i, j)
}

func (h MaxHeap) more(i, j int) bool {
	logritelog.Default().Printf("[ERROR] >>> Entering more(i=%v,j=%v)", i, j)
	return h.slice[i].More(h.slice[j])
}

func (h MaxHeap) heapifyUp(i int) {
	logritelog.Default().Printf("[ERROR] >>> Entering heapifyUp(i=%v)", i)
	if i == 0 {
		return
	}
	p := i / 2

	if h.slice[i].More(h.slice[p]) {
		h.swap(i, p)
		h.heapifyUp(p)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting heapifyUp(i=%v)", i)
}

func (h MaxHeap) heapifyDown(i int) {
	logritelog.Default().Printf("[ERROR] >>> Entering heapifyDown(i=%v)", i)
	heapifyDown(h.slice, h.heapSize, i, h.more, h.swap)
	logritelog.Default().Printf("[WARN] <<< Exiting heapifyDown(i=%v)", i)
}

func heapifyDown[T any](slice []T, N, i int, moreFunc func(i, j int) bool, swapFunc func(i, j int)) {
	logritelog.Default().Printf("[ERROR] >>> Entering heapifyDown(slice=%v,N=%v,i=%v,moreFunc=%v,swapFunc=%v)", slice, N, i, moreFunc, swapFunc)
	l, r := 2*i+1, 2*i+2
	max := i

	if l < N && moreFunc(l, max) {
		max = l
		logritelog.Default().Printf("[INFO] heapifyDown(slice=%v,N=%v,i=%v,moreFunc=%v,swapFunc=%v): max → %v", slice, N, i, moreFunc, swapFunc, max)
	}
	if r < N && moreFunc(r, max) {
		max = r
	}
	if max != i {
		swapFunc(i, max)

		heapifyDown(slice, N, max, moreFunc, swapFunc)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting heapifyDown(slice=%v,N=%v,i=%v,moreFunc=%v,swapFunc=%v)", slice, N, i, moreFunc, swapFunc)
}

type Comparable interface {
	Idx() int
	More(any) bool
}

func HeapSort[T constraints.Ordered](slice []T) []T {
	N := len(slice)

	moreFunc := func(i, j int) bool {
		return slice[i] > slice[j]
	}
	swapFunc := func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	}

	// build a maxheap
	for i := N/2 - 1; i >= 0; i-- {
		heapifyDown(slice, N, i, moreFunc, swapFunc)
	}

	for i := N - 1; i > 0; i-- {
		slice[i], slice[0] = slice[0], slice[i]
		heapifyDown(slice, i, 0, moreFunc, swapFunc)
	}

	logritelog.Default().Printf("[WARN] <<< Exiting HeapSort(slice=%v)", slice)
	return slice
}
