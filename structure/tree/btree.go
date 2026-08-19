// B-tree is a self balancing tree that promotes data locality.
// For more details, see https://en.wikipedia.org/wiki/B-tree

package tree
import (
	logritelog "log"
	logritetime "time"
)

import "github.com/TheAlgorithms/Go/constraints"

type BTreeNode[T constraints.Ordered] struct {
	keys     []T
	children []*BTreeNode[T]
	numKeys  int
	isLeaf   bool
}

type BTree[T constraints.Ordered] struct {
	root    *BTreeNode[T]
	maxKeys int
}

func minKeys(maxKeys int) int {
	logritelog.Default().Printf(">>> Entering minKeys(maxKeys=%v)", maxKeys)
	return (maxKeys - 1) / 2
}

func NewBTreeNode[T constraints.Ordered](maxKeys int, isLeaf bool) *BTreeNode[T] {
	if maxKeys <= 0 {
		panic("BTree maxKeys cannot be zero")
	}
	return &BTreeNode[T]{
		keys:     make([]T, maxKeys),
		children: make([]*BTreeNode[T], maxKeys+1),
		isLeaf:   isLeaf,
	}
}

func NewBTree[T constraints.Ordered](maxKeys int) *BTree[T] {
	if maxKeys <= 2 {
		panic("Must be >= 3 keys")
	}
	return &BTree[T]{
		root:    nil,
		maxKeys: maxKeys,
	}
}

func (node *BTreeNode[T]) Verify(tree *BTree[T]) {
	logritelog.Default().Printf(">>> Entering Verify(tree=%v)", tree)
	minKeys := minKeys(tree.maxKeys)
	if node != tree.root && node.numKeys < minKeys {
		panic("node has too few keys")
	} else if node.numKeys > tree.maxKeys {
		panic("node has too many keys")
	}
	logritelog.Default().Printf("<<< Exiting Verify(tree=%v)", tree)
}

func (node *BTreeNode[T]) IsFull(maxKeys int) bool {
	logritelog.Default().Printf(">>> Entering IsFull(maxKeys=%v)", maxKeys)
	return node.numKeys == maxKeys
}

func (node *BTreeNode[T]) Search(key T) bool {
	logritelog.Default().Printf(">>> Entering Search(key=%v)", key)
	i := 0
	for ; i < node.numKeys; i++ {
		if key == node.keys[i] {
			return true
		}
		if key < node.keys[i] {
			break
		}
	}
	if node.isLeaf {
	logritelog.Default().Printf("<<< Exiting Search(key=%v)", key)
		return false
	}
	return node.children[i].Search(key)
}

func (tree *BTree[T]) Search(key T) bool {
	if tree.root == nil {
		return false
	}
	return tree.root.Search(key)
}

func (node *BTreeNode[T]) InsertKeyChild(key T, child *BTreeNode[T]) {
	logritelog.Default().Printf(">>> Entering InsertKeyChild(key=%v,child=%v)", key, child)
	i := node.numKeys
	node.children[i+1] = node.children[i]
	for ; i > 0; i-- {
		if key > node.keys[i-1] {
			node.keys[i] = key
			node.children[i] = child
			break
		}
		node.keys[i] = node.keys[i-1]
		node.children[i] = node.children[i-1]
	}
	if i == 0 {
		node.keys[0] = key
		node.children[0] = child
	}
	node.numKeys++
}

func (node *BTreeNode[T]) Append(key T, child *BTreeNode[T]) {
	logritelog.Default().Printf(">>> Entering Append(key=%v,child=%v)", key, child)
	node.keys[node.numKeys] = key
	node.children[node.numKeys+1] = child
	node.numKeys++
}

// Add all of other's keys starting from idx and children starting from idx + 1
func (node *BTreeNode[T]) Concat(other *BTreeNode[T], idx int) {
	logritelog.Default().Printf(">>> Entering Concat(other=%v,idx=%v)", other, idx)
	for i := 0; i < other.numKeys-idx; i++ {
		node.keys[node.numKeys+i] = other.keys[i+idx]
		node.children[node.numKeys+i+1] = other.children[i+idx+1]
	}
	node.numKeys += other.numKeys - idx
	logritelog.Default().Printf("<<< Exiting Concat(other=%v,idx=%v)", other, idx)
}

