package main

import (
	"blockchain/chain"
	"fmt"
)

type chainData struct {
	chain.IData
	data string
}

func (c *chainData) setData(data string) *chainData {
	c.data = data
	return c
}

func (c *chainData) GetByte() []byte {
	return []byte(c.data)
}

func (c *chainData) GetData() string {
	return c.data
}

func main() {
	// __source__ = 'https://github.com/Jeiwan/blockchain_go'
	blockchain := chain.New()
	blockchain.NewBlock(&chainData{data: "First"})
	blockchain.NewBlock(&chainData{data: "Second"})
	blockchain.NewBlock(&chainData{data: "Third"})
	blockchain.NewBlock(&chainData{data: "Fourth"})
	for _, block := range blockchain.GetBlocks() {
		fmt.Println((*block).GetData())
	}

}
