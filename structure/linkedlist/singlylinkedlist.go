package linkedlist
import logritelog "log"

// demonstration of singly linked list in golang
import (
	"errors"
	"fmt"
)

// Singly structure with length of the list and its head
type Singly[T any] struct {
	length int

	// Note that Node here holds both Next and Prev Node
	// however only the Next node is used in Singly methods.
	Head *Node[T]
}

// NewSingly returns a new instance of a linked list
func NewSingly[T any]() *Singly[T] {
	return &Singly[T]{}
}

// AddAtBeg adds a new snode with given value at the beginning of the list.
func (ll *Singly[T]) AddAtBeg(val T) {
	logritelog.Default().Printf(">>> Entering AddAtBeg(val=%v)", val)
	n := NewNode(val)
	n.Next = ll.Head
	logritelog.Default().Printf("AddAtBeg(val=%v): n.Next → %v", val, n.Next)
	logritelog.Default().Printf("AddAtBeg(val=%v): ll.Head → %v", val, ll.Head)
	ll.Head = n
	ll.length++
}

// AddAtEnd adds a new snode with given value at the end of the list.
func (ll *Singly[T]) AddAtEnd(val T) {
	logritelog.Default().Printf(">>> Entering AddAtEnd(val=%v)", val)
	n := NewNode(val)

	if ll.Head == nil {
		ll.Head = n
		logritelog.Default().Printf("AddAtEnd(val=%v): ll.Head → %v", val, ll.Head)
		ll.length++
		return
	}

	cur := ll.Head
	for ; cur.Next != nil; cur = cur.Next {
	}
	cur.Next = n
	logritelog.Default().Printf("AddAtEnd(val=%v): cur.Next → %v", val, cur.Next)
	ll.length++
}

// DelAtBeg deletes the snode at the head(beginning) of the list
// and returns its value. Returns false if the list is empty.
func (ll *Singly[T]) DelAtBeg() (T, bool) {
	logritelog.Default().Printf(">>> Entering DelAtBeg()")
	if ll.Head == nil {
		var r T
		return r, false
	}

	cur := ll.Head
	ll.Head = cur.Next
	logritelog.Default().Printf("DelAtBeg(): ll.Head → %v", ll.Head)
	ll.length--

	logritelog.Default().Printf("<<< Exiting DelAtBeg()")
	return cur.Val, true
}

// DelAtEnd deletes the snode at the tail(end) of the list
// and returns its value. Returns false if the list is empty.
func (ll *Singly[T]) DelAtEnd() (T, bool) {
	logritelog.Default().Printf(">>> Entering DelAtEnd()")
	if ll.Head == nil {
		var r T
		return r, false
	}

	if ll.Head.Next == nil {
		return ll.DelAtBeg()
	}

	cur := ll.Head

	for ; cur.Next.Next != nil; cur = cur.Next {
	}

	retval := cur.Next.Val
	cur.Next = nil
	logritelog.Default().Printf("DelAtEnd(): cur.Next → %v", cur.Next)
	ll.length--
	logritelog.Default().Printf("<<< Exiting DelAtEnd()")
	return retval, true

}

// DelByPos deletes the node at the middle based on position in the list
// and returns its value. Returns false if the list is empty or length is not more than given position
func (ll *Singly[T]) DelByPos(pos int) (T, bool) {
	logritelog.Default().Printf(">>> Entering DelByPos(pos=%v)", pos)
	switch {
	case ll.Head == nil:
		var r T
		return r, false
	case pos-1 > ll.length:
		var r T
		return r, false
	case pos-1 == 0:
		return ll.DelAtBeg()
	case pos-1 == ll.Count():
		return ll.DelAtEnd()
	}

	var prev *Node[T]
	var val T
	cur := ll.Head
	count := 0

	for count < pos-1 {
		prev = cur
		logritelog.Default().Printf("DelByPos(pos=%v): prev → %v", pos, prev)
		logritelog.Default().Printf("DelByPos(pos=%v): cur → %v", pos, cur)
		cur = cur.Next
		count++
	}

	val = cur.Val
	logritelog.Default().Printf("DelByPos(pos=%v): val → %v", pos, val)
	logritelog.Default().Printf("DelByPos(pos=%v): prev.Next → %v", pos, prev.Next)
	prev.Next = cur.Next
	ll.length--

	logritelog.Default().Printf("<<< Exiting DelByPos(pos=%v)", pos)
	return val, true
}

