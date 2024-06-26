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
	clientServer.SendMyData(p2p.GetSeedNodeUrl(), myData)
	clientServer.SyncBlockchain(p2p.GetSeedNodeUrl())
	clientServer.SyncNodeAddresses(p2p.GetSeedNodeUrl())

	clientServer.StartListening(os.Getenv("SERVER_1_PORT"))
}
