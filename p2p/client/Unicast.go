package client

import (
	"my-blockchain/blockchain"
	"my-blockchain/p2p"
	"net"
)

func SendNewBlock(conn net.Conn, block *blockchain.Block) {
	request := &p2p.Request{
		RequestType: p2p.NewBlockType,
		Payload:     block,
	}
	sendMessge(conn, request)
}

func SendProcessedBlock(conn net.Conn, block *blockchain.Block) {
	request := &p2p.Request{
		RequestType: p2p.ProcessedBlockType,
		Payload:     block,
	}
	sendMessge(conn, request)
}

func SendNodeAddresses(conn net.Conn, nodeAddresses *map[string]struct{}) {
	request := &p2p.Request{
		RequestType: p2p.NodeAddressesType,
		Payload:     nodeAddresses,
	}
	sendMessge(conn, request)
}

func SendBlockchain(conn net.Conn, blockchain blockchain.Blockchain) {
	request := &p2p.Request{
		RequestType: p2p.BlockchainType,
		Payload:     blockchain,
	}
	sendMessge(conn, request)
}
