package main

import (
	"my-blockchain/blockchain"
	"my-blockchain/p2p"
	clientServer "my-blockchain/p2p/client_server"
	"my-blockchain/utils"
	"os"
)

func main() {
	utils.InitEnvVariables()
	blockchain.LocalBlockchain = blockchain.NewBlockchain()
	p2p.NodeAddresses = map[string]struct{}{}

	clientServer.StartListening(os.Getenv("SEED_NODE_PORT"))
}
