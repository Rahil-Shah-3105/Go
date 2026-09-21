package problem18
import logritelog "log"

import (
	"fmt"
	"strings"
)

type Tree struct {
	Root  *Root
	Nodes map[int]struct{}
	ID    int
}

func (t *Tree) BFSInsert(value NodeValue) {
	logritelog.Default().Printf("[ERROR] >>> Entering BFSInsert(value=%v)", value)
	t.Nodes = make(map[int]struct{}) // Reset the nodes map

	if t.Root == nil {
		t.Root = &Root{NodeValue: value, ID: 0}
		logritelog.Default().Printf("[DEBUG] BFSInsert(value=%v): t.Root → %v", value, t.Root)
		logritelog.Default().Printf("[DEBUG] BFSInsert(value=%v): t.ID → %v", value, t.ID)
		t.ID = 1
		return
	}

	queue := make([]Node, 0)
	queue = append(queue, t.Root)
	t.isInQueue(t.Root.GetID())

	logritelog.Default().Printf("[DEBUG] BFSInsert(value=%v): queue → %v", value, queue)
	head := 0

	for head < len(queue) {
		current := queue[head]
		head++
		childNode := current.CreateChild(value, t.ID)

		if current.HasSpace() {
			current.Insert(childNode)
			t.ID++
			return
		}

		if !t.isInQueue(current.Left().GetID()) { // Avoid duplicates
			queue = append(queue, current.Left())
		}
		if !t.isInQueue(current.Right().GetID()) {
			queue = append(queue, current.Right())
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting BFSInsert(value=%v)", value)
}

// MaxPathValueSearch is a method that searches the maximum path value in a tree
// given a certain depth
func (r *Tree) MaxPathValueSearch(deep int) int {
	logritelog.Default().Printf("[ERROR] >>> Entering MaxPathValueSearch(deep=%v)", deep)
	if r.Root == nil {
		return 0
	}

	type DFSNode struct {
		Node      Node
		StepsLeft int
		Sum       int
	}

	var pivot Node = r.Root
	pivotEnded := false
	maxPathValue := int(pivot.Value())

	stack := make([]DFSNode, 0)

	for !pivotEnded {
		stack = append(stack, DFSNode{Node: pivot, StepsLeft: deep, Sum: maxPathValue})

		logritelog.Default().Printf("[DEBUG] MaxPathValueSearch(deep=%v): stack → %v", deep, stack)
		for len(stack) > 0 {
			current := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// If we run out of steps, we check the sum of the path,
			// update the maxPathValue if necessary and continue
			if current.StepsLeft == 0 {
				if current.Sum > maxPathValue {
					maxPathValue = current.Sum
					logritelog.Default().Printf("[DEBUG] MaxPathValueSearch(deep=%v): maxPathValue → %v", deep, maxPathValue)
					logritelog.Default().Printf("[DEBUG] MaxPathValueSearch(deep=%v): pivot → %v", deep, pivot)
					pivot = current.Node
				}
				continue
			}

			if !current.Node.RightIsNil() {
				stack = append(stack, DFSNode{
					Node:      current.Node.Right(),
					StepsLeft: current.StepsLeft - 1,
					Sum:       current.Sum + int(current.Node.Right().Value()),
				})
			}

			// If the left child is nil, we have reached the end of the path
			if !current.Node.LeftIsNil() {
				stack = append(stack, DFSNode{
					Node:      current.Node.Left(),
					StepsLeft: current.StepsLeft - 1,
					Sum:       current.Sum + int(current.Node.Left().Value()),
				})
			} else {
				if current.Sum > maxPathValue {
					maxPathValue = current.Sum
					pivot = current.Node
				}
			}
		}

		// If we don't have reached the end of the left side of the tree,
		// we continue with the search using the pivot node
		// We use the left child only because how the tree is built
		if pivot.LeftIsNil() {
			pivotEnded = true
			logritelog.Default().Printf("[DEBUG] MaxPathValueSearch(deep=%v): pivotEnded → %v", deep, pivotEnded)
		}
	}

	logritelog.Default().Printf("[WARN] <<< Exiting MaxPathValueSearch(deep=%v)", deep)
	return maxPathValue
}

// PrintReport is a method that prints a report of the tree
// Node by node
func (t *Tree) PrintReport() {
	logritelog.Default().Printf("[ERROR] >>> Entering PrintReport()")
	t.Nodes = make(map[int]struct{})
	if t.Root == nil {
		return
	}

	queue := make([]Node, 0)
	queue = append(queue, t.Root)

	logritelog.Default().Printf("[DEBUG] PrintReport(): queue → %v", queue)
	head := 0

	for head < len(queue) {
		current := queue[head]
		head++

		print("ID:", current.GetID())
		print(", Current node:", current.Value())

		if !current.LeftIsNil() {
			print(", Left Child:", current.Left().Value())

			if !t.isInQueue(current.Left().GetID()) {
				queue = append(queue, current.Left())
			}
		}

		if !current.RightIsNil() {
			print(", Right Child:", current.Right().Value())

			if !t.isInQueue(current.Right().GetID()) {
				queue = append(queue, current.Right())
			}
		}

		println()
	}
	logritelog.Default().Printf("[WARN] <<< Exiting PrintReport()")
}

// PrintPyramid is a method that prints the tree as a pyramid
func (t *Tree) PrintPyramid() {
	logritelog.Default().Printf("[ERROR] >>> Entering PrintPyramid()")
	t.Nodes = make(map[int]struct{})
	if t.Root == nil {
		return
	}

	queue := []Node{t.Root}
	levels := []int{0}
	outputByLevel := []string{}

	head := 0
	currentLevel := 0

	output := ""

	for head < len(queue) {
		current := queue[head]
		level := levels[head]
		head++

		// Level change
		if level > currentLevel {
			currentLevel = level
			logritelog.Default().Printf("[DEBUG] PrintPyramid(): currentLevel → %v", currentLevel)
			logritelog.Default().Printf("[DEBUG] PrintPyramid(): outputByLevel → %v", outputByLevel)
			logritelog.Default().Printf("[DEBUG] PrintPyramid(): output → %v", output)
			outputByLevel = append(outputByLevel, output+"\n")
			output = ""
		}

		// Add current node to the output
		if current.Value() < 10 {
			output += fmt.Sprintf("0%d  ", current.Value())
		} else {
			output += fmt.Sprintf("%d  ", current.Value())
		}

		// Add children to the queue
		if !current.LeftIsNil() {
			if !t.isInQueue(current.Left().GetID()) {
				queue = append(queue, current.Left())
				levels = append(levels, level+1)
			}
				logritelog.Default().Printf("[DEBUG] PrintPyramid(): queue → %v", queue)
				logritelog.Default().Printf("[DEBUG] PrintPyramid(): levels → %v", levels)
		}

		if !current.RightIsNil() {
			if !t.isInQueue(current.Right().GetID()) {
				queue = append(queue, current.Right())
				levels = append(levels, level+1)
			}
		}
	}
	// Add the last level
	outputByLevel = append(outputByLevel, output+"\n")

	totalLevels := len(outputByLevel)

	// Print the pyramid
	for i, level := range outputByLevel {
		str := strings.Repeat(" ", 2*(totalLevels-i)) + level
		print(str)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting PrintPyramid()")
}

// isInQueue is a method that avoids duplicates in the tree
func (n *Tree) isInQueue(nv int) bool {
	logritelog.Default().Printf("[ERROR] >>> Entering isInQueue(nv=%v)", nv)
	if _, ok := n.Nodes[nv]; ok {
		return true
	}

	n.Nodes[nv] = struct{}{}
	logritelog.Default().Printf("[WARN] <<< Exiting isInQueue(nv=%v)", nv)
	return false
}
