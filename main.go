package main

import (
	"boggle-solver/wordtree"
	"fmt"
	"log"
)

func main() {
	fmt.Println("working")
	t := wordtree.NewWordTree()
	t.Add("test")
	t.Add("testing")
	t.Add("tertiary")
	log.Print(t)
}
