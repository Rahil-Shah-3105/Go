package tree_test
import logritelog "log"

import (
	"fmt"
	"github.com/TheAlgorithms/Go/constraints"
	"testing"

	bt "github.com/TheAlgorithms/Go/structure/tree"
)

type TestTree[T constraints.Ordered] interface {
	Push(...T)
	Delete(T) bool
	Get(T) (bt.Node[T], bool)
	Empty() bool
	Has(T) bool
	Depth() int
	Max() (T, bool)
	Min() (T, bool)
	Predecessor(T) (T, bool)
	Successor(T) (T, bool)
	PreOrder() []T
	InOrder() []T
	PostOrder() []T
	LevelOrder() []T
	AccessNodesByLayer() [][]T
}

// BinarySearch, AVL and RB have completed the `TestTree` interface.
var (
	_ TestTree[int] = (*bt.BinarySearch[int])(nil)
	_ TestTree[int] = (*bt.AVL[int])(nil)
	_ TestTree[int] = (*bt.RB[int])(nil)
)

var tree TestTree[int]

func TestBinarySearch(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestBinarySearch(t=%v)", t)
	tree = bt.NewBinarySearch[int]()

	logritelog.Default().Printf("[DEBUG] TestBinarySearch(t=%v): tree → %v", t, tree)
	if tree.Empty() {
		t.Log("Binary Search Tree is empty now.")
	}

	tree.Push(1, 4, 10)
	tree.Push(-8)
	nums := []int{87, 18, 10, -34}
	tree.Push(nums...)
	tree.Push(4) // duplicate key, dismiss it

	if tree.Has(4) {
		t.Logf("There is a node of 4")
	}

	if n, ok := tree.Get(10); ok {
		t.Logf("node of 10: %T", n)
	}

	if ret, ok := tree.Min(); ok {
		t.Logf("tree.Min() = %v", ret)
	}

	if ret, ok := tree.Max(); ok {
		t.Logf("tree.Max() = %v", ret)
	}

	if ret, ok := tree.Predecessor(1); ok {
		t.Logf("tree.Preducessor(1) = %v", ret)
	}

	if ret, ok := tree.Successor(18); ok {
		t.Logf("tree.Successor(18) = %v", ret)
	}

	fmt.Println(tree.InOrder())
	fmt.Println(tree.AccessNodesByLayer())

	tree.Delete(18)
	fmt.Println("Delete 18")
	fmt.Println(tree.InOrder())
	logritelog.Default().Printf("[WARN] <<< Exiting TestBinarySearch(t=%v)", t)
}

func TestAVL(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestAVL(t=%v)", t)
	tree = bt.NewAVL[int]()

	logritelog.Default().Printf("[DEBUG] TestAVL(t=%v): tree → %v", t, tree)
	if tree.Empty() {
		t.Log("AVL Tree is empty now.")
	}

	tree.Push(1, 4, 10)
	tree.Push(-8)
	nums := []int{87, 18, 10, -34}
	tree.Push(nums...)
	tree.Push(4) // duplicate key, dismiss it

	if tree.Has(4) {
		t.Logf("There is a node of 4")
	}

	if n, ok := tree.Get(10); ok {
		t.Logf("node of 10: %T", n)
	}

	if ret, ok := tree.Min(); ok {
		t.Logf("tree.Min() = %v", ret)
	}

	if ret, ok := tree.Max(); ok {
		t.Logf("tree.Max() = %v", ret)
	}

	if ret, ok := tree.Predecessor(1); ok {
		t.Logf("tree.Preducessor(1) = %v", ret)
	}

	if ret, ok := tree.Successor(18); ok {
		t.Logf("tree.Successor(18) = %v", ret)
	}

	fmt.Println(tree.InOrder())
	fmt.Println(tree.AccessNodesByLayer())

	tree.Delete(18)
	fmt.Println("Delete 18")
	fmt.Println(tree.InOrder())
	logritelog.Default().Printf("[WARN] <<< Exiting TestAVL(t=%v)", t)
}

func TestRB(t *testing.T) {
	logritelog.Default().Printf("[ERROR] >>> Entering TestRB(t=%v)", t)
	tree = bt.NewRB[int]()

	logritelog.Default().Printf("[DEBUG] TestRB(t=%v): tree → %v", t, tree)
	if tree.Empty() {
		t.Log("RB Tree is empty now.")
	}

	tree.Push(1, 4, 10)
	tree.Push(-8)
	nums := []int{87, 18, 10, -34}
	tree.Push(nums...)
	tree.Push(4) // duplicate key, dismiss it

	if tree.Has(4) {
		t.Logf("There is a node of 4")
	}

	if n, ok := tree.Get(10); ok {
		t.Logf("node of 10: %T", n)
	}

	if ret, ok := tree.Min(); ok {
		t.Logf("tree.Min() = %v", ret)
	}

	if ret, ok := tree.Max(); ok {
		t.Logf("tree.Max() = %v", ret)
	}

	if ret, ok := tree.Predecessor(1); ok {
		t.Logf("tree.Preducessor(1) = %v", ret)
	}

	if ret, ok := tree.Successor(18); ok {
		t.Logf("tree.Successor(18) = %v", ret)
	}

	fmt.Println(tree.InOrder())
	fmt.Println(tree.AccessNodesByLayer())
	logritelog.Default().Printf("[WARN] <<< Exiting TestRB(t=%v)", t)
}
