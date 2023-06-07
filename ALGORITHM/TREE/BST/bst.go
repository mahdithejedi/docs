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

func (n *Node) Search(search_value int) *Node {
	// time complexity -> o(h) * h: hiehgt of the tree
	if n.value == search_value {
		return n
	}
	if n.value > search_value && n.left != nil {
		return n.left.Search(search_value)
	}
	if n.value < search_value && n.right != nil {
		return n.right.Search(search_value)
	}
	return nil
}

func (n *Node) InorderTraverse(output *[]int) {
	// time complexity -> o(n)
	if n.left != nil {
		n.left.InorderTraverse(output)
	}
	if n.value == 0 {
		return
	} else {
		*output = append(*output, n.value)
	}
	if n.right != nil {
		n.right.InorderTraverse(output)
	}
}

func (n *Node) Max() int {
	for ; n.right != nil; n = n.right {

	}
	return n.value
}

func (n *Node) Min() int {
	for ; n.left != nil; n = n.left {

	}
	return n.value
}

func (n *Node) MinWithPointer() *Node {
	for ; n.left != nil; n = n.left {

	}
	return n
}

func (n *Node) Delete(delete_value int) bool {
	node := n.Search(delete_value)
	fmt.Print(node)
	if node == nil {
		return false
	} else if node.left == nil && node.right == nil {
		// no child !
		fmt.Println("1")
		defer node.Reset()
		return true
	} else if node.left != nil && node.right == nil {
		// just left child is availabe!
		fmt.Println("2")
		node = node.left
		return true
	} else if node.right != nil && node.left == nil {
		// just right child is available
		fmt.Println("3")
		node = node.right
		return true

	} else {
		fmt.Println("4")
		// has two children
		// successor := node.right.MinWithPointer()
		del_node := node.right.MinWithPointer()
		defer del_node.Reset()
		node = del_node
		return true
	}

}

func main() {
	n := Node{}
	// node_values := []int{100, 500, 20, 10, 30, 40}
	// node_values := []int{2, 252, 401, 398, 330, 344, 397, 363}
	// node_values := []int{10, 5, 1, 7, 40, 50}
	// node_values := []int{3, 0, 2, 9, 3, 4, 3, 0, 7, 6, 7, 1, 6, 6, 7, 5, 3, 1, 4, 1, 3, 2, 5, 7, 8, 5, 3, 5, 4, 8, 6, 4, 3, 4, 8, 5, 6, 4, 5}
	node_values := []int{15, 10, 20, 8, 12, 25}
	n.ListInsert(node_values)
	success := n.Delete(15)
	if success != true {
		fmt.Println("not !")

	}
	// fmt.Println(n.value)
	output := []int{}
	n.InorderTraverse(&output)
	fmt.Print(output)
	// n.Search(363)
	// inorder := []int{}
	// n.InorderTraverse(&inorder)
	// fmt.Print(n.Min())
	// node, existence := n.Search(85)
	// fmt.Print(node.value, existence)
}