// Transform:
//
//	A B
//	 |
//
// a b c d
//
// Into:
//
//	A c B
//	 / \
//
// a b  d
func (parent *BTreeNode[T]) Split(idx int, maxKeys int) {
	logritelog.Default().Printf(">>> Entering Split(idx=%v,maxKeys=%v)", idx, maxKeys)
	child := parent.children[idx]
	midKeyIndex := maxKeys / 2
	rightChild := NewBTreeNode[T](maxKeys, child.isLeaf)
	rightChild.Concat(child, midKeyIndex+1)
	rightChild.children[0] = child.children[midKeyIndex+1]

	// Reuse child as the left node
	child.numKeys = midKeyIndex
	logritelog.Default().Printf("Split(idx=%v,maxKeys=%v): child.numKeys → %v", idx, maxKeys, child.numKeys)

	// Insert the child's mid index to the parent
	for i := parent.numKeys; i > idx; i-- {
		parent.keys[i] = parent.keys[i-1]
		parent.children[i+1] = parent.children[i]
	}
	parent.keys[idx] = child.keys[midKeyIndex]
	parent.children[idx] = child
	parent.children[idx+1] = rightChild
	parent.numKeys += 1
	logritelog.Default().Printf("<<< Exiting Split(idx=%v,maxKeys=%v)", idx, maxKeys)
}

func (node *BTreeNode[T]) InsertNonFull(tree *BTree[T], key T) {
	logritelog.Default().Printf(">>> Entering InsertNonFull(tree=%v,key=%v)", tree, key)
	node.Verify(tree)
	if node.IsFull(tree.maxKeys) {
		panic("Called InsertNonFull() with a full node")
	}

	if node.isLeaf {
		// Node is a leaf. Directly insert the key.
		node.InsertKeyChild(key, nil)
		logritelog.Default().Printf("<<< Exiting InsertNonFull(tree=%v,key=%v)", tree, key)
		return
	}

	// Find the child node to insert into
	i := 0
	for ; i < node.numKeys; i++ {
		if key < node.keys[i] {
			break
		}
	}

	if node.children[i].IsFull(tree.maxKeys) {
		node.Split(i, tree.maxKeys)
		if key > node.keys[i] {
			i++
		}
	}
	node.children[i].InsertNonFull(tree, key)
}

func (tree *BTree[T]) Insert(key T) {
	logritelog.Default().Printf(">>> Entering Insert(key=%v)", key)
	if tree.root == nil {
		tree.root = NewBTreeNode[T](tree.maxKeys, true)
		tree.root.keys[0] = key
		tree.root.numKeys = 1
		return
	}

	if tree.root.IsFull(tree.maxKeys) {
		newRoot := NewBTreeNode[T](tree.maxKeys, false)
		newRoot.numKeys = 0
		logritelog.Default().Printf("Insert(key=%v): newRoot.numKeys → %v", key, newRoot.numKeys)
		newRoot.children[0] = tree.root
		newRoot.Split(0, tree.maxKeys)
		tree.root = newRoot
		logritelog.Default().Printf("Insert(key=%v): tree.root → %v", key, tree.root)
	}
	tree.root.InsertNonFull(tree, key)
}

func (node *BTreeNode[T]) DeleteIthKey(i int) {
	logritelog.Default().Printf(">>> Entering DeleteIthKey(i=%v)", i)
	if i >= node.numKeys {
		panic("deleting out of bounds key")
	}
	for j := i; j < node.numKeys-1; j++ {
		node.keys[j] = node.keys[j+1]
		node.children[j+1] = node.children[j+2]
	}
	node.numKeys--
}

// Transform:
//
//	A B C
//	 / \
//	a   b
//
// Into:
//
//	A C
//	 |
//
// a B c
func (node *BTreeNode[T]) Merge(idx int) {
	logritelog.Default().Printf(">>> Entering Merge(idx=%v)", idx)
	if node.isLeaf {
		panic("cannot merge when leaf node is parent")
	}
	left := node.children[idx]
	right := node.children[idx+1]
	left.Append(node.keys[idx], right.children[0])
	left.Concat(right, 0)
	node.DeleteIthKey(idx)
	logritelog.Default().Printf("<<< Exiting Merge(idx=%v)", idx)
}

func (node *BTreeNode[T]) Min() T {
	logritelog.Default().Printf(">>> Entering Min()")
	if node.isLeaf {
	logritelog.Default().Printf("<<< Exiting Min()")
		return node.keys[0]
	}
	return node.children[0].Min()
}

