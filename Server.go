package main

import (
	clientServer "my-blockchain/p2p/client-server"
	"my-blockchain/utils"
	"os"
)

func main() {
	utils.InitEnvVariables()

	clientServer.SyncBlockchain(os.Getenv("SEED_NODE_HOST"), os.Getenv("SEED_NODE_PORT"))
	clientServer.SyncNodeAddresses(os.Getenv("SEED_NODE_HOST"), os.Getenv("SEED_NODE_PORT"))

	clientServer.StartListening(os.Getenv("SEED_NODE_PORT"))
}
