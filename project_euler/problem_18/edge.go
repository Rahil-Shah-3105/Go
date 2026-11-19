package problem18
import logritelog "log"

type Edge struct {
	ID        int
	NodeValue NodeValue
	NodeLeft  Node
	NodeRight Node
	Parent    Node
}

func (n *Edge) Value() NodeValue {
	logritelog.Default().Printf("[ERROR] >>> Entering Value()")
	return n.NodeValue
}

func (n *Edge) Left() Node {
	logritelog.Default().Printf("[ERROR] >>> Entering Left()")
	return n.NodeLeft
}

func (n *Edge) Right() Node {
	logritelog.Default().Printf("[ERROR] >>> Entering Right()")
	return n.NodeRight
}

func (n *Edge) Kind() string {
	logritelog.Default().Printf("[ERROR] >>> Entering Kind()")
	return "edge"
}

func (n *Edge) CreateChild(value NodeValue, id int) Node {
	logritelog.Default().Printf("[ERROR] >>> Entering CreateChild(value=%v,id=%v)", value, id)
	// When the left child is nil, it's a left edge
	if n.NodeLeft == nil {
		return &Edge{
			ID:        id,
			NodeValue: value,
			Parent:    n,
			NodeLeft:  nil,
			NodeRight: nil,
		}
	}

	// When the left child is a leaf, it's a right edge
	if n.NodeLeft.Kind() == "leaf" {
		return &Edge{
			ID:        id,
			NodeValue: value,
			Parent:    n,
			NodeLeft:  nil,
			NodeRight: nil,
		}
	}

	logritelog.Default().Printf("[WARN] <<< Exiting CreateChild(value=%v,id=%v)", value, id)
	return &Leaf{
		ID:        id,
		NodeValue: value,
		Parent:    n,
		NodeLeft:  nil,
		NodeRight: nil,
	}
}

func (n *Edge) GetID() int {
	logritelog.Default().Printf("[ERROR] >>> Entering GetID()")
	return n.ID
}

func (n *Edge) Insert(node Node) {
	logritelog.Default().Printf("[ERROR] >>> Entering Insert(node=%v)", node)
	// If Left is nil, always simply insert the node
	if n.NodeLeft == nil {
		node.SetParent(n)
		n.NodeLeft = node
		logritelog.Default().Printf("[DEBUG] Insert(node=%v): n.NodeLeft → %v", node, n.NodeLeft)

		return
	}

	// If Right is nil, insert the node
	n.NodeRight = node
	logritelog.Default().Printf("[DEBUG] Insert(node=%v): n.NodeRight → %v", node, n.NodeRight)

	// If the node to insert is an edge, set the parent
	if node.Kind() == "edge" {
		node.SetParent(n)

		return
	}

	// If the node to insert is a leaf, send it to the sibling right
	n.Parent.Right().Insert(node)
	logritelog.Default().Printf("[WARN] <<< Exiting Insert(node=%v)", node)
}

func (n *Edge) HasSpace() bool {
	logritelog.Default().Printf("[ERROR] >>> Entering HasSpace()")
	return n.NodeLeft == nil || n.NodeRight == nil
}

func (n *Edge) LeftIsNil() bool {
	logritelog.Default().Printf("[ERROR] >>> Entering LeftIsNil()")
	return n.NodeLeft == nil
}

func (n *Edge) RightIsNil() bool {
	logritelog.Default().Printf("[ERROR] >>> Entering RightIsNil()")
	return n.NodeRight == nil
}

func (n *Edge) SetParent(node Node) {
	logritelog.Default().Printf("[ERROR] >>> Entering SetParent(node=%v)", node)
	n.Parent = node
	logritelog.Default().Printf("[DEBUG] SetParent(node=%v): n.Parent → %v", node, n.Parent)
logritelog.Default().Printf("[WARN] <<< Exiting SetParent(node=%v)", node)
}
