package main

import (
	"fmt"
	"errors"
)
// `Node` contains the search value, some data, a left child node, and a right child node.
type Node struct {
	Value int
	Data  int
	Left  *Node
	Right *Node
}
// A `Tree` basically consists of a root node.

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
	tree :- &Tree{}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(15)
	tree.Insert(14)
	tree.Insert(16)
	tree.Insert(18)
	tree.Insert(19)
	 //10
	//5    /15\
//4 6 // 14    16\
				//18
					//19
	// err := tree.Root.Delete(5,tree.Root)
	// if err != nil {
	// 	// Errors.New("sdf")
	// }
	tree.inOrderTraversal(tree.Root)
	tree.bfs(tree.Root, 16)
	//first := tree.Root
	
}

// func (t *Tree) cycleDetect(){
// 	visited := make(map[*Node]bool)
// 	for n := first; n != nil; n = n.Right {
// 		if visited[n] {
// 			fmt.Println("cycle detected")	
// 			break
// 		}
// 		visited[n] = true
// 		fmt.Println(n.Value, "..", n.Right)
// 	}
// }

func(t *Tree) Insert(val int) error {
	if t.Root == nil {
		t.Root = &Node{Value: val}
		return nil
	}
	
	return t.Root.Insert(val)
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

// func (n *Node) w(s string) (int, bool) {
// 	if n == nil {
// 		return "nil", false
// 	}
// 	switch {
// 	case s == n.value:
// 		return n.Data, true
// 	case s < n.value:
// 		return n.Left.Find(s)
// 	default:
// 		return n.Right.Find(s)
// 	}
// }

// Delete Helper Functions
// Either findmax = traverse left subtree, or find min if you traverse the right side.

func (n *Node) findMax(parent *Node) (*Node, *Node) {
	//parent has no child. return parent
	if n == nil {
		return nil, parent
	}
	//child has no child. return parent and child.
	if n.Right == nil {
		return n, parent
	}
	return n.Right.findMax(n)
}

// func (n *Node) findMax(n *Node) {

// 	for n.Right == nil {
// 		findMax(n.Right)
// 	}
// 	if n.Left != nil {
// 		findMax(n.Left)
// 	}
// 	return n

// }


func (n *Node) replaceNode(parent, replacement *Node) error {
	if n == nil {
		errors.New("replaceNode() not allowed on a nil node")
	}
	if parent.Left == n {
		parent.Left = replacement
	}

	parent.Right = replacement
	return nil
}	

func (n *Node) Delete(s int, parent *Node) error {
	if n == nil {
		errors.New("nothing to delete")
	}

	switch {
	case s < n.Value:
		return n.Left.Delete(s, n)
	case s > n.Value:
		return n.Right.Delete(s, n)
	default:
		if n.Left == nil && n.Right == nil {
			n.replaceNode(parent,nil)
		}

		if n.Left == nil {
			n.replaceNode(parent, n.Right)
			return nil
		}
		if n.Right == nil {
			n.replaceNode(parent, n.Left)
			return nil
		}
		//If node has 2 children
		replacement, replParent := n.Left.findMax(n)
		n.Value = replacement.Value
		n.Data = replacement.Data
		
		return replacement.Delete(replacement.Value, replParent)
	}
}

//that the amount of space we need in terms of memory depends upon the height of our tree, or O(h). 
func (t *Tree) inOrderTraversal(n *Node) {
	if n == nil {
		return
	}
	//PreOrder
	// fmt.Print("InorderTraversal: ",n.Value, "\t")
	t.inOrderTraversal(n.Left)
	//InOrder
	// fmt.Print(n.Value, "\t")
	t.inOrderTraversal(n.Right)
	//PostORder
	//fmt.Print(n.Value, "\t")
}
// func (t *Tree) Delete(s int) error {

// 	if t.Root == nil {
// 		return errors.New("Cannot delete from an empty tree")
// 	}
// 	fakeParent := &Node{Right: t.Root}
// 	err := t.Root.Delete(s, fakeParent)
// 	if err != nil {
// 		return err
// 	}
// 	if fakeParent.Right == nil {
// 		t.Root = nil
// 	}
// 	return nil
// }

// func main() {

// 	// Set up a slice of strings.
// 	values := []string{"d", "b", "c", "e", "a"}
// 	data := []string{"delta", "bravo", "charlie", "echo", "alpha"}

// 	// Create a tree and fill it from the values.
// 	tree := &Tree{}
// 	for i := 0; i < len(values); i++ {
// 		err := tree.Insert(values[i], data[i])
// 		if err != nil {
// 			log.Fatal("Error inserting value '", values[i], "': ", err)
// 		}
// 	}


// 	// Print the sorted values.
// 	fmt.Print("Sorted values: | ")
// 	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
// 	fmt.Println()

// }


/* BFS: O(n) Space Complexity: O(n)
* 1. enqueue root node, 
* 2. print out node
* enqueue left child,
* enqueue right child
* dequeue node
*/
func (t *Tree) bfs(n *Node, value int) int {

	queue := []*Node{}
    queue = append(queue, n)
    
    for 0 < len(queue) {
			node := queue[0]
			fmt.Println(node.Value)
			
            // if (node.Value == value) {
				//     return node.Value
				// }
				if (node.Left != nil) {
					queue = append(queue, node.Left)
				}
				if (node.Right != nil) {
					queue = append(queue, node.Right)
				}
				queue = queue[1:]
        // }
    }
    return 0
}