package main

type Node struct {
	value int
	left  *Node
	right *Node
}

func FindBestPosition(n *Node, new_value int) Node {
	if new_value == n.value {
		return *n
	} else if n.value > new_value {
		return FindBestPosition(n.left, new_value)
	} else {
		return FindBestPosition(n.right, new_value)
	}
}

func (n *Node) Insert(new_value int) {
	if n.value == 0 {
		n.value = new_value
		return
	}
	// postion_to_insert := FindBestPosition(n, new_value)
	// postion_to_insert.value = new_value

}

func main() {
	n := Node{}
	n.Insert(12)
	n.Insert(10)
	n.Insert(14)
}
