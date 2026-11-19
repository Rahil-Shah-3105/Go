package problem18
import logritelog "log"

type Leaf struct {
	ID        int
	NodeValue NodeValue
	NodeLeft  *Leaf
	NodeRight *Leaf
	Parent    Node
}

func (n *Leaf) Value() NodeValue {
	logritelog.Default().Printf("[ERROR] >>> Entering Value()")
	return n.NodeValue
}

func (n *Leaf) Left() Node {
	logritelog.Default().Printf("[ERROR] >>> Entering Left()")
	if n.NodeLeft != nil {
		n.NodeLeft.Parent = n // Leaf is the parent of its left child always
	}

	logritelog.Default().Printf("[WARN] <<< Exiting Left()")
	return n.NodeLeft
}

func (n *Leaf) Right() Node {
	logritelog.Default().Printf("[ERROR] >>> Entering Right()")
	return n.NodeRight
}

func (n *Leaf) Kind() string {
	logritelog.Default().Printf("[ERROR] >>> Entering Kind()")
	return "leaf"
}

func (n *Leaf) CreateChild(value NodeValue, id int) Node {
	logritelog.Default().Printf("[ERROR] >>> Entering CreateChild(value=%v,id=%v)", value, id)
	// Leafs only have leaf children
	logritelog.Default().Printf("[WARN] <<< Exiting CreateChild(value=%v,id=%v)", value, id)
	return &Leaf{
		ID:        id,
		NodeValue: value,
		Parent:    n,
		NodeLeft:  nil,
		NodeRight: nil,
	}
}

func (n *Leaf) GetID() int {
	logritelog.Default().Printf("[ERROR] >>> Entering GetID()")
	return n.ID
}

func (n *Leaf) Insert(node Node) {
	logritelog.Default().Printf("[ERROR] >>> Entering Insert(node=%v)", node)
	// If Left is nil, always simply insert the node
	if n.NodeLeft == nil {
		node.SetParent(n)
		n.NodeLeft = node.(*Leaf)

		return
	}

	// If Right is nil, insert the node
	n.NodeRight = node.(*Leaf)
	// Send it to the sibling right
	n.Parent.Right().Insert(node)
	logritelog.Default().Printf("[WARN] <<< Exiting Insert(node=%v)", node)
}

func (n *Leaf) HasSpace() bool {
	logritelog.Default().Printf("[ERROR] >>> Entering HasSpace()")
	return n.NodeLeft == nil || n.NodeRight == nil
}

func (n *Leaf) LeftIsNil() bool {
	logritelog.Default().Printf("[ERROR] >>> Entering LeftIsNil()")
	return n.NodeLeft == nil
}

func (n *Leaf) RightIsNil() bool {
	logritelog.Default().Printf("[ERROR] >>> Entering RightIsNil()")
	return n.NodeRight == nil
}

func (n *Leaf) SetParent(node Node) {
	logritelog.Default().Printf("[ERROR] >>> Entering SetParent(node=%v)", node)
	n.Parent = node
	logritelog.Default().Printf("[DEBUG] SetParent(node=%v): n.Parent → %v", node, n.Parent)
logritelog.Default().Printf("[WARN] <<< Exiting SetParent(node=%v)", node)
}