// Count returns the current size of the list.
func (ll *Singly[T]) Count() int {
	logritelog.Default().Printf(">>> Entering Count()")
	return ll.length
}

// Reverse reverses the list.
func (ll *Singly[T]) Reverse() {
	logritelog.Default().Printf(">>> Entering Reverse()")
	var prev, next *Node[T]
	cur := ll.Head

	for cur != nil {
		next = cur.Next
		logritelog.Default().Printf("Reverse(): next → %v", next)
		logritelog.Default().Printf("Reverse(): cur.Next → %v", cur.Next)
		logritelog.Default().Printf("Reverse(): prev → %v", prev)
		logritelog.Default().Printf("Reverse(): cur → %v", cur)
		cur.Next = prev
		prev = cur
		cur = next
	}

	ll.Head = prev
	logritelog.Default().Printf("Reverse(): ll.Head → %v", ll.Head)
logritelog.Default().Printf("<<< Exiting Reverse()")
}

// ReversePartition Reverse the linked list from the ath to the bth node
func (ll *Singly[T]) ReversePartition(left, right int) error {
	logritelog.Default().Printf(">>> Entering ReversePartition(left=%v,right=%v)", left, right)
	err := ll.CheckRangeFromIndex(left, right)
	if err != nil {
		logritelog.Default().Printf("[ERROR] Exception in ReversePartition(left=%v,right=%v): %v", left, right, err)
		logritelog.Default().Printf(">>> Entering CheckRangeFromIndex(left=%v,right=%v)", left, right)
		logritelog.Default().Printf("<<< Exiting CheckRangeFromIndex(left=%v,right=%v)", left, right)
		return err
	}
	tmpNode := &Node[T]{}
	tmpNode.Next = ll.Head
	logritelog.Default().Printf("ReversePartition(left=%v,right=%v): tmpNode.Next → %v", left, right, tmpNode.Next)
	pre := tmpNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
		logritelog.Default().Printf("ReversePartition(left=%v,right=%v): pre → %v", left, right, pre)
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		logritelog.Default().Printf("ReversePartition(left=%v,right=%v): cur.Next → %v", left, right, cur.Next)
		logritelog.Default().Printf("ReversePartition(left=%v,right=%v): next.Next → %v", left, right, next.Next)
		logritelog.Default().Printf("ReversePartition(left=%v,right=%v): pre.Next → %v", left, right, pre.Next)
		next.Next = pre.Next
		pre.Next = next
	}
	ll.Head = tmpNode.Next
	logritelog.Default().Printf("ReversePartition(left=%v,right=%v): ll.Head → %v", left, right, ll.Head)
	logritelog.Default().Printf("<<< Exiting ReversePartition(left=%v,right=%v)", left, right)
	return nil
}
func (ll *Singly[T]) CheckRangeFromIndex(left, right int) error {
	if left > right {
		return errors.New("left boundary must smaller than right")
	} else if left < 1 {
		return errors.New("left boundary starts from the first node")
	} else if right > ll.length {
		return errors.New("right boundary cannot be greater than the length of the linked list")
	}
	return nil
}

// Display prints out the elements of the list.
func (ll *Singly[T]) Display() {
	logritelog.Default().Printf(">>> Entering Display()")
	for cur := ll.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}

	fmt.Print("\n")
	logritelog.Default().Printf("<<< Exiting Display()")
}
