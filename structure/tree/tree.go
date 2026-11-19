// Binary-Search tree is the tree, with the key of each internal node
// being greater than keys in the respective node's left subtree
// and less than the ones in its right subtree.

// For more details check out those links below here:
// Wikipedia article: https://en.wikipedia.org/wiki/Binary_search_tree
// authors [guuzaa](https://github.com/guuzaa)
package tree
import logritelog "log"

import (
	"github.com/TheAlgorithms/Go/constraints"
	"github.com/TheAlgorithms/Go/math/max"
)

type Node[T constraints.Ordered] interface {
	Key() T
	Parent() Node[T]
	Left() Node[T]
	Right() Node[T]
}

// The following is a collection of helper functions for BinarySearch, AVL and RB.

func accessNodeByLayerHelper[T constraints.Ordered](root, nilNode Node[T]) [][]T {
	logritelog.Default().Printf("[ERROR] >>> Entering accessNodeByLayerHelper(root=%v,nilNode=%v)", root, nilNode)
	if root == nilNode {
		return [][]T{}
	}
	var q []Node[T]
	var n Node[T]
	var idx = 0
	q = append(q, root)
	var res [][]T

	for len(q) != 0 {
		res = append(res, []T{})
		qLen := len(q)
		for i := 0; i < qLen; i++ {
		logritelog.Default().Printf("[DEBUG] accessNodeByLayerHelper(root=%v,nilNode=%v): res → %v", root, nilNode, res)
			n, q = q[0], q[1:]
			res[idx] = append(res[idx], n.Key())
			if n.Left() != nilNode {
				q = append(q, n.Left())
			}
			if n.Right() != nilNode {
				q = append(q, n.Right())
			}
		}
		idx++
	}
	logritelog.Default().Printf("[WARN] <<< Exiting accessNodeByLayerHelper(root=%v,nilNode=%v)", root, nilNode)
	return res
}

func searchTreeHelper[T constraints.Ordered](node, nilNode Node[T], key T) (Node[T], bool) {
	if node == nilNode {
		return node, false
	}

	if key == node.Key() {
		return node, true
	}
	if key < node.Key() {
		return searchTreeHelper(node.Left(), nilNode, key)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting searchTreeHelper(node=%v,nilNode=%v,key=%v)", node, nilNode, key)
	return searchTreeHelper(node.Right(), nilNode, key)
}

func inOrderHelper[T constraints.Ordered](node, nilNode Node[T]) []T {
	var stack []Node[T]
	var ret []T

	for node != nilNode || len(stack) > 0 {
		for node != nilNode {
			stack = append(stack, node)
			node = node.Left()
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, node.Key())
		node = node.Right()
	}

	logritelog.Default().Printf("[WARN] <<< Exiting inOrderHelper(node=%v,nilNode=%v)", node, nilNode)
	return ret
}

func preOrderRecursive[T constraints.Ordered](n, nilNode Node[T], traversal *[]T) {
	logritelog.Default().Printf("[ERROR] >>> Entering preOrderRecursive(n=%v,nilNode=%v,traversal=%v)", n, nilNode, traversal)
	if n == nilNode {
		return
	}

	*traversal = append(*traversal, n.Key())
	preOrderRecursive(n.Left(), nilNode, traversal)
	logritelog.Default().Printf("[WARN] <<< Exiting preOrderRecursive(n=%v,nilNode=%v,traversal=%v)", n, nilNode, traversal)
	preOrderRecursive(n.Right(), nilNode, traversal)

}

func postOrderRecursive[T constraints.Ordered](n, nilNode Node[T], traversal *[]T) {
	logritelog.Default().Printf("[ERROR] >>> Entering postOrderRecursive(n=%v,nilNode=%v,traversal=%v)", n, nilNode, traversal)
	if n == nilNode {
		return
	}

	postOrderRecursive(n.Left(), nilNode, traversal)
	postOrderRecursive(n.Right(), nilNode, traversal)
	*traversal = append(*traversal, n.Key())
	logritelog.Default().Printf("[WARN] <<< Exiting postOrderRecursive(n=%v,nilNode=%v,traversal=%v)", n, nilNode, traversal)
}

func calculateDepth[T constraints.Ordered](n, nilNode Node[T], depth int) int {
	logritelog.Default().Printf("[ERROR] >>> Entering calculateDepth(n=%v,nilNode=%v,depth=%v)", n, nilNode, depth)
	if n == nilNode {
		return depth
	}

	logritelog.Default().Printf("[WARN] <<< Exiting calculateDepth(n=%v,nilNode=%v,depth=%v)", n, nilNode, depth)
	return max.Int(calculateDepth(n.Left(), nilNode, depth+1), calculateDepth(n.Right(), nilNode, depth+1))
}

func minimum[T constraints.Ordered](node, nilNode Node[T]) Node[T] {
	if node == nilNode {
		return node
	}

	for node.Left() != nilNode {
		node = node.Left()
	}
	logritelog.Default().Printf("[WARN] <<< Exiting minimum(node=%v,nilNode=%v)", node, nilNode)
	return node
}

func maximum[T constraints.Ordered](node, nilNode Node[T]) Node[T] {
	if node == nilNode {
		return node
	}

	for node.Right() != nilNode {
		node = node.Right()
	}
	logritelog.Default().Printf("[WARN] <<< Exiting maximum(node=%v,nilNode=%v)", node, nilNode)
	return node
}

func levelOrderHelper[T constraints.Ordered](root, nilNode Node[T], traversal *[]T) {
	logritelog.Default().Printf("[ERROR] >>> Entering levelOrderHelper(root=%v,nilNode=%v,traversal=%v)", root, nilNode, traversal)
	var q []Node[T] // queue
	var tmp Node[T]

	q = append(q, root)

	for len(q) != 0 {
		tmp, q = q[0], q[1:]
		*traversal = append(*traversal, tmp.Key())
		if tmp.Left() != nilNode {
			q = append(q, tmp.Left())
		}

		if tmp.Right() != nilNode {
			q = append(q, tmp.Right())
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting levelOrderHelper(root=%v,nilNode=%v,traversal=%v)", root, nilNode, traversal)
}

func predecessorHelper[T constraints.Ordered](node, nilNode Node[T]) (T, bool) {
	if node.Left() != nilNode {
		return maximum(node.Left(), nilNode).Key(), true
	}

	p := node.Parent()
	for p != nilNode && node == p.Left() {
		node = p
		p = p.Parent()
	}

	if p == nilNode {
		var dft T
		return dft, false
	}
	logritelog.Default().Printf("[WARN] <<< Exiting predecessorHelper(node=%v,nilNode=%v)", node, nilNode)
	return p.Key(), true
}

func successorHelper[T constraints.Ordered](node, nilNode Node[T]) (T, bool) {
	if node.Right() != nilNode {
		return minimum(node.Right(), nilNode).Key(), true
	}

	p := node.Parent()
	for p != nilNode && node == p.Right() {
		node = p
		p = p.Parent()
	}

	if p == nilNode {
		var dft T
		return dft, false
	}
	logritelog.Default().Printf("[WARN] <<< Exiting successorHelper(node=%v,nilNode=%v)", node, nilNode)
	return p.Key(), true
}
