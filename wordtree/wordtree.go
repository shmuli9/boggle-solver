package wordtree

import (
	"strings"
)

type WordTree struct {
	children  NodeMap
	voidNodes []WTNode
	voidWords []WTNode
}

func NewWordTree() WordTree {
	return WordTree{children: NodeMap{}, voidNodes: []WTNode{}, voidWords: []WTNode{}}
}

func (tree WordTree) String() string {
	return "WordTree[" + tree.children.String() + "]"
}

func (tree WordTree) Add(word string) {
	length := len(word)

	if length >= 3 {
		children := tree.children
		var node *WTNode
		letterIndex := 0

		for letterIndex < length {
			letter := strings.ToUpper(string(word[letterIndex]))

			if letter == "Q" {
				letter = "QU"
				letterIndex += 1
			}

			letterIndex += 1

			if _, exists := children[letter]; !exists {
				children[letter] = NewWTNode(letter)
			}

			node = children[letter]
			children = node.children
		}

		node.isWord = true
	}
}

func (tree WordTree) Find(word string) bool {
	length := len(word)

	if length >= 3 {
		children := tree.children
		var node *WTNode
		letterIndex := 0

		for letterIndex < length {
			letter := strings.ToUpper(string(word[letterIndex]))

			if letter == "Q" {
				letter = "QU"
				letterIndex += 1
			}

			letterIndex += 1

			if _, exists := children[letter]; !exists {
				return false
			}

			node = children[letter]
			children = node.children
		}

		return node.isWord
	}

	return true
}

func (tree WordTree) ResetTree() {

}

func (tree WordTree) BuildWordTree(minWordSize int) {}

func (tree WordTree) verifyWordTree() {}
