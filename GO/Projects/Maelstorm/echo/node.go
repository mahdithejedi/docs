package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Node struct {
	ID          string
	ReceiveChan chan InputMsg
	SendChan    chan string
}

var MaxOutputBufferSizw = 120
var GlobalOutPutHandler = make(chan string, MaxOutputBufferSizw)

type NodeMap map[string]*Node

var Nodes = make(NodeMap)

func (n *Node) Close() {
	close(n.ReceiveChan)
	close(GlobalOutPutHandler)
}

func replyInitOkMsg(msgId int, output OutputMsg) string {
	output.Body = OutputBody{
		Type:      "init_ok",
		InReplyTo: msgId,
	}
	body, err := json.Marshal(output)
	LogError(err)
	return fmt.Sprintf("Sent %s", string(body))
}

func (n *NodeMap) AddNode(Body InputBody, output OutputMsg) {
	if (*n)[Body.NodeID] != nil {
		return
	}
	(*n)[Body.NodeID] = &Node{
		ID:          Body.NodeID,
		ReceiveChan: make(chan InputMsg),
		SendChan:    GlobalOutPutHandler,
	}
	GlobalOutPutHandler <- fmt.Sprintf("Node %s initialized", Body.NodeID)
	GlobalOutPutHandler <- replyInitOkMsg(Body.MsgID, output)

}

func formatOutputMsg(msg string) string {
	if strings.HasSuffix(msg, "\n") {
		return msg[len(msg) : len(msg)-len("\n")]
	}
	return msg
}

func (n *NodeMap) InitOutputHandler() {
	for {
		msg := <-GlobalOutPutHandler
		if msg == "0" {
			break
		}
		var err error
		log.Println(formatOutputMsg(msg))
		LogError(err)
	}
	defer close(GlobalOutPutHandler)
}

func (n *NodeMap) Close() {
	for _, node := range *n {
		//fmt.Printf("Going to stop %s\n", nodeID)
		node.Close()
		//fmt.Printf("Node %s stopped successfully\n", nodeID)
	}
}
