package main

import "fmt"

type Node struct {
	next *Node
	key  int
}

type List struct {
	head  *Node
	first *Node
}

func (L *List) Insert(key int) bool {
	node := &Node{
		key: key,
	}
	if L.first == nil {
		L.first = node
	} else {
		L.head.next = node
	}
	L.head = node
	return true
}

func (L *List) Print() {
	l := L.first
	for l != nil {
		fmt.Println(l.key)
		l = l.next
	}
}

func main() {
	list := List{}
	list.Insert(20)
	list.Insert(23)
	list.Insert(32)
	list.Insert(46)
	list.Insert(57)
	list.Insert(60)
	list.Print()
}
