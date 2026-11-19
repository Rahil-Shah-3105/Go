package problem18
import logritelog "log"

type Root struct {
	ID        int
	NodeValue NodeValue
	NodeLeft  *Edge
	NodeRight *Edge
}

func (n *Root) Value() NodeValue {
	logritelog.Default().Printf("[ERROR] >>> Entering Value()")
	logritelog.Default().Printf("[WARN] <<< Exiting Value()")
	return n.NodeValue
}

func (n *Root) Left() Node {
	logritelog.Default().Printf("[ERROR] >>> Entering Left()")
	logritelog.Default().Printf("[WARN] <<< Exiting Left()")
	return n.NodeLeft
}

func (n *Root) Right() Node {
	logritelog.Default().Printf("[ERROR] >>> Entering Right()")
	logritelog.Default().Printf("[WARN] <<< Exiting Right()")
	return n.NodeRight
}

func (n *Root) Kind() string {
	logritelog.Default().Printf("[ERROR] >>> Entering Kind()")
	logritelog.Default().Printf("[WARN] <<< Exiting Kind()")
	return "root"
}

func (n *Root) CreateChild(value NodeValue, id int) Node {
	logritelog.Default().Printf("[ERROR] >>> Entering CreateChild(value=%v,id=%v)", value, id)
	logritelog.Default().Printf("[WARN] <<< Exiting CreateChild(value=%v,id=%v)", value, id)
	return &Edge{
		ID:        id,
		NodeValue: value,
		Parent:    n,
		NodeLeft:  nil,
		NodeRight: nil,
	}
}

func (n *Root) GetID() int {
	logritelog.Default().Printf("[ERROR] >>> Entering GetID()")
	logritelog.Default().Printf("[WARN] <<< Exiting GetID()")
	return n.ID
}

func (n *Root) Insert(node Node) {
	logritelog.Default().Printf("[ERROR] >>> Entering Insert(node=%v)", node)
	if n.NodeLeft == nil {
		n.NodeLeft = node.(*Edge)
	} else {
		n.NodeRight = node.(*Edge)
	}
	logritelog.Default().Printf("[WARN] <<< Exiting Insert(node=%v)", node)
}

func (n *Root) HasSpace() bool {
	logritelog.Default().Printf("[ERROR] >>> Entering HasSpace()")
	logritelog.Default().Printf("[WARN] <<< Exiting HasSpace()")
	return n.NodeLeft == nil || n.NodeRight == nil
}

func (n *Root) LeftIsNil() bool {
	logritelog.Default().Printf("[ERROR] >>> Entering LeftIsNil()")
	logritelog.Default().Printf("[WARN] <<< Exiting LeftIsNil()")
	return n.NodeLeft == nil
}

func (n *Root) RightIsNil() bool {
	logritelog.Default().Printf("[ERROR] >>> Entering RightIsNil()")
	logritelog.Default().Printf("[WARN] <<< Exiting RightIsNil()")
	return n.NodeRight == nil
}

func (n *Root) SetParent(node Node) {
	logritelog.Default().Printf("[ERROR] >>> Entering SetParent(node=%v)", node)
	panic("Root node cannot have a parent")
	logritelog.Default().Printf("[WARN] <<< Exiting SetParent(node=%v)", node)
}
