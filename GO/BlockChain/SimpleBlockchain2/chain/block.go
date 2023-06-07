package chain

import (
	"log"
	"time"
)

type IData interface {
	GetByte() []byte
	GetData() string
}

type block struct {
	timestamp int64
	preBlock  []byte
	hash      []byte
	Nonce     int
	data      IData
}

type blockchain struct {
	blocks []*block
}

func (b *block) GetData() string {
	return b.data.GetData()
}

func (b *block) GetByte() []byte {
	return b.data.GetByte()
}

func New() *blockchain {
	return &blockchain{}
}

func (b *blockchain) NewBlock(data IData) *block {
	addedBlock := newBlock(data, func() []byte {
		if len(b.blocks) == 0 {
			return []byte{}
		} else {
			return b.blocks[len(b.blocks)-1].hash
		}
	}())
	b.blocks = append(b.blocks, addedBlock)
	return addedBlock

}

func newBlock(data IData, preBlockHash []byte) *block {
	_block := &block{
		timestamp: time.Now().Unix(),
		preBlock:  preBlockHash,
		data:      data,
	}
	//_block.setHash()
	pow := NewPOW(_block)
	var err error
	_block.Nonce, _block.hash, err = pow.Run()
	if err != nil {
		log.Fatalln("Error", err.Error())
	}
	return _block
}

func (b *blockchain) GetBlocks() []*block {
	return b.blocks
}
