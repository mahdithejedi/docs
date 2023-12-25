package maelstormcore

type Node struct {
	ID          string
	ReceiveChan chan InputMsg
	SendChan    chan string
}

type NodeMap map[string]*Node

var Nodes = make(NodeMap)

func (n *Node) Close() {
	close(n.ReceiveChan)
	close(GlobalLogChannel)
}

func (n *NodeMap) AddNode(Body InputBody) bool {
	if (*n)[Body.NodeID] != nil {
		return false
	}
	(*n)[Body.NodeID] = &Node{
		ID:          Body.NodeID,
		ReceiveChan: make(chan InputMsg),
		SendChan:    GlobalLogChannel,
	}
	return true

}

func (n *NodeMap) Close() {
	for _, node := range *n {
		//fmt.Printf("Going to stop %s\n", nodeID)
		node.Close()
		//fmt.Printf("Node %s stopped successfully\n", nodeID)
	}
}
