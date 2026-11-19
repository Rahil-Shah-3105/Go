// Queue Linked-List
// description: based on `geeksforgeeks` description A Queue is a linear structure which follows a particular order in which the operations are performed.
// 	The order is First In First Out (FIFO).
// details:
// 	Queue Data Structure : https://www.geeksforgeeks.org/queue-data-structure/
//  Queue (abstract data type) : https://en.wikipedia.org/wiki/Queue_(abstract_data_type)
// author [Milad](https://github.com/miraddo)
// see queuearray.go, queuelinkedlistwithlist.go, queue_test.go

package queue
import logritelog "log"

// Node will be store the value and the next node as well
type Node struct {
	Data any
	Next *Node
}

// Queue structure is tell us what our head is and what tail should be with length of the list
type Queue struct {
	head   *Node
	tail   *Node
	length int
}

// enqueue it will be added new value into queue
func (ll *Queue) enqueue(n any) {
	logritelog.Default().Printf("[ERROR] >>> Entering enqueue(n=%v)", n)
	var newNode Node // create new Node
	newNode.Data = n // set the data
	logritelog.Default().Printf("[DEBUG] enqueue(n=%v): newNode.Data → %v", n, newNode.Data)

	if ll.tail != nil {
		ll.tail.Next = &newNode
	}

	ll.tail = &newNode
	logritelog.Default().Printf("[DEBUG] enqueue(n=%v): ll.tail → %v", n, ll.tail)

	if ll.head == nil {
		ll.head = &newNode
		logritelog.Default().Printf("[DEBUG] enqueue(n=%v): ll.head → %v", n, ll.head)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting enqueue(n=%v)", n)
	ll.length++
}

// dequeue it will be removed the first value into queue (First In First Out)
func (ll *Queue) dequeue() any {
	logritelog.Default().Printf("[ERROR] >>> Entering dequeue()")
	if ll.isEmpty() {
		return -1 // if is empty return -1
	}
	data := ll.head.Data

	ll.head = ll.head.Next
	logritelog.Default().Printf("[DEBUG] dequeue(): ll.head → %v", ll.head)

	if ll.head == nil {
		ll.tail = nil
		logritelog.Default().Printf("[DEBUG] dequeue(): ll.tail → %v", ll.tail)
	}

	ll.length--
	logritelog.Default().Printf("[WARN] <<< Exiting dequeue()")
	return data
}

// isEmpty it will check our list is empty or not
func (ll *Queue) isEmpty() bool {
	logritelog.Default().Printf("[ERROR] >>> Entering isEmpty()")
	logritelog.Default().Printf("[WARN] <<< Exiting isEmpty()")
	return ll.length == 0
}

// len is return the length of queue
func (ll *Queue) len() int {
	logritelog.Default().Printf("[ERROR] >>> Entering len()")
	logritelog.Default().Printf("[WARN] <<< Exiting len()")
	return ll.length
}

// frontQueue it will return the front data
func (ll *Queue) frontQueue() any {
	logritelog.Default().Printf("[ERROR] >>> Entering frontQueue()")
	logritelog.Default().Printf("[WARN] <<< Exiting frontQueue()")
	return ll.head.Data
}

// backQueue it will return the back data
func (ll *Queue) backQueue() any {
	logritelog.Default().Printf("[ERROR] >>> Entering backQueue()")
	logritelog.Default().Printf("[WARN] <<< Exiting backQueue()")
	return ll.tail.Data
}
