package main

import (
	"my-blockchain/p2p"
	clientServer "my-blockchain/p2p/client_server"
	"my-blockchain/utils"
	"os"
)

func main() {
	utils.InitEnvVariables()

	myData := &p2p.NodeDataPayload{
		Port:         os.Getenv("SEED_NODE_PORT"),
		NodeDataType: p2p.MinerType,
	}
	clientServer.SendMyData(os.Getenv("SEED_NODE_HOST"), os.Getenv("SEED_NODE_PORT"), myData)
	clientServer.SyncBlockchain(os.Getenv("SEED_NODE_HOST"), os.Getenv("SEED_NODE_PORT"))
	clientServer.SyncNodeAddresses(os.Getenv("SEED_NODE_HOST"), os.Getenv("SEED_NODE_PORT"))

	clientServer.StartListening(os.Getenv("SEED_NODE_PORT"))
}
