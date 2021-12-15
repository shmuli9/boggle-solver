package wordtree

type WTNode struct {
	data     string
	isWord   bool
	void     bool
	children map[string]WTNode
}

func NewWTNode(str string) WTNode {
	return WTNode{data: str, isWord: false, void: false, children: make(map[string]WTNode)}
}
