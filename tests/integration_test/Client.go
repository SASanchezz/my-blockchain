package integrationTest

import (
	"my-blockchain/blockchain"
	clientServer "my-blockchain/p2p/client_server"
	"os"
)

func client_1() {
	clientServer.SyncBlockchain("localhost", os.Getenv("SEED_NODE_PORT"))
	clientServer.SyncNodeAddresses("localhost", os.Getenv("SEED_NODE_PORT"))

	fromAddress := "0x0002"
	toAddress := "0x003"
	amount := uint64(100)
	data := "first transaction"

	newBlock := blockchain.NewBlock(fromAddress, toAddress, amount, data)
	clientServer.SendNewBlockToRandomNode(newBlock)
}

func client_2() {
	clientServer.SyncBlockchain("localhost", os.Getenv("SEED_NODE_PORT"))
	clientServer.SyncNodeAddresses("localhost", os.Getenv("SEED_NODE_PORT"))

	fromAddress := "0x001"
	toAddress := "0x002"
	amount := uint64(50)
	data := "second transaction"

	newBlock := blockchain.NewBlock(fromAddress, toAddress, amount, data)
	clientServer.SendNewBlockToRandomNode(newBlock)
}
