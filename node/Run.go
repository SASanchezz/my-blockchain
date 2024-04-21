package node

import "my-blockchain/p2p"

func Run(nodeType string, port string) {
	switch nodeType {
	case p2p.SeedNode:
		seedNode(port)
	case p2p.MinerType:
		miner(port)
	}
}
