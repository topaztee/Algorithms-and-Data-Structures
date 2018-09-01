// Without the tree struct you cant create a nil tree because node key will
// initialize to 0 rather than nil
package main

import (
	"fmt"
)

type Node struct {
    key    int
    left   *Node
    right  *Node
    parent *Node
}

type Tree struct {
	Root *Node
}

func main() {
	fmt.Println("Hello, playground")
	var n = &Tree{}
	
	fmt.Println(n.Root)
}