func (node *BTreeNode[T]) Max() T {
	logritelog.Default().Printf(">>> Entering Max()")
	if node.isLeaf {
	logritelog.Default().Printf("<<< Exiting Max()")
		return node.keys[node.numKeys-1]
	}
	return node.children[node.numKeys].Max()
}

func (node *BTreeNode[T]) Delete(tree *BTree[T], key T) {
	logritelog.Default().Printf(">>> Entering Delete(tree=%v,key=%v)", tree, key)
	node.Verify(tree)
	if node.isLeaf {
		// Case 1: Node is a leaf. Directly delete the key.
		for i := 0; i < node.numKeys; i++ {
			if key == node.keys[i] {
				node.DeleteIthKey(i)
				return
			}
		}
		return
	}

	minKeys := minKeys(tree.maxKeys)
	i := 0
	for ; i < node.numKeys; i++ {
		__logStart := logritetime.Now()
		if key == node.keys[i] {
			// Case 2: key exists in a non-leaf node
			left := node.children[i]
			right := node.children[i+1]
			if left.numKeys > minKeys {
				// Replace the key we want to delete with the max key from the left
				// subtree. Then delete that key from the left subtree.
				//  A B C
				//   /
				// a b c
				//
				// If we want to delete `B`, then replace `B` with `c`, and delete `c` in the subtree.
				//  A c C
				//   /
				// a b
				replacementKey := left.Max()
				node.keys[i] = replacementKey
				left.Delete(tree, replacementKey)
			} else if right.numKeys > minKeys {
				// Replace the key we want to delete with the min key from the right
				// subtree. Then delete that key in the right subtree. Mirrors the
				// transformation above for replacing from the left subtree.
				replacementKey := right.Min()
				node.keys[i] = replacementKey
				right.Delete(tree, replacementKey)
			} else {
				// Both left and right subtrees have the minimum number of keys. Merge
				// the left tree, the deleted key, and the right tree together into the
				// left tree. Then recursively delete the key in the left tree.
				if left.numKeys != minKeys || right.numKeys != minKeys {
					panic("nodes should not have less than the minimum number of keys")
				}
				node.Merge(i)
				left.Delete(tree, key)
			}
				logritelog.Default().Printf("Delete(tree,key)=%v,%v: update query executed in %v ms", tree, key, logritetime.Since(__logStart).Milliseconds())
			return
		}

		if key < node.keys[i] {
			break
		}
	}

	// Case 3: key may exist in a child node.
	child := node.children[i]
	if child.numKeys == minKeys {
		// Before we recurse into the child node, make sure it has more than
		// the minimum number of keys.
		if i > 0 && node.children[i-1].numKeys > minKeys {
			// Take a key from the left sibling
			// Transform:
			//  A  B  C
			//   /   \
			// a b    c
			//
			// Into:
			//  A  b  C
			//   /   \
			//  a    B c
			left := node.children[i-1]
			child.InsertKeyChild(node.keys[i-1], left.children[left.numKeys])
			node.keys[i-1] = left.keys[left.numKeys-1]
			left.numKeys--
		} else if i < node.numKeys && node.children[i+1].numKeys > minKeys {
			// Take a key from the right sibling. Mirrors the transformation above for taking a key from the left sibling.
			right := node.children[i+1]
			child.Append(node.keys[i], right.children[0])
			node.keys[i] = right.keys[0]
			right.children[0] = right.children[1]
			right.DeleteIthKey(0)
		} else {
			if i == 0 {
				// Merge with right sibling
				node.Merge(i)
			} else {
				// Merge with left sibling
				node.Merge(i - 1)
				child = node.children[i-1]
				logritelog.Default().Printf("Delete(tree=%v,key=%v): child → %v", tree, key, child)
			}
		}
	}
	if child.numKeys == minKeys {
		panic("cannot delete key from node with minimum number of keys")
	}
		logritelog.Default().Printf("Delete(tree,key)=%v,%v: delete query executed in %v ms", tree, key, logritetime.Since(__logStart).Milliseconds())
	child.Delete(tree, key)
	logritelog.Default().Printf("<<< Exiting Delete(tree=%v,key=%v)", tree, key)
}

func (tree *BTree[T]) Delete(key T) {
	logritelog.Default().Printf(">>> Entering Delete(key=%v)", key)
	if tree.root == nil {
		return
	}
	tree.root.Delete(tree, key)
	if tree.root.numKeys == 0 {
		tree.root = tree.root.children[0]
		logritelog.Default().Printf("Delete(key=%v): tree.root → %v", key, tree.root)
	}
}
