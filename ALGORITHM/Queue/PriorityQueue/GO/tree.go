package main

var initTrackerLen uint8 = 12

type MaxHeapTree struct {
	tracker  []*Task
	lastNode *MaxHeapTree
	Node     *Task
	Left     *Task
	Right    *Task
}

func (tree *MaxHeapTree) Insert(priority uint8, method func()) {
	if tree.Node == nil {
		tree.Node = InitTask(priority)
		tree.Node.Push(method)
		tree.lastNode = tree
	}
	node := tree.find(priority)
	if node != nil {
		node.Push(method)
	} else {
		tree.insert(priority, method)
	}
}
func (tree *MaxHeapTree) insert(priority uint8, method func()) {
	if tree.lastNode.Left != nil && tree.lastNode.Right != nil {
		reloadLastNode(tree)
	}
	if tree.lastNode.Left == nil {
		tree.lastNode.Left = InitTask(priority)
		tree.lastNode.Left.Push(method)
		tree.trackInsert(priority, tree.lastNode.Left)
	} else if tree.lastNode.Right == nil {
		tree.lastNode.Right = InitTask(priority)
		tree.lastNode.Right.Push(method)
		tree.trackInsert(priority, tree.lastNode.Right)
	}
}

func reloadLastNode(tree *MaxHeapTree) {
	if tree.Left == tree.lastNode.Node {
		tree.lastNode = tree.Right
	}

}

func (tree *MaxHeapTree) find(priority uint8) *Task {
	return tree.tracker[priority%initTrackerLen]
}

func (tree *MaxHeapTree) trackInsert(priority uint8, node *Task) {
	tree.tracker[priority%initTrackerLen] = node
}

func InitTree() *MaxHeapTree {
	return &MaxHeapTree{
		tracker: make([]*Task, initTrackerLen),
	}
}
