package main

import "fmt"

// Node sdd
type Node struct {
	char   rune
	Data   interface{}
	parent *Node
	Depth  int
	childs map[rune]*Node
	term   bool
}

// Trie sss
type Trie struct {
	root *Node
	size int
}

// NewNode ss
func NewNode() *Node {
	return &Node{
		childs: make(map[rune]*Node, 32),
	}
}

// NewTrie ss
func NewTrie() *Trie {
	return &Trie{
		root: NewNode(),
	}
}



func main() {
	a := Node{
		char: 10,
		Data: 1,
	}
	b := Node{
		parent: &a,
	}
	fmt.Printf("%v\n", b.parent)
}
