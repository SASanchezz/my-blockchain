package clientServer

import (
	"fmt"
	"my-blockchain/blockchain"
	"my-blockchain/p2p"
	"my-blockchain/utils"
	"net"
)

func processRequestType(conn net.Conn, request *p2p.Request) {
	switch request.RequestType {
	case p2p.NodeDataType:
		var nodeData *p2p.NodeDataPayload = &p2p.NodeDataPayload{}
		utils.ConvertMapToObject(request.Payload.(map[string]interface{}), nodeData)
		HandleNodeData(conn, nodeData)

	case p2p.NewBlockType:
		var block *blockchain.Block = &blockchain.Block{}
		utils.ConvertMapToObject(request.Payload.(map[string]interface{}), block)
		HandleNewBlock(block)

	case p2p.ProcessedBlockType:
		var block *blockchain.Block = &blockchain.Block{}
		utils.ConvertMapToObject(request.Payload.(map[string]interface{}), block)
		HandleProcessedBlock(block)

	case p2p.GetSyncNodeAddressesType:
		SendNodeAddresses(conn, &p2p.NodeAddresses)

	case p2p.GetSyncBlockchainType:
		SendBlockchain(conn, blockchain.LocalBlockchain)

	default:
		fmt.Println("Received unknown request type")
	}

}
