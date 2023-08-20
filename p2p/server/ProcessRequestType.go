package server

import (
	"fmt"
	"my-blockchain/blockchain"
	"my-blockchain/p2p"
	"my-blockchain/p2p/client"
	"net"
)

func processRequestType(conn net.Conn, request *p2p.Request) {
	switch request.RequestType {
	case p2p.NewBlockType:
		var block *blockchain.Block = &blockchain.Block{}
		p2p.ConvertMapToObject(request.Payload.(map[string]interface{}), block)
		p2p.HandleNewBlock(block)
		client.BroadcastProcessedBlock(block)

	case p2p.ProcessedBlockType:
		var block *blockchain.Block = &blockchain.Block{}
		p2p.ConvertMapToObject(request.Payload.(map[string]interface{}), block)
		fmt.Println("Received processed block2", block)
		p2p.HandleProcessedBlock(block)

	case p2p.GetSyncNodeAddressesType:
		fmt.Println("Received get sync node addresses")
		client.SendNodeAddresses(conn, &p2p.NodeAddresses)

	case p2p.GetSyncBlockchainType:
		fmt.Println("Received get sync blockchain")
		client.SendBlockchain(conn, blockchain.LocalBlockchain)

	default:
		fmt.Println("Received unknown request type")
	}

}
