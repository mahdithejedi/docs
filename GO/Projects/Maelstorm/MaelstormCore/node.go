package maelstormcore

type Node struct {
	ID          string
	ReceiveChan chan InputMsg
	SendChan    chan string
	latestMsgID int
}

type NodeMap map[string]*Node

var Nodes = make(NodeMap)

func (n *Node) Close() {
	close(n.ReceiveChan)
	close(GlobalLogChannel)
}

func (n *Node) Echo(msg string) string {
	return msg
}

func (n *Node) MsgID() (_id int) {
	_id = (*n).latestMsgID
	n.latestMsgID++
	return
}

func (n *NodeMap) AddNode(NodeID string) bool {
	if (*n)[NodeID] != nil {
		return false
	}
	(*n)[NodeID] = &Node{
		ID:          NodeID,
		ReceiveChan: make(chan InputMsg),
		SendChan:    GlobalLogChannel,
		latestMsgID: 1,
	}
	return true

}

func (n *NodeMap) GetNode(NodeID string) *Node {
	return (*n)[NodeID]
}

func (n *NodeMap) Close() {
	for _, node := range *n {
		//fmt.Printf("Going to stop %s\n", nodeID)
		node.Close()
		//fmt.Printf("Node %s stopped successfully\n", nodeID)
	}
}
