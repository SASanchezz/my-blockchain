package main

import (
	"my-blockchain/blockchain"
	"my-blockchain/p2p/server"
)

func main() {
	blockchain.LocalBlockchain = blockchain.NewBlockchain()
	server.StartListening("3001")
}
