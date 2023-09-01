package integrationTest

import (
	"my-blockchain/p2p"
	clientServer "my-blockchain/p2p/client_server"
	"os"
)

func miner() {
	myData := &p2p.NodeDataPayload{
		Port:         os.Getenv("SERVER_1_PORT"),
		NodeDataType: p2p.MinerType,
	}
	clientServer.SendMyData(os.Getenv("SEED_NODE_HOST"), os.Getenv("SEED_NODE_PORT"), myData)
	clientServer.SyncBlockchain(os.Getenv("SEED_NODE_HOST"), os.Getenv("SEED_NODE_PORT"))
	clientServer.SyncNodeAddresses(os.Getenv("SEED_NODE_HOST"), os.Getenv("SEED_NODE_PORT"))

	clientServer.StartListening(os.Getenv("SERVER_1_PORT"))
}
