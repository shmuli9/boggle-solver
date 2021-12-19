package main

import (
	"boggle-solver/wordtree"
	"log"
)

func main() {
	t := wordtree.NewWordTree()
	t.Add("test")
	t.Add("testing")
	t.Add("tertiary")
	log.Print(t)
}
