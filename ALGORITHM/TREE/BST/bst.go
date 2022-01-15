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

func (n *Node) ListInsert(new_values []int) {
	for _, element := range new_values {
		n.Insert(element)
	}
}

func (n *Node) Reset() {
	n.value = 0
	n.left = nil
	n.right = nil
}

func (n *Node) Search(search_value int) (*Node, bool) {
	if n.value == search_value {
		return n, true
	}
	if n.value > search_value && n.left != nil {
		return n.left.Search(search_value)
	}
	if n.value < search_value && n.right != nil {
		return n.right.Search(search_value)
	}
	return &Node{}, false
}

// func (n *Node) Delete(delete_value int) bool {
// 	node, exitence := n.Search(delete_value)
// 	if exitence != false {
// 		return false
// 	}
// 	if node.left == nil && n.right == nil {
// 		node.value = 0
// 	} else if node.left != nil && node.right == nil {
// 		node = node.left
// 	} else if node.right != nil && node.left == nil{
// 		node = node.right
// 	}
// 	else{

// 	}

// }

func main() {
	n := Node{}
	node_values := []int{100, 500, 20, 10, 30, 40}
	n.ListInsert(node_values)
	node, existence := n.Search(85)
	fmt.Print(node.value, existence)

}
