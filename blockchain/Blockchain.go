package blockchain

import (
	"bytes"
	"fmt"
)

func (bl *Blockchain) AddBlock(from string, to string, amount uint64, data string) *Block {
	newBlock := NewBlock(from, to, amount, data) // create a new block containing the data and the hash of the previous block
	newBlock.Mine()
	bl.Chain = append(bl.Chain, newBlock) // add that block to the Chain to create a Chain of blocks

	return newBlock
}

func (bl *Blockchain) AppendBlock(b *Block) {
	bl.Chain = append(bl.Chain, b)
}

func (bl *Blockchain) RemoveLastBlock() {
	bl.Chain = bl.Chain[:len(bl.Chain)-1]
}

func NewBlockchain() *Blockchain { // the function is created
	return &Blockchain{
		ConfirmationsNeeded: 2,
		Chain:               []*Block{NewGenesisBlock()},
	} // the genesis block is added first to the Chain
}

func (bl *Blockchain) Has(b *Block) bool {
	for _, block := range bl.Chain {
		if bytes.Equal(block.BlockHash, b.BlockHash) {
			return true
		}
	}
	return false
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
