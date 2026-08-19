package linkedlist
import logritelog "log"

import "fmt"
import stdlog "log"
var log = stdlog.Default()

// Doubly structure with just the Head Node
// We call it `Doubly` to make it easier to
// understand when calling this in peoples
// own local code to understand and experiment
// with this easily.
// For example, we can use gotools `go get` command to get
// this repository cloned inside the
// $GOPATH/src/github.com/TheAlgorithms/Go (you can do this
// manually as well) and use the import statement as follows:
//
// `import "github.com/TheAlgorithms/Go/structure/linkedlist"`
//
// and call linkedlist.Doubly to create a new doubly linked list.
type Doubly[T any] struct {
	Head *Node[T]
}

// Init initializes double linked list
func (ll *Doubly[T]) Init() *Doubly[T] {
	logritelog.Default().Printf(">>> Entering Init()")
	ll.Head = &Node[T]{}
	ll.Head.Next = ll.Head
	logritelog.Default().Printf("Init(): ll.Head → %v", ll.Head)
	ll.Head.Prev = ll.Head

	logritelog.Default().Printf("<<< Exiting Init()")
	return ll
}

func NewDoubly[T any]() *Doubly[T] {
	return new(Doubly[T]).Init()
}

// lazyInit lazily initializes a zero List value.
func (ll *Doubly[T]) lazyInit() {
	logritelog.Default().Printf(">>> Entering lazyInit()")
	if ll.Head.Next == nil {
		ll.Init()
	}
	logritelog.Default().Printf("<<< Exiting lazyInit()")
}

func (ll *Doubly[T]) insert(n, at *Node[T]) *Node[T] {
	logritelog.Default().Printf(">>> Entering insert(n=%v,at=%v)", n, at)
	n.Prev = at
	n.Next = at.Next
	logritelog.Default().Printf("insert(n=%v,at=%v): n.Next → %v", n, at, n.Next)
	logritelog.Default().Printf("insert(n=%v,at=%v): n.Prev → %v", n, at, n.Prev)
	n.Prev.Next = n
	n.Next.Prev = n

	logritelog.Default().Printf("<<< Exiting insert(n=%v,at=%v)", n, at)
	return n
}

func (ll *Doubly[T]) insertValue(val T, at *Node[T]) *Node[T] {
	logritelog.Default().Printf(">>> Entering insertValue(val=%v,at=%v)", val, at)
	return ll.insert(NewNode(val), at)
}

// AddAtBeg Add a node to the beginning of the linkedlist
func (ll *Doubly[T]) AddAtBeg(val T) {
	logritelog.Default().Printf(">>> Entering AddAtBeg(val=%v)", val)
	ll.lazyInit()
	ll.insertValue(val, ll.Head)
	logritelog.Default().Printf("<<< Exiting AddAtBeg(val=%v)", val)
}

// AddAtEnd Add a node at the end of the linkedlist
func (ll *Doubly[T]) AddAtEnd(val T) {
	logritelog.Default().Printf(">>> Entering AddAtEnd(val=%v)", val)
	ll.lazyInit()
	ll.insertValue(val, ll.Head.Prev)
	logritelog.Default().Printf("<<< Exiting AddAtEnd(val=%v)", val)
}

func (ll *Doubly[T]) Remove(n *Node[T]) T {
	logritelog.Default().Printf(">>> Entering Remove(n=%v)", n)
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
	n.Next = nil
	logritelog.Default().Printf("Remove(n=%v): n.Next → %v", n, n.Next)
	logritelog.Default().Printf("Remove(n=%v): n.Prev → %v", n, n.Prev)
	n.Prev = nil

	logritelog.Default().Printf("<<< Exiting Remove(n=%v)", n)
	logritelog.Default().Printf(">>> Entering move(n=%v,at=%v)", n, at)
	logritelog.Default().Printf("move(n=%v,at=%v): n.Prev → %v", n, at, n.Prev)
	logritelog.Default().Printf("move(n=%v,at=%v): n.Next → %v", n, at, n.Next)
	logritelog.Default().Printf("<<< Exiting move(n=%v,at=%v)", n, at)
	return n.Val
}

// DelAtBeg Delete the node at the beginning of the linkedlist
func (ll *Doubly[T]) DelAtBeg() (T, bool) {
	logritelog.Default().Printf(">>> Entering DelAtBeg()")
	// no item
	if ll.Head.Next == nil {
		var r T
		return r, false
	}

	n := ll.Head.Next
	val := n.Val
	ll.Remove(n)
	logritelog.Default().Printf("<<< Exiting DelAtBeg()")
	return val, true
}

