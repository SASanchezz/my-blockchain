package blockchain

import (
	"bytes"         // need to convert data into byte in order to be sent on the network, computer understands better the byte(8bits)language
	"crypto/sha256" //crypto library to hash the data
	"fmt"
	"math"
	"math/big"
	"strconv" // for conversion
	"time"
	// the time for our timestamp
)

// GetHash - we just concatenate all the data and hash it to obtain the block hash
func (b *Block) GetHash() []byte {
	headers := b.GetBytes()
	hash := sha256.Sum256(headers)
	return hash[:]
}

func NewBlock(from string, to string, amount uint64, data string, prevBlockHash []byte) *Block {
	block := &Block{
		from,
		to,
		amount,
		data,
		time.Now().Unix(),
		prevBlockHash,
		[]byte{},
		0,
		0,
	}
	return block
}

func (b *Block) Mine() {
	const (
		targetBits = 24            // Number of leading zeros required in the hash
		maxNonce   = math.MaxInt64 // Maximum value for the nonce
	)

	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	for b.Nonce < maxNonce {
		hash := b.GetHash()

		var hashInt big.Int
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(target) == -1 {
			fmt.Printf("Block mined! Nonce: %d\n", b.Nonce)
			b.BlockHash = hash[:]
			break
		} else {
			b.Nonce++
		}
	}
}

func NewGenesisBlock() *Block {
	var genesisBlock = NewBlock("", "", 0, "Genesis Block1", []byte{})
	genesisBlock.Mine()

	return genesisBlock
}

// GetBytes - we just concatenate all the data into a byte array
func (b *Block) GetBytes() []byte {
	return bytes.Join([][]byte{
		[]byte(b.From),
		[]byte(b.To),
		[]byte(strconv.FormatUint(b.Amount, 10)),
		[]byte(b.Data),
		[]byte(strconv.FormatInt(b.Timestamp, 10)),
		[]byte(b.PreviousBlockHash),
		[]byte(strconv.FormatUint(b.Nonce, 10)),
	}, []byte{})
}

func (b *Block) String() string {
	var result string
	result += fmt.Sprintf("From : %s \n", b.From)
	result += fmt.Sprintf("To : %s \n", b.To)
	result += fmt.Sprintf("Amount : %d \n", b.Amount)
	result += fmt.Sprintf("Transaction data: %s\n", b.Data)
	result += fmt.Sprintf("Timestamp : %d \n", b.Timestamp)
	result += fmt.Sprintf("Hash of the block %x\n", b.BlockHash)
	result += fmt.Sprintf("Hash of the previous Block: %x\n", b.PreviousBlockHash)

	return result
}
