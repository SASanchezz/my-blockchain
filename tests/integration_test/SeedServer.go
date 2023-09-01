package integrationTest

import (
	"my-blockchain/blockchain"
	"my-blockchain/p2p"
	clientServer "my-blockchain/p2p/client_server"
	"os"
)

func seedServer() {
	blockchain.LocalBlockchain = blockchain.NewBlockchain()
	p2p.NodeAddresses = map[string]struct{}{}

	clientServer.StartListening(os.Getenv("SEED_NODE_PORT"))
}
