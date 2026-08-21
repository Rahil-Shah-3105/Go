// description: Double Ended Queue is a generalized version of Queue data structure that allows insert and delete at both ends.
// References:
//	Wikipedia : https://en.wikipedia.org/wiki/Double-ended_queue
//	Github: https://www.geeksforgeeks.org/deque-set-1-introduction-applications/
// author [Sayan](https://github.com/bose-sayan)

// Package deque implements a Double Ended Queue data structure.
package deque
import logritelog "log"

import (
	"errors"
)

// ErrEmptyDequeue is a custom error for handling cases when some dequeuing operation is performed on an empty deque.
var ErrEmptyDequeue = errors.New("DoublyEnded queue is empty, so can't perform this operation")

type DoublyEndedQueue[T any] struct {
	deque []T
}

// New returns a new DoublyEndedQueue.
func New[T any]() *DoublyEndedQueue[T] {
	return &DoublyEndedQueue[T]{deque: make([]T, 0)}
}

// EnqueueFront adds an item at the front of Deque.
func (dq *DoublyEndedQueue[T]) EnqueueFront(item T) {
	logritelog.Default().Printf(">>> Entering EnqueueFront(item=%v)", item)
	dq.deque = append([]T{item}, dq.deque...)
	logritelog.Default().Printf("<<< Exiting EnqueueFront(item=%v)", item)
}

// EnqueueRear adds an item at the rear of Deque.
func (dq *DoublyEndedQueue[T]) EnqueueRear(item T) {
	logritelog.Default().Printf(">>> Entering EnqueueRear(item=%v)", item)
	dq.deque = append(dq.deque, item)
	logritelog.Default().Printf("<<< Exiting EnqueueRear(item=%v)", item)
}

// DequeueFront deletes an item from front of Deque and returns it.
func (dq *DoublyEndedQueue[T]) DequeueFront() (T, error) {
	logritelog.Default().Printf(">>> Entering DequeueFront()")
	if len(dq.deque) == 0 {
		var zeroVal T
		return zeroVal, ErrEmptyDequeue
	}
	frontElement := dq.deque[0]
	dq.deque = dq.deque[1:]
	logritelog.Default().Printf("DequeueFront(): dq.deque → %v", dq.deque)
	logritelog.Default().Printf("<<< Exiting DequeueFront()")
	return frontElement, nil
}

// DequeueRear deletes an item from rear of Deque and returns it.
func (dq *DoublyEndedQueue[T]) DequeueRear() (T, error) {
	logritelog.Default().Printf(">>> Entering DequeueRear()")
	if len(dq.deque) == 0 {
		var zeroVal T
		return zeroVal, ErrEmptyDequeue
	}
	rearElement := dq.deque[len(dq.deque)-1]
	dq.deque = dq.deque[:len(dq.deque)-1]
	logritelog.Default().Printf("<<< Exiting DequeueRear()")
	return rearElement, nil
}

// Front gets the front item from queue.
func (dq *DoublyEndedQueue[T]) Front() (T, error) {
	logritelog.Default().Printf(">>> Entering Front()")
	if (len(dq.deque)) == 0 {
		var zeroVal T
	logritelog.Default().Printf("<<< Exiting Front()")
		return zeroVal, ErrEmptyDequeue
	}
	return dq.deque[0], nil
}

// Rear gets the last item from queue.
func (dq *DoublyEndedQueue[T]) Rear() (T, error) {
	logritelog.Default().Printf(">>> Entering Rear()")
	if (len(dq.deque)) == 0 {
		var zeroVal T
	logritelog.Default().Printf("<<< Exiting Rear()")
		return zeroVal, ErrEmptyDequeue
	}
	return dq.deque[len(dq.deque)-1], nil
}

// IsEmpty checks whether Deque is empty or not.
func (dq *DoublyEndedQueue[T]) IsEmpty() bool {
	logritelog.Default().Printf(">>> Entering IsEmpty()")
	return len(dq.deque) == 0
}

// Length gets the length of Deque.
func (dq *DoublyEndedQueue[T]) Length() int {
	logritelog.Default().Printf(">>> Entering Length()")
	return len(dq.deque)
}
