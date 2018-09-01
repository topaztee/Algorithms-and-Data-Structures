package main

import (
	"fmt"
	"container/list"
)

type Node struct {
	data [int]string
	next *Node
	prev *Node
}

type DList struct {
	Head *Node
	Tail *Node
	count int
}

func (ll *DList) Insert(val int){
	new := Node{data: 4}
}

// Insert at end.
func (ll *DoublyLinkedList) Insert(val int) {
	newNode := &node{value: val}
	ll.insertNode(newNode)
}


func (ll *DoublyLinkedList) insertNode(newNode *node) {
	ll.count++
	if ll.tail == nil {
		ll.head, ll.tail = newNode, newNode
	} else {
		ll.tail.next = newNode
		newNode.prev = ll.tail
		ll.tail = newNode
	}
}