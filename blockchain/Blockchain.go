package blockchain

import (
	"bytes"
	"fmt"
)

func (bl *Blockchain) AddBlock(from string, to string, amount int64, data string) *Block {
	previousBlock := bl.Chain[len(bl.Chain)-1]                            // the previous block is needed, so let's get it
	newBlock := NewBlock(from, to, amount, data, previousBlock.BlockHash) // create a new block containing the data and the hash of the previous block
	newBlock.Mine()
	bl.Chain = append(bl.Chain, newBlock) // add that block to the Chain to create a Chain of blocks

	return newBlock
}

func NewBlockchain() *Blockchain { // the function is created
	return &Blockchain{[]*Block{NewGenesisBlock()}} // the genesis block is added first to the Chain
}

func (bl *Blockchain) IsValid() bool {
	for i := range bl.Chain[1:] {
		previousBlock := bl.Chain[i]
		currentBlock := bl.Chain[i+1]
		currentBlockHash := currentBlock.GetHash()   //Error
		previousBlockHash := previousBlock.GetHash() // Hashes of same block are different
		if !bytes.Equal(currentBlock.BlockHash, currentBlockHash) || !bytes.Equal(currentBlock.PreviousBlockHash, previousBlockHash) {
			return false
		}
	}
	return true
}

func (bl Blockchain) String() string {
	var result string
	for i, b := range bl.Chain { // iterate on each block
		result += fmt.Sprintf("Block ID : %d \n", i)
		result += b.String()
		result += "\n"
	}

	return result
}
