package problem18
import logritelog "log"

type Root struct {
	ID        int
	NodeValue NodeValue
	NodeLeft  *Edge
	NodeRight *Edge
}

func (n *Root) Value() NodeValue {
	logritelog.Default().Printf(">>> Entering Value()")
	return n.NodeValue
}

func (n *Root) Left() Node {
	logritelog.Default().Printf(">>> Entering Left()")
	return n.NodeLeft
}

func (n *Root) Right() Node {
	logritelog.Default().Printf(">>> Entering Right()")
	return n.NodeRight
}

func (n *Root) Kind() string {
	logritelog.Default().Printf(">>> Entering Kind()")
	return "root"
}

func (n *Root) CreateChild(value NodeValue, id int) Node {
	logritelog.Default().Printf(">>> Entering CreateChild(value=%v,id=%v)", value, id)
	return &Edge{
		ID:        id,
		NodeValue: value,
		Parent:    n,
		NodeLeft:  nil,
		NodeRight: nil,
	}
}

func (n *Root) GetID() int {
	logritelog.Default().Printf(">>> Entering GetID()")
	return n.ID
}

func (n *Root) Insert(node Node) {
	logritelog.Default().Printf(">>> Entering Insert(node=%v)", node)
	if n.NodeLeft == nil {
		n.NodeLeft = node.(*Edge)
	} else {
		n.NodeRight = node.(*Edge)
	}
	logritelog.Default().Printf("<<< Exiting Insert(node=%v)", node)
}

func (n *Root) HasSpace() bool {
	logritelog.Default().Printf(">>> Entering HasSpace()")
	return n.NodeLeft == nil || n.NodeRight == nil
}

func (n *Root) LeftIsNil() bool {
	logritelog.Default().Printf(">>> Entering LeftIsNil()")
	return n.NodeLeft == nil
}

func (n *Root) RightIsNil() bool {
	logritelog.Default().Printf(">>> Entering RightIsNil()")
	return n.NodeRight == nil
}

func (n *Root) SetParent(node Node) {
	logritelog.Default().Printf(">>> Entering SetParent(node=%v)", node)
	panic("Root node cannot have a parent")
	logritelog.Default().Printf("<<< Exiting SetParent(node=%v)", node)
}
