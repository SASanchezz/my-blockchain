package integrationTest

import (
	"my-blockchain/blockchain"
	"my-blockchain/p2p"
	clientServer "my-blockchain/p2p/client_server"
)

func Client_1() {
	clientServer.SyncBlockchain(p2p.GetSeedNodeUrl())
	clientServer.SyncNodeAddresses(p2p.GetSeedNodeUrl())

	fromAddress := "0x0002"
	toAddress := "0x003"
	amount := uint64(100)
	data := "first transaction"

	newBlock := blockchain.NewBlock(fromAddress, toAddress, amount, data)
	clientServer.SendNewBlockToRandomNode(newBlock)
}

func Client_2() {
	clientServer.SyncBlockchain(p2p.GetSeedNodeUrl())
	clientServer.SyncNodeAddresses(p2p.GetSeedNodeUrl())

	fromAddress := "0x001"
	toAddress := "0x002"
	amount := uint64(50)
	data := "second transaction"

	newBlock := blockchain.NewBlock(fromAddress, toAddress, amount, data)
	clientServer.SendNewBlockToRandomNode(newBlock)
}
