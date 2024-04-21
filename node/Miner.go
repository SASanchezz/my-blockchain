package node

import (
	"my-blockchain/p2p"
	clientServer "my-blockchain/p2p/client_server"
)

func miner(port string) {
	myData := &p2p.NodeDataPayload{
		Port:         port,
		NodeDataType: p2p.MinerType,
	}
	clientServer.SendMyData(p2p.GetSeedNodeUrl(), myData)
	clientServer.SyncBlockchain(p2p.GetSeedNodeUrl())
	clientServer.SyncNodeAddresses(p2p.GetSeedNodeUrl())

	clientServer.StartListening(port)
}
