package main

import (
	"my-blockchain/blockchain"
	clientServer "my-blockchain/p2p/client-server"
	"my-blockchain/utils"
	"os"
)

func main() {
	utils.InitEnvVariables()

	clientServer.SyncBlockchain("localhost", os.Getenv("SEED_NODE_PORT"))
	clientServer.SyncNodeAddresses("localhost", os.Getenv("SEED_NODE_PORT"))

	fromAddress1 := "0x0002"
	toAddress1 := "0x003"
	amount1 := uint64(100)
	data1 := "first transaction"

	// fromAddress2 := "0x001"
	// toAddress2 := "0x002"
	// amount2 := int64(50)
	// data2 := "second transaction"

	bl := *blockchain.LocalBlockchain // Get the blockchain

	// processedBlock := newBlockchain.AddBlock(fromAddress1, toAddress1, amount1, data1)
	// p2p.BroadcastProcessedBlock(processedBlock)

	newBlock := blockchain.NewBlock(fromAddress1, toAddress1, amount1, data1)
	clientServer.SendNewBlockToRandomNode(newBlock)

	bl = *blockchain.LocalBlockchain // Initialize the blockchain
	//newBlockchain.AddBlock(fromAddress2, toAddress2, amount2, data2)
	//// Now print all the blocks and their contents
	println("last blockhain", bl.String())
}
