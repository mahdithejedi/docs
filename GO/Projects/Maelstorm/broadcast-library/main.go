package main

import (
	"fmt"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"log"
)

func handler(msg maelstrom.Message) (err error) {
	err = nil
	fmt.Printf("\n ************** \n\n msg %s \n\n ************** \n \n", msg)
	return
}

func main() {
	n := maelstrom.NewNode()

	n.Handle("broadcast", handler)

	if err := n.Run(); err != nil {
		log.Fatal("error while running server", err)
	}
}
