package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

var Blockchain []Block

func (b *Block) calculateHash() {
	record := string(b.Index) + b.Timestamp + string(b.BPM) + b.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	b.Hash = hex.EncodeToString(h.Sum(nil))
}

func CreateBlock(prevBlock Block, BPM int) (Block, error) {
	t := time.Now()
	new_block := &Block{
		Index:     prevBlock.Index + 1,
		Timestamp: t.String(),
		BPM:       BPM,
		PrevHash:  prevBlock.Hash,
	}
	new_block.calculateHash()
	return *new_block, nil
}
