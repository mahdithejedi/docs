package main

import (
	"fmt"
	core "maelstormcore"
)

func main() {
	handler := core.InitHandler(true)
	handler.AddHandler(core.Topology, func(currentNodeID string, message core.InputBody) core.OutputBody {
		for rootNode, parents := range message.Topology {
			if rootNode == currentNodeID {
				node := core.Nodes.GetNode(rootNode)
				node.SetNeighbors(
					core.Nodes.GetNodes(
						parents,
					),
				)
				core.GlobalLogChannel <- fmt.Sprintf("My neighbors are %v\n",
					node.GetNeighbors(),
				)
			}
		}
		return core.OutputBody{
			Type:      string(core.TopologyOK),
			InReplyTo: message.MsgID,
		}
	})
	handler.Run()

}
