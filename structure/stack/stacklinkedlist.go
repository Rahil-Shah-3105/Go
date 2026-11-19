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
	logritelog.Default().Printf("[ERROR] >>> Entering Push(n=%v)", n)
	newStack := &Node{} // new node

	newStack.Val = n
	logritelog.Default().Printf("[DEBUG] Push(n=%v): newStack.Val → %v", n, newStack.Val)
	newStack.Next = ll.top
	logritelog.Default().Printf("[DEBUG] Push(n=%v): newStack.Next → %v", n, newStack.Next)

	ll.top = newStack
	logritelog.Default().Printf("[DEBUG] Push(n=%v): ll.top → %v", n, ll.top)
	logritelog.Default().Printf("[WARN] <<< Exiting Push(n=%v)", n)
	ll.length++
}

// pop remove last item as first output
func (ll *Stack) Pop() any {
	logritelog.Default().Printf("[ERROR] >>> Entering Pop()")
	result := ll.top.Val
	if ll.top.Next == nil {
		ll.top = nil
	} else {
		ll.top.Val, ll.top.Next = ll.top.Next.Val, ll.top.Next.Next
	}

	ll.length--
	logritelog.Default().Printf("[WARN] <<< Exiting Pop()")
	return result
}

// isEmpty to check our array is empty or not
func (ll *Stack) IsEmpty() bool {
	logritelog.Default().Printf("[ERROR] >>> Entering IsEmpty()")
	logritelog.Default().Printf("[WARN] <<< Exiting IsEmpty()")
	return ll.length == 0
}

// len use to return length of our stack
func (ll *Stack) Length() int {
	logritelog.Default().Printf("[ERROR] >>> Entering Length()")
	logritelog.Default().Printf("[WARN] <<< Exiting Length()")
	return ll.length
}

// peak return last input value
func (ll *Stack) Peek() any {
	logritelog.Default().Printf("[ERROR] >>> Entering Peek()")
	logritelog.Default().Printf("[WARN] <<< Exiting Peek()")
	return ll.top.Val
}

// show all value as an interface array
func (ll *Stack) Show() (in []any) {
	logritelog.Default().Printf("[ERROR] >>> Entering Show()")
	current := ll.top

	for current != nil {
		in = append(in, current.Val)
		current = current.Next
		logritelog.Default().Printf("[DEBUG] Show(): current → %v", current)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Show()")
	return
}
