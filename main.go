package main

import (
	"boggle-solver/wordtree"
	"log"
)

func main() {
	t := wordtree.NewWordTree()
	testWords := []string{"test", "testing", "tertiary"}

	for _, word := range testWords {
		t.Add(word)
	}

	for _, word := range testWords {
		found := t.Find(word)
		log.Print(found, " ", word)
	}

	log.Print(t)
}
