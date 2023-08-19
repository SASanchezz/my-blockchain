package client

import (
	"my-blockchain/blockchain"
	"my-blockchain/p2p"
)

func BroadcastProcessedBlock(block *blockchain.Block) {
	request := &p2p.Request{
		RequestType: p2p.ProcessedBlockType,
		Payload:     block,
	}
	broadcastMessge(request)
}

func BroadcastNewBlock(block *blockchain.Block) {
	request := &p2p.Request{
		RequestType: p2p.NewBlockType,
		Payload:     block,
	}
	broadcastMessge(request)
}
