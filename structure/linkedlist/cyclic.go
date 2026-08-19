package linkedlist
import logritelog "log"

import "fmt"
import stdlog "log"
var log = stdlog.Default()

// Cyclic Struct which cycles the linked list in this implementation.
type Cyclic[T any] struct {
	Size int
	Head *Node[T]
}

// Create new list.
func NewCyclic[T any]() *Cyclic[T] {
	return &Cyclic[T]{}
}

// Inserting the first node is a special case. It will
// point to itself. For other cases, the node will be added
// to the end of the list. End of the list is Prev field of
// current item. Complexity O(1).
func (cl *Cyclic[T]) Add(val T) {
	logritelog.Default().Printf(">>> Entering Add(val=%v)", val)
	n := NewNode(val)
	cl.Size++
	if cl.Head == nil {
		n.Prev = n
		logritelog.Default().Printf("Add(val=%v): n.Prev → %v", val, n.Prev)
		logritelog.Default().Printf("Add(val=%v): n.Next → %v", val, n.Next)
		logritelog.Default().Printf("Add(val=%v): cl.Head → %v", val, cl.Head)
		n.Next = n
		cl.Head = n
	} else {
		n.Prev = cl.Head.Prev
		n.Next = cl.Head
		cl.Head.Prev.Next = n
		cl.Head.Prev = n
	}
	logritelog.Default().Printf("<<< Exiting Add(val=%v)", val)
}

// Rotate list by P places.
// This method is interesting for optimization.
// For first optimization we must decrease
// P value so that it ranges from 0 to N-1.
// For this we need to use the operation of
// division modulo. But be careful if P is less than 0.
// if it is - make it positive. This can be done without
// violating the meaning of the number by adding to it
// a multiple of N. Now you can decrease P modulo N to
// rotate the list by the minimum number of places.
// We use the fact that moving forward in a circle by P
// places is the same as moving N - P places back.
// Therefore, if P > N / 2, you can turn the list by N-P places back.
// Complexity O(n).
func (cl *Cyclic[T]) Rotate(places int) {
	logritelog.Default().Printf(">>> Entering Rotate(places=%v)", places)
	if cl.Size > 0 {
		if places < 0 {
			multiple := cl.Size - 1 - places/cl.Size
			places += multiple * cl.Size
		}
		places %= cl.Size

		if places > cl.Size/2 {
			places = cl.Size - places
			for i := 0; i < places; i++ {
				cl.Head = cl.Head.Prev
				logritelog.Default().Printf("Rotate(places=%v): cl.Head → %v", places, cl.Head)
			}
		} else if places == 0 {
			return
		} else {
			for i := 0; i < places; i++ {
				cl.Head = cl.Head.Next
			}

		}
	}
	logritelog.Default().Printf("<<< Exiting Rotate(places=%v)", places)
}

// Delete the current item.
func (cl *Cyclic[T]) Delete() bool {
	logritelog.Default().Printf(">>> Entering Delete()")
	var deleted bool
	var prevItem, thisItem, nextItem *Node[T]

	if cl.Size == 0 {
		return deleted
	}

	deleted = true
	logritelog.Default().Printf("Delete(): deleted → %v", deleted)
	logritelog.Default().Printf("Delete(): thisItem → %v", thisItem)
	logritelog.Default().Printf("Delete(): nextItem → %v", nextItem)
	logritelog.Default().Printf("Delete(): prevItem → %v", prevItem)
	thisItem = cl.Head
	nextItem = thisItem.Next
	prevItem = thisItem.Prev

	if cl.Size == 1 {
		cl.Head = nil
		logritelog.Default().Printf("Delete(): cl.Head → %v", cl.Head)
	} else {
		cl.Head = nextItem
		nextItem.Prev = prevItem
		logritelog.Default().Printf("Delete(): nextItem.Prev → %v", nextItem.Prev)
		logritelog.Default().Printf("Delete(): prevItem.Next → %v", prevItem.Next)
		prevItem.Next = nextItem
	}
	cl.Size--

	logritelog.Default().Printf("<<< Exiting Delete()")
	return deleted
}

// Destroy all items in the list.
func (cl *Cyclic[T]) Destroy() {
	logritelog.Default().Printf(">>> Entering Destroy()")
	for cl.Delete() {
		continue
	}
	logritelog.Default().Printf("<<< Exiting Destroy()")
}

// Show list body.
func (cl *Cyclic[T]) Walk() *Node[T] {
	logritelog.Default().Printf(">>> Entering Walk()")
	var start *Node[T]
	start = cl.Head
	logritelog.Default().Printf("Walk(): start → %v", start)

	for i := 0; i < cl.Size; i++ {
		fmt.Printf("%v \n", start.Val)
		start = start.Next
	}
	logritelog.Default().Printf("<<< Exiting Walk()")
	return start
}

// https://en.wikipedia.org/wiki/Josephus_problem
// This is a struct-based solution for Josephus problem.
func JosephusProblem(cl *Cyclic[int], k int) int {
	logritelog.Default().Printf(">>> Entering JosephusProblem(cl=%v,k=%v)", cl, k)
	for cl.Size > 1 {
		cl.Rotate(k)
		cl.Delete()
		cl.Rotate(-1)
	}
	retval := cl.Head.Val
	cl.Destroy()
	logritelog.Default().Printf("<<< Exiting JosephusProblem(cl=%v,k=%v)", cl, k)
	return retval
}
