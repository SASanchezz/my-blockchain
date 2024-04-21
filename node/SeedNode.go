package node

import (
	"my-blockchain/blockchain"
	"my-blockchain/p2p"
	clientServer "my-blockchain/p2p/client_server"
)

func seedNode(port string) {
	blockchain.LocalBlockchain = blockchain.NewBlockchain()
	p2p.NodeAddresses = map[string]struct{}{}

	clientServer.StartListening(port)
}
