package entity

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
