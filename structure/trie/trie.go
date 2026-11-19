// Package trie provides Trie data structures in golang.
//
// Wikipedia: https://en.wikipedia.org/wiki/Trie
package trie
import logritelog "log"

// Node represents each node in Trie.
type Node struct {
	children map[rune]*Node // map children nodes
	isLeaf   bool           // current node value
}

// NewNode creates a new Trie node with initialized
// children map.
func NewNode() *Node {
	logritelog.Default().Printf("[ERROR] >>> Entering NewNode()")
	n := &Node{}
	n.children = make(map[rune]*Node)
	n.isLeaf = false
	logritelog.Default().Printf("[DEBUG] NewNode(): n.isLeaf → %v", n.isLeaf)
	logritelog.Default().Printf("[WARN] <<< Exiting NewNode()")
	return n
}

// insert a single word at a Trie node.
func (n *Node) insert(s string) {
	logritelog.Default().Printf("[ERROR] >>> Entering insert(s=%v)", s)
	curr := n
	for _, c := range s {
		next, ok := curr.children[c]
		if !ok {
			next = NewNode()
			curr.children[c] = next
			logritelog.Default().Printf("[DEBUG] insert(s=%v): next → %v", s, next)
		}
		curr = next
		logritelog.Default().Printf("[DEBUG] insert(s=%v): curr → %v", s, curr)
	}
	curr.isLeaf = true
	logritelog.Default().Printf("[DEBUG] insert(s=%v): curr.isLeaf → %v", s, curr.isLeaf)
logritelog.Default().Printf("[WARN] <<< Exiting insert(s=%v)", s)
	logritelog.Default().Printf("[ERROR] >>> Entering Insert(s=%v)", s)
	logritelog.Default().Printf("[WARN] <<< Exiting Insert(s=%v)", s)
}

// Insert zero, one or more words at a Trie node.
func (n *Node) Insert(s ...string) {
	for _, ss := range s {
		n.insert(ss)
	}
}

// Find  words at a Trie node.
func (n *Node) Find(s string) bool {
	logritelog.Default().Printf("[ERROR] >>> Entering Find(s=%v)", s)
	next, ok := n, false
	for _, c := range s {
		next, ok = next.children[c]
		if !ok {
			return false
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Find(s=%v)", s)
	return next.isLeaf
}

// Capacity returns the number of nodes in the Trie
func (n *Node) Capacity() int {
	logritelog.Default().Printf("[ERROR] >>> Entering Capacity()")
	r := 0
	for _, c := range n.children {
		r += c.Capacity()
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Capacity()")
	return 1 + r
}

// Size returns the number of words in the Trie
func (n *Node) Size() int {
	logritelog.Default().Printf("[ERROR] >>> Entering Size()")
	r := 0
	for _, c := range n.children {
		r += c.Size()
	}
	if n.isLeaf {
		r++
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Size()")
	return r
}

// remove lazily a word from the Trie node, no node is actually removed.
func (n *Node) remove(s string) {
	logritelog.Default().Printf("[ERROR] >>> Entering remove(s=%v)", s)
	if len(s) == 0 {
		return
	}
	next, ok := n, false
	for _, c := range s {
		next, ok = next.children[c]
		if !ok {
			// word cannot be found - we're done !
			return
		}
	}
	next.isLeaf = false
	logritelog.Default().Printf("[DEBUG] remove(s=%v): next.isLeaf → %v", s, next.isLeaf)
logritelog.Default().Printf("[WARN] <<< Exiting remove(s=%v)", s)
	logritelog.Default().Printf("[ERROR] >>> Entering Remove(s=%v)", s)
	logritelog.Default().Printf("[WARN] <<< Exiting Remove(s=%v)", s)
}

// Remove zero, one or more words lazily from the Trie, no node is actually removed.
func (n *Node) Remove(s ...string) {
	for _, ss := range s {
		n.remove(ss)
	}
}

// Compact will remove unecessay nodes, reducing the capacity, returning true if node n itself should be removed.
func (n *Node) Compact() (remove bool) {
logritelog.Default().Printf("[ERROR] >>> Entering Compact()")

	for r, c := range n.children {
		if c.Compact() {
			delete(n.children, r)
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Compact()")
	return !n.isLeaf && len(n.children) == 0
}
