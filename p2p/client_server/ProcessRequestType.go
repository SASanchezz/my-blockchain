package clientServer

import (
	"fmt"
	"my-blockchain/blockchain"
	"my-blockchain/p2p"
	"net"
)

func processRequestType(conn net.Conn, request *p2p.Request) {
	switch request.RequestType {
	case p2p.NodeDataType:
		var nodeData *p2p.NodeDataPayload = &p2p.NodeDataPayload{}
		p2p.ConvertMapToObject(request.Payload.(map[string]interface{}), nodeData)
		fmt.Println("Received node data")
		HandleNodeData(conn, nodeData)

	case p2p.NewBlockType:
		var block *blockchain.Block = &blockchain.Block{}
		p2p.ConvertMapToObject(request.Payload.(map[string]interface{}), block)
		fmt.Println("Received new block")
		HandleNewBlock(block)

	case p2p.ProcessedBlockType:
		var block *blockchain.Block = &blockchain.Block{}
		p2p.ConvertMapToObject(request.Payload.(map[string]interface{}), block)
		fmt.Println("Received processed block")
		HandleProcessedBlock(block)

	case p2p.GetSyncNodeAddressesType:
		fmt.Println("Received get sync node addresses")
		SendNodeAddresses(conn, &p2p.NodeAddresses)

	case p2p.GetSyncBlockchainType:
		fmt.Println("Received get sync blockchain")
		SendBlockchain(conn, blockchain.LocalBlockchain)

	default:
		fmt.Println("Received unknown request type")
	}

}
