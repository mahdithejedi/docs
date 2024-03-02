package maelstormcore

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

type Node struct {
	ID          string
	ReceiveChan chan InputMsg
	SendChan    chan string
	neighbors   []*Node
	latestMsgID int
	lock        sync.Mutex
}

func raiseNodeDoesNotExists(nodeID string) {
	LogError(
		errors.New(
			fmt.Sprintf("Node ID %s does not exists", nodeID),
		),
	)
}

func (n *Node) Close() {
	close(n.ReceiveChan)
	close(GlobalLogChannel)
}

func (n *Node) Echo(msg string) string {
	return msg
}

func (n *Node) SetNeighbor(node *Node) {
	n.neighbors = append(n.neighbors, node)
}

func (n *Node) SetNeighbors(node []*Node) {
	n.neighbors = append(n.neighbors, node...)
}

func (n *Node) GetNeighbors() []*Node {
	return n.neighbors
}

func (n *Node) increaseID() {
	n.lock.Lock()
	defer n.lock.Unlock()
	n.latestMsgID++
}

func (n *Node) MsgID() (_id int) {
	_id = (*n).latestMsgID
	n.increaseID()
	return
}
func (n *Node) ReceiveBroadcast(msg string) {

}

func (n *Node) Broadcast(msg string) {
	for _, node := range n.neighbors {
		node.ReceiveBroadcast(msg)
	}
}

// **********************
// ****** Node Map ******
// **********************

type NodeMap struct {
	nodes map[string]*Node
}

var Nodes *NodeMap
var runOnce sync.Once

func InitNodeHandler() {
	runOnce.Do(func() {
		Nodes = &NodeMap{
			nodes: make(map[string]*Node),
		}
	})
}

func (n *NodeMap) AddNode(NodeID string) bool {
	_, exists := n.nodes[NodeID]
	if exists {
		return false
	}
	n.nodes[NodeID] = &Node{
		ID:          NodeID,
		ReceiveChan: make(chan InputMsg),
		SendChan:    GlobalLogChannel,
		latestMsgID: 1,
	}
	return true

}

func (n *NodeMap) GetNode(NodeID string) *Node {
	log.Printf("ALL NODES ARE %#v while getting node %s", n.nodes, NodeID)
	node, exists := n.nodes[NodeID]
	if exists != true {
		raiseNodeDoesNotExists(NodeID)
	}
	return node
}

func (n *NodeMap) GetNodes(NodeIDs []string) []*Node {
	nodeObjects := make([]*Node, len(NodeIDs))
	for _, nodeID := range NodeIDs {
		nodeObjects = append(nodeObjects, n.GetNode(nodeID))
	}
	return nodeObjects
}

func (n *NodeMap) Close() {
	for _, node := range n.nodes {
		//fmt.Printf("Going to stop %s\n", nodeID)
		node.Close()
		//fmt.Printf("Node %s stopped successfully\n", nodeID)
	}
}
