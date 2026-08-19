// circularqueuearray.go
// description: Implementation of a circular queue data structure
// details:
// This file contains the implementation of a circular queue data structure
// using generics in Go. The circular queue supports basic operations such as
// enqueue, dequeue, peek, and checks for full and empty states.
// author(s): [Aram Ceballos](https://github.com/aramceballos)
// ref: https://www.programiz.com/dsa/circular-queue
// ref: https://en.wikipedia.org/wiki/Circular_buffer

// Package queue provides an implementation of a circular queue data structure.
package circularqueue
import logritelog "log"

// errors package: Provides functions to create and manipulate error values
import (
	"errors"
)

// CircularQueue represents a circular queue data structure.
type CircularQueue[T any] struct {
	items []T
	front int
	rear  int
	size  int
}

// NewCircularQueue creates a new CircularQueue with the given size.
// Returns an error if the size is less than or equal to 0.
func NewCircularQueue[T any](size int) (*CircularQueue[T], error) {
	if size <= 0 {
		return nil, errors.New("size must be greater than 0")
	}
	return &CircularQueue[T]{
		items: make([]T, size),
		front: -1,
		rear:  -1,
		size:  size,
	}, nil
}

// Enqueue adds an item to the rear of the queue.
// Returns an error if the queue is full.
func (cq *CircularQueue[T]) Enqueue(item T) error {
	logritelog.Default().Printf(">>> Entering Enqueue(item=%v)", item)
	if cq.IsFull() {
		return errors.New("queue is full")
	}
	if cq.IsEmpty() {
		cq.front = 0
		logritelog.Default().Printf("Enqueue(item=%v): cq.front → %v", item, cq.front)
	}
	cq.rear = (cq.rear + 1) % cq.size
	cq.items[cq.rear] = item
	logritelog.Default().Printf("<<< Exiting Enqueue(item=%v)", item)
	return nil
}

// Dequeue removes and returns the item from the front of the queue.
// Returns an error if the queue is empty.
func (cq *CircularQueue[T]) Dequeue() (T, error) {
	logritelog.Default().Printf(">>> Entering Dequeue()")
	if cq.IsEmpty() {
		var zeroValue T
		return zeroValue, errors.New("queue is empty")
	}
	retVal := cq.items[cq.front]
	if cq.front == cq.rear {
		cq.front = -1
		logritelog.Default().Printf("Dequeue(): cq.front → %v", cq.front)
		logritelog.Default().Printf("Dequeue(): cq.rear → %v", cq.rear)
		cq.rear = -1
	} else {
		cq.front = (cq.front + 1) % cq.size
	}
	logritelog.Default().Printf("<<< Exiting Dequeue()")
	return retVal, nil
}

// IsFull checks if the queue is full.
func (cq *CircularQueue[T]) IsFull() bool {
	logritelog.Default().Printf(">>> Entering IsFull()")
	return (cq.front == 0 && cq.rear == cq.size-1) || cq.front == cq.rear+1
}

// IsEmpty checks if the queue is empty.
func (cq *CircularQueue[T]) IsEmpty() bool {
	logritelog.Default().Printf(">>> Entering IsEmpty()")
	return cq.front == -1 && cq.rear == -1
}

// Peek returns the item at the front of the queue without removing it.
// Returns an error if the queue is empty.
func (cq *CircularQueue[T]) Peek() (T, error) {
	logritelog.Default().Printf(">>> Entering Peek()")
	if cq.IsEmpty() {
		var zeroValue T
	logritelog.Default().Printf("<<< Exiting Peek()")
		return zeroValue, errors.New("queue is empty")
	}

	return cq.items[cq.front], nil
}

// Size returns the size of the queue.
func (cq *CircularQueue[T]) Size() int {
	logritelog.Default().Printf(">>> Entering Size()")
	return cq.size
}
