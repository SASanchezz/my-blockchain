package main

import (
	"my-blockchain/blockchain"
	"my-blockchain/p2p"
)

func main() {
	// p2p.StartListening("localhost:3001")

	// Connect to a peer node
	conn := p2p.Connect("localhost:3001")
	p2p.SendGetSyncBlockchain(conn) //TODO: make it wait for responsee to sync blockchain
	conn.Close()

	fromAddress1 := "0x000"
	toAddress1 := "0x001"
	amount1 := int64(100)
	data1 := "first transaction"

	// fromAddress2 := "0x001"
	// toAddress2 := "0x002"
	// amount2 := int64(50)
	// data2 := "second transaction"

	bl := blockchain.NewBlockchain() // Initialize the blockchain

	// processedBlock := newBlockchain.AddBlock(fromAddress1, toAddress1, amount1, data1)
	// p2p.BroadcastProcessedBlock(processedBlock)

	previousBlock := bl.Chain[len(bl.Chain)-1] // the previous block is needed, so let's get it
	newBlock := blockchain.NewBlock(fromAddress1, toAddress1, amount1, data1, previousBlock.BlockHash)
	p2p.BroadcastNewBlock(newBlock)

	conn.Close()

	//newBlockchain.AddBlock(fromAddress2, toAddress2, amount2, data2)
	//// Now print all the blocks and their contents
	//println(newBlockchain.String())
}
