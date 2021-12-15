package wordtree

type NodeMap map[string]*WTNode

type WTNode struct {
	data     string
	isWord   bool
	void     bool
	children NodeMap
}

func NewWTNode(str string) *WTNode {
	return &WTNode{data: str, isWord: false, void: false, children: make(NodeMap)}
}
