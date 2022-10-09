package helpers

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	first *Node
	head  *Node
}

func (l *LinkedList) Add(data interface{}) {
	node := Node{
		data: data,
	}
	if l.first == nil {
		l.first = &node
		l.head = l.first
		return
	}
	l.head.next = &node
}

func (l *LinkedList) Exists(data interface{}) bool {
	ls := l.first
	for ls != nil {
		if ls.data == data {
			return true
		}
		ls = ls.next
	}
	return false
}

func (l *LinkedList) Show() {
	ls := l.first
	for ls != nil {
		fmt.Print(ls.data.(string))
		ls = ls.next
	}
}
