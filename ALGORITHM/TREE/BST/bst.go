package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) insertInTree(new_value int) *Node {
	if n == nil {
		n = &Node{value: new_value}
	} else if n.value > new_value {
		n.left = n.left.insertInTree(new_value)
	} else {
		n.right = n.right.insertInTree(new_value)
	}
	return n
}

func (n *Node) Insert(new_value int) {
	// initiation!
	if n.value == 0 {
		n.value = new_value
		return
	}
	n.insertInTree(new_value)

}

func main() {
	n := Node{}
	n.Insert(12)
	n.Insert(10)
	n.Insert(14)
	n.Insert(11)
	fmt.Println(n.left.right.value)
}
