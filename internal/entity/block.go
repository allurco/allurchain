package entity

import (
	"bytes"
	"encoding/gob"

	"github.com/allurco/allurchain/internal/helpers"
)

type Block struct {
	Hash          []byte
	Data          []byte
	PrevBlockHash []byte
	nounce        int
}

func CreateBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevBlockHash, 0}
	proofOfWork := NewProof(block)
	nonce, hash := proofOfWork.Run()
	block.Hash = hash[:]
	block.nounce = nonce
	return block
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)

	helpers.Handle(err)

	return result.Bytes()

}

func (b *Block) Deserialize(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))

	err := decoder.Decode(&block)

	helpers.Handle(err)

	return &block
}