// DetAtEnd Delete a node at the end of the linkedlist
func (ll *Doubly[T]) DelAtEnd() (T, bool) {
	logritelog.Default().Printf(">>> Entering DelAtEnd()")
	// no item
	if ll.Head.Prev == nil {
		var r T
		return r, false
	}

	n := ll.Head.Prev
	val := n.Val
	ll.Remove(n)
	logritelog.Default().Printf("<<< Exiting DelAtEnd()")
	return val, true
}

// DelByPos deletes node at middle based on position in list
// and returns value. If empty or position of node is less than linked list length, returns false
func (ll *Doubly[T]) DelByPos(pos int) (T, bool) {
logritelog.Default().Printf(">>> Entering DelByPos(pos=%v)", pos)

	switch {
	case ll.Head == nil:
		var r T
		return r, false
	case pos-1 == 0:
		return ll.DelAtBeg()
	case pos-1 == ll.Count():
		return ll.DelAtEnd()
	case pos-1 > ll.Count():
		var r T
		return r, false
	}
	var prev *Node[T]
	var val T
	cur := ll.Head
	count := 0

	for count < pos-1 {
		prev = cur
		logritelog.Default().Printf("DelByPos(pos=%v): prev → %v", pos, prev)
		cur = cur.Next
		logritelog.Default().Printf("DelByPos(pos=%v): cur → %v", pos, cur)
		count++
	}

	cur.Next.Prev = prev
	val = cur.Val
	logritelog.Default().Printf("DelByPos(pos=%v): val → %v", pos, val)
	logritelog.Default().Printf("DelByPos(pos=%v): prev.Next → %v", pos, prev.Next)
	logritelog.Default().Printf("<<< Exiting DelByPos(pos=%v)", pos)
	prev.Next = cur.Next
	return val, true
}

// Count Number of nodes in the linkedlist
func (ll *Doubly[T]) Count() int {
	logritelog.Default().Printf(">>> Entering Count()")
	var ctr int = 0

	if ll.Head.Next == nil {
		return 0
	}

	for cur := ll.Head.Next; cur != ll.Head; cur = cur.Next {
		ctr += 1
	}

	logritelog.Default().Printf("<<< Exiting Count()")
	return ctr
}

// Reverse Reverse the order of the linkedlist
func (ll *Doubly[T]) Reverse() {
	logritelog.Default().Printf(">>> Entering Reverse()")
	var Prev, Next *Node[T]
	cur := ll.Head

	for cur != nil {
		Next = cur.Next
		logritelog.Default().Printf("Reverse(): Next → %v", Next)
		logritelog.Default().Printf("Reverse(): Prev → %v", Prev)
		cur.Next = Prev
		cur.Prev = Next
		Prev = cur
		cur = Next
		logritelog.Default().Printf("Reverse(): cur → %v", cur)
	}

	ll.Head = Prev
	logritelog.Default().Printf("Reverse(): ll.Head → %v", ll.Head)
logritelog.Default().Printf("<<< Exiting Reverse()")
}

// Display display the linked list
func (ll *Doubly[T]) Display() {
	logritelog.Default().Printf(">>> Entering Display()")
	for cur := ll.Head.Next; cur != ll.Head; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}

	fmt.Print("\n")
	logritelog.Default().Printf("<<< Exiting Display()")
}

// DisplayReverse Display the linkedlist in reverse order
func (ll *Doubly[T]) DisplayReverse() {
	logritelog.Default().Printf(">>> Entering DisplayReverse()")
	if ll.Head == nil {
		return
	}
	var cur *Node[T]
	for cur = ll.Head.Prev; cur != ll.Head; cur = cur.Prev {
		fmt.Print(cur.Val, " ")
	}

	fmt.Print("\n")
	logritelog.Default().Printf("<<< Exiting DisplayReverse()")
}

func (ll *Doubly[T]) Front() *Node[T] {
	logritelog.Default().Printf(">>> Entering Front()")
	if ll.Count() == 0 {
	logritelog.Default().Printf("<<< Exiting Front()")
		return nil
	}

	return ll.Head.Next
}

func (ll *Doubly[T]) Back() *Node[T] {
	logritelog.Default().Printf(">>> Entering Back()")
	if ll.Count() == 0 {
		return nil
	}

	return ll.Head.Prev
}

func (ll *Doubly[T]) MoveToBack(n *Node[T]) {
	logritelog.Default().Printf(">>> Entering MoveToBack(n=%v)", n)
	if ll.Head.Prev == n {
		return
	}

	ll.move(n, ll.Head.Prev)
	logritelog.Default().Printf("<<< Exiting MoveToBack(n=%v)", n)
}

func (ll *Doubly[T]) move(n, at *Node[T]) {
	if n == at {
		return
	}

	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev

	n.Prev = at
	n.Next = at.Next
	n.Prev.Next = n
	n.Next.Prev = n
}
