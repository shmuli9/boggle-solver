package wordtree

type WTNode struct {
	data     string
	isWord   bool
	void     bool
	children NodeMap
}

func NewWTNode(str string) *WTNode {
	return &WTNode{data: str, isWord: false, void: false, children: make(NodeMap)}
}

func (node WTNode) String() string {
	output := ""
	output += node.data
	if node.isWord {
		output += "*"
	}
	output += node.children.String()
	return output
}
