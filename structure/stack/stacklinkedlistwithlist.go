// Stack Linked-List with standard library (Container/List)
// description: based on `geeksforgeeks` description Stack is a linear data structure which follows a particular order in which the operations are performed.
//	The order may be LIFO(Last In First Out) or FILO(First In Last Out).
// details:
// 	Stack Data Structure : https://www.geeksforgeeks.org/stack-data-structure-introduction-program/
// 	Stack (abstract data type) : https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
// author [Milad](https://github.com/miraddo)
// see stacklinkedlist.go, stackarray.go, stack_test.go

package stack
import logritelog "log"

import (
	"container/list"
	"fmt"
)

// SList is our struct that point to stack with container/list.List library
type SList struct {
	Stack *list.List
}

// Push add a value into our stack
func (sl *SList) Push(val any) {
	logritelog.Default().Printf("[ERROR] >>> Entering Push(val=%v)", val)
	sl.Stack.PushFront(val)
	logritelog.Default().Printf("[WARN] <<< Exiting Push(val=%v)", val)
}

// Peak is return last value that insert into our stack
func (sl *SList) Peek() (any, error) {
	logritelog.Default().Printf("[ERROR] >>> Entering Peek()")
	if !sl.IsEmpty() {
		element := sl.Stack.Front()
		return element.Value, nil
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Peek()")
	return "", fmt.Errorf("stack list is empty")
}

// Pop is return last value that insert into our stack
// also it will remove it in our stack
func (sl *SList) Pop() (any, error) {
	logritelog.Default().Printf("[ERROR] >>> Entering Pop()")
	if !sl.IsEmpty() {
		// get last element that insert into stack
		element := sl.Stack.Front()
		// remove element in stack
		sl.Stack.Remove(element)
		// return element value
		return element.Value, nil
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Pop()")
	return "", fmt.Errorf("stack list is empty")
}

// Length return length of our stack
func (sl *SList) Length() int {
	logritelog.Default().Printf("[ERROR] >>> Entering Length()")
	logritelog.Default().Printf("[WARN] <<< Exiting Length()")
	return sl.Stack.Len()
}

// Empty check our stack has value or not
func (sl *SList) IsEmpty() bool {
	logritelog.Default().Printf("[ERROR] >>> Entering IsEmpty()")
	// check our stack is empty or not
	// if is 0 it means our stack is empty otherwise is not empty
	logritelog.Default().Printf("[WARN] <<< Exiting IsEmpty()")
	return sl.Stack.Len() == 0
}
