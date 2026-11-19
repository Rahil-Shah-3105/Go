package linkedlist
import logritelog "log"

// Node Structure representing the linkedlist node.
// This node is shared across different implementations.
type Node[T any] struct {
	Val  T
	Prev *Node[T]
	Next *Node[T]
}

// Create new node.
func NewNode[T any](val T) *Node[T] {
	logritelog.Default().Printf("[ERROR] >>> Entering NewNode(val=%v)", val)
	logritelog.Default().Printf("[WARN] <<< Exiting NewNode(val=%v)", val)
	return &Node[T]{val, nil, nil}
}
