package main

import (
	"my-blockchain/blockchain"
	"my-blockchain/p2p"
)

func main() {
	// p2p.StartListening("localhost:3001")

	// Connect to a peer node
	conn := p2p.Connect("localhost:3001")

	// Send a message to the peer node

	fromAddress1 := "0x000"
	toAddress1 := "0x001"
	amount1 := int64(100)
	data1 := "first transaction"

	// fromAddress2 := "0x001"
	// toAddress2 := "0x002"
	// amount2 := int64(50)
	// data2 := "second transaction"

	newBlockchain := blockchain.NewBlockchain() // Initialize the blockchain
	block := newBlockchain.AddBlock(fromAddress1, toAddress1, amount1, data1)

	p2p.BroadcastProcessedBlock(block)
	conn.Close()

	//newBlockchain.AddBlock(fromAddress2, toAddress2, amount2, data2)
	//// Now print all the blocks and their contents
	//println(newBlockchain.String())
}
