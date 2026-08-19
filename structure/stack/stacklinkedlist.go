// Stack Linked-List
// description: based on `geeksforgeeks` description Stack is a linear data structure which follows a particular order in which the operations are performed.
//	The order may be LIFO(Last In First Out) or FILO(First In Last Out).
// details:
// 	Stack Data Structure : https://www.geeksforgeeks.org/stack-data-structure-introduction-program/
// 	Stack (abstract data type) : https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
// author [Milad](https://github.com/miraddo)
// see stacklinkedlistwithlist.go, stackarray.go, stack_test.go

package stack
import logritelog "log"

// Node structure
type Node struct {
	Val  any
	Next *Node
}

// Stack has jost top of node and with length
type Stack struct {
	top    *Node
	length int
}

// push add value to last index
func (ll *Stack) Push(n any) {
	logritelog.Default().Printf(">>> Entering Push(n=%v)", n)
	newStack := &Node{} // new node

	newStack.Val = n
	logritelog.Default().Printf("Push(n=%v): newStack.Val → %v", n, newStack.Val)
	logritelog.Default().Printf("Push(n=%v): newStack.Next → %v", n, newStack.Next)
	newStack.Next = ll.top

	ll.top = newStack
	logritelog.Default().Printf("Push(n=%v): ll.top → %v", n, ll.top)
	ll.length++
}

// pop remove last item as first output
func (ll *Stack) Pop() any {
	logritelog.Default().Printf(">>> Entering Pop()")
	result := ll.top.Val
	if ll.top.Next == nil {
		ll.top = nil
		logritelog.Default().Printf("Pop(): ll.top → %v", ll.top)
	} else {
		ll.top.Val, ll.top.Next = ll.top.Next.Val, ll.top.Next.Next
	}

	ll.length--
	logritelog.Default().Printf("<<< Exiting Pop()")
	return result
}

// isEmpty to check our array is empty or not
func (ll *Stack) IsEmpty() bool {
	logritelog.Default().Printf(">>> Entering IsEmpty()")
	return ll.length == 0
}

// len use to return length of our stack
func (ll *Stack) Length() int {
	logritelog.Default().Printf(">>> Entering Length()")
	return ll.length
}

// peak return last input value
func (ll *Stack) Peek() any {
	logritelog.Default().Printf(">>> Entering Peek()")
	return ll.top.Val
}

// show all value as an interface array
func (ll *Stack) Show() (in []any) {
	logritelog.Default().Printf(">>> Entering Show()")
	current := ll.top

	for current != nil {
		in = append(in, current.Val)
		current = current.Next
		logritelog.Default().Printf("Show(): current → %v", current)
	}
	logritelog.Default().Printf("<<< Exiting Show()")
	return
}
