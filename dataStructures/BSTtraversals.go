package main

import (
	"fmt"
)

type Node struct {
	Value int
	Left *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func main() {
	// values := []int{10,9,8,7,6,5,4,3,2}
	// tree := &Tree{}
	// for i := range values {
	// 	tree.Insert(values[i])
	// }
	tree := &Tree{}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(15)
	tree.Insert(14)
	tree.Insert(16)

	tree.inOrderTraversal(tree.Root)
	
}

func (n *Node) Insert(val int) error {
	
	switch {
	case val == n.Value:
		return nil
	case val < n.Value:
		if n.Left == nil {
			n.Left = &Node{Value: val}
			return nil
		}
		return n.Left.Insert(val)
		
	case val > n.Value:
		if n.Right == nil {
			n.Right = &Node{Value: val}
			return nil
		}
		return n.Right.Insert(val)
	}
	return nil
}

func(t *Tree) Insert(val int) error {
	if t.Root == nil {
		t.Root = &Node{Value: val}
		return nil
	}
	
	return t.Root.Insert(val)
}

// This function traverses a binary tree in preorder (Rooot, L ,R)
// For a BST with integer keys, it is often most useful to 
// utilise in-order tree traversal (tree walk). 
// Such a tree traversal will give the keys in ascending sorted order.

func (t *Tree) inOrderTraversal(n *Node) {
	if n == nil {
		return
	}
	//PreOrder
	//fmt.Print("InorderTraversal: ",n.Value, "\t")
	t.inOrderTraversal(n.Left)
	//InOrder
	fmt.Print(n.Value, "\t")
	t.inOrderTraversal(n.Right)
	//PostORder
	//fmt.Print(n.Value, "\t")
}

func (t *Tree) Traverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)
}

// Flatten: return inorder slice of keys
// func (t *Tree) Flatten() []int {
// 	var keys []int
// 	fn := func(n *Node) {
// 		keys = append(keys, n.key)
// 	}
// 	t.root.inorder(fn)
// 	return keys
// }

// func (t *Tree) findMax() (int, error) {
// 	if t.root == nil {
// 		return nil, fmt.Error("empty tree")
// 	}
// 	return t.Root.Right.findMax()

// }


	