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

func main() {
	// __source__ = 'https://github.com/Jeiwan/blockchain_go'
	blockchain := chain.New()
	_chaindata := chainData{data: "First"}
	fmt.Println(_chaindata, _chaindata.IData, (_chaindata).IData.GetByte())
	for _, block := range blockchain.GetBlocks() {
		fmt.Println(block.GetData())
	}

}
