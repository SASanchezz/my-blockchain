package main

import (
	"my-blockchain/blockchain"
	"my-blockchain/p2p/client"
)

func main() {
	conn, err := client.Connect("localhost:3001")
	if err != nil {
		println("Error connecting to peer:", err.Error())
		return
	}
	client.SendGetSyncBlockchain(conn)
	conn.Close()

	fromAddress1 := "0x0002"
	toAddress1 := "0x003"
	amount1 := int64(100)
	data1 := "first transaction"

	// fromAddress2 := "0x001"
	// toAddress2 := "0x002"
	// amount2 := int64(50)
	// data2 := "second transaction"

	bl := *blockchain.LocalBlockchain // Initialize the blockchain

	// processedBlock := newBlockchain.AddBlock(fromAddress1, toAddress1, amount1, data1)
	// p2p.BroadcastProcessedBlock(processedBlock)

	previousBlock := bl.Chain[len(bl.Chain)-1] // the previous block is needed, so let's get it
	newBlock := blockchain.NewBlock(fromAddress1, toAddress1, amount1, data1, previousBlock.BlockHash)
	client.BroadcastNewBlock(newBlock)

	bl = *blockchain.LocalBlockchain // Initialize the blockchain
	//newBlockchain.AddBlock(fromAddress2, toAddress2, amount2, data2)
	//// Now print all the blocks and their contents
	println("last blockhain", bl.String())
}
