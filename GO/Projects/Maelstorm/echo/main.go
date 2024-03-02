package main

import (
	"log"
	core "maelstormcore"
	"os"
)

func init() {
	log.Printf("RUNNED!!!!!!")
}

func main() {
	log.Printf("NODES ADDRESS IS %s and process ID %d 1", &core.Nodes, os.Getppid())
	h := core.InitHandler(true)
	log.Printf("NODES ADDRESS IS %s and process ID %d 2", &core.Nodes, os.Getppid())
	h.AddHandler(core.Echo, func(NodeID string, message core.InputBody) core.OutputBody {
		node := core.Nodes.GetNode(NodeID)
		return core.OutputBody{
			Type:      string(core.EchoOk),
			MsgID:     node.MsgID(),
			InReplyTo: message.MsgID,
			Echo:      node.Echo(message.Echo),
		}
	})
	log.Printf("NODES ADDRESS IS %s and process ID %d 3", &core.Nodes, os.Getppid())
	h.Run()
	log.Printf("NODES ADDRESS IS %s and process ID %d 4", &core.Nodes, os.Getppid())
}
