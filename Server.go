package main

import (
	"my-blockchain/p2p/client"
	"my-blockchain/p2p/server"
	"my-blockchain/utils"
	"os"
)

func main() {
	utils.InitEnvVariables()

	client.SyncBlockchain(os.Getenv("SEED_NODE_HOST"), os.Getenv("SEED_NODE_PORT"))
	client.SyncNodeAddresses(os.Getenv("SEED_NODE_HOST"), os.Getenv("SEED_NODE_PORT"))

	server.StartListening(os.Getenv("SEED_NODE_PORT"))
}
