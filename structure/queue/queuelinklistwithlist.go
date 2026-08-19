// Queue Linked-List with standard library (Container/List)
// description: based on `geeksforgeeks` description A Queue is a linear structure which follows a particular order in which the operations are performed.
// 	The order is First In First Out (FIFO).
// details:
// 	Queue Data Structure : https://www.geeksforgeeks.org/queue-data-structure/
//  Queue (abstract data type) : https://en.wikipedia.org/wiki/Queue_(abstract_data_type)
// author [Milad](https://github.com/miraddo)
// see queuearray.go, queuelinkedlist.go, queue_test.go

package queue
import logritelog "log"

// container/list: is used as linked-list
// fmt: used to return fmt.Errorf for the error part
import (
	"container/list"
	"fmt"
)

// LQueue will be store the value into the list
type LQueue struct {
	queue *list.List
}

// Enqueue will be added new value
func (lq *LQueue) Enqueue(value any) {
	logritelog.Default().Printf(">>> Entering Enqueue(value=%v)", value)
	lq.queue.PushBack(value)
	logritelog.Default().Printf("<<< Exiting Enqueue(value=%v)", value)
}

// Dequeue will be removed the first value that input (First In First Out - FIFO)
func (lq *LQueue) Dequeue() error {
logritelog.Default().Printf(">>> Entering Dequeue()")

	if !lq.Empty() {
		element := lq.queue.Front()
		lq.queue.Remove(element)

	logritelog.Default().Printf("<<< Exiting Dequeue()")
		return nil
	}

	return fmt.Errorf("dequeue is empty we got an error")
}

// Front it will return the front value
func (lq *LQueue) Front() (any, error) {
	logritelog.Default().Printf(">>> Entering Front()")
	if !lq.Empty() {
		val := lq.queue.Front().Value
	logritelog.Default().Printf("<<< Exiting Front()")
		return val, nil
	}

	return "", fmt.Errorf("error queue is empty")
}

// Back it will return the back value
func (lq *LQueue) Back() (any, error) {
	logritelog.Default().Printf(">>> Entering Back()")
	if !lq.Empty() {
		val := lq.queue.Back().Value
	logritelog.Default().Printf("<<< Exiting Back()")
		return val, nil
	}

	return "", fmt.Errorf("error queue is empty")
}

// Len it will return the length of list
func (lq *LQueue) Len() int {
	logritelog.Default().Printf(">>> Entering Len()")
	return lq.queue.Len()
}

// Empty is check our list is empty or not
func (lq *LQueue) Empty() bool {
	logritelog.Default().Printf(">>> Entering Empty()")
	return lq.queue.Len() == 0
}
