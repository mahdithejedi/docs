package main

import (
	"fmt"
	core "maelstormcore"
)

func main() {
	h := core.InitHandler()
	h.AddHandler(core.Init, func(NodeID string, message core.InputBody) (output core.OutputBody) {
		if core.Nodes.AddNode(NodeID) == true {
			core.GlobalLogChannel <- fmt.Sprintf("Node %s initiated", message.NodeID)
		}
		output = core.OutputBody{
			Type:      string(core.InitOk),
			InReplyTo: message.MsgID,
		}
		return
	})
	h.AddHandler(core.Echo, func(NodeID string, message core.InputBody) core.OutputBody {
		node := core.Nodes.GetNode(NodeID)
		return core.OutputBody{
			Type:      string(core.EchoOk),
			MsgID:     node.MsgID(),
			InReplyTo: message.MsgID,
			Echo:      node.Echo(message.Echo),
		}
	})
	h.Run()

}
