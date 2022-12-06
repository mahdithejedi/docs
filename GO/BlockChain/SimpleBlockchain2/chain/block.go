package chain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
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
	data      IData
}

type blockchain struct {
	blocks []*block
}

func (b *block) setHash() {
	timestamp := []byte(strconv.FormatInt(b.timestamp, 10))
	headers := bytes.Join([][]byte{b.preBlock, b.data.GetByte(), timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.hash = hash[:]
}

func (b *block) GetData() string {
	return b.data.GetData()
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
	_block.setHash()
	return _block
}

func (b *blockchain) GetBlocks() []*block {
	return b.blocks
}
