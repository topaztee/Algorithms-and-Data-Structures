package main

import 
(
	"fmt"
)

type Node struct {
	value int
	next *Node
	prev *Node
}

type List struct {
	head *Node
	tail *Node
	count int
}
func main(){

}

func (dll *DLinkedList) Delete(del *Node) {
	nodeToDel:= dll.Find(del)
	prev := nodeToDel.prev
}