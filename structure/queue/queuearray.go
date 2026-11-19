// Queue Array
// description: based on `geeksforgeeks` description A Queue is a linear structure which follows a particular order in which the operations are performed.
// 	The order is First In First Out (FIFO).
// details:
// 	Queue Data Structure : https://www.geeksforgeeks.org/queue-data-structure/
//  Queue (abstract data type) : https://en.wikipedia.org/wiki/Queue_(abstract_data_type)
// author [Milad](https://github.com/miraddo)
// see queuelinkedlist.go, queuelinkedlistwithlist.go, queue_test.go

package queue
import logritelog "log"

var ListQueue []any

// EnQueue it will be added new value into our list
func EnQueue(n any) {
	logritelog.Default().Printf("[ERROR] >>> Entering EnQueue(n=%v)", n)
	ListQueue = append(ListQueue, n)
	logritelog.Default().Printf("[DEBUG] EnQueue(n=%v): ListQueue → %v", n, ListQueue)
logritelog.Default().Printf("[WARN] <<< Exiting EnQueue(n=%v)", n)
}

// DeQueue it will be removed the first value that added into the list
func DeQueue() any {
	logritelog.Default().Printf("[ERROR] >>> Entering DeQueue()")
	data := ListQueue[0]
	ListQueue = ListQueue[1:]
	logritelog.Default().Printf("[DEBUG] DeQueue(): ListQueue → %v", ListQueue)
	logritelog.Default().Printf("[WARN] <<< Exiting DeQueue()")
	return data
}

// FrontQueue return the Front value
func FrontQueue() any {
	logritelog.Default().Printf("[ERROR] >>> Entering FrontQueue()")
	return ListQueue[0]
}

// BackQueue return the Back value
func BackQueue() any {
	logritelog.Default().Printf("[ERROR] >>> Entering BackQueue()")
	return ListQueue[len(ListQueue)-1]
}

// LenQueue will return the length of the queue list
func LenQueue() int {
	logritelog.Default().Printf("[ERROR] >>> Entering LenQueue()")
	return len(ListQueue)
}

// IsEmptyQueue check our list is empty or not
func IsEmptyQueue() bool {
	logritelog.Default().Printf("[ERROR] >>> Entering IsEmptyQueue()")
	return len(ListQueue) == 0
}
