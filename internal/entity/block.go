package entity

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash          []byte
	Data          []byte
	PrevBlockHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevBlockHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevBlockHash}
	block.DeriveHash()
	return block
}
