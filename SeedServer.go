package main

import (
	"my-blockchain/blockchain"
	"my-blockchain/p2p"
	"my-blockchain/p2p/server"
	"my-blockchain/utils"
	"os"
)

func main() {
	utils.InitEnvVariables()
	blockchain.LocalBlockchain = blockchain.NewBlockchain()
	p2p.NodeAddresses = map[string]struct{}{}

	server.StartListening(os.Getenv("SEED_NODE_PORT"))
}
