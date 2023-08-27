package clientServer

import (
	"my-blockchain/blockchain"
	"my-blockchain/p2p"
	"net"
)

func HandleNodeData(conn net.Conn, nd *p2p.NodeDataPayload) {
	if nd.NodeDataType == p2p.MinerType {
		p2p.NodeAddresses[conn.RemoteAddr().String()] = struct{}{}
	}
}

func HandleNewBlock(b *blockchain.Block) {
	b.Confirmations++
	b.PreviousBlockHash = blockchain.LocalBlockchain.Chain[len(blockchain.LocalBlockchain.Chain)-1].BlockHash
	b.Mine()
	blockchain.LocalBlockchain.Chain = append(blockchain.LocalBlockchain.Chain, b)
	if !blockchain.LocalBlockchain.IsValid() {
		println("Block is invalid")
		//Do something else
		return
	}
	if b.Confirmations >= blockchain.LocalBlockchain.ConfirmationsNeeded {
		BroadcastProcessedBlock(b)
		return
	}
	SendNewBlockToRandomNode(b)
}

func HandleProcessedBlock(b *blockchain.Block) {
	if blockchain.LocalBlockchain.Has(b) {
		return
	}

	blockchain.LocalBlockchain.Chain = append(blockchain.LocalBlockchain.Chain, b)
	if !blockchain.LocalBlockchain.IsValid() {
		print("Block is invalid")
		return
	}
}

func HandleBlockchain(bl *blockchain.Blockchain) {
	blockchain.LocalBlockchain = bl
}

func HandleNodeAddresses(nAddreses *map[string]struct{}) {
	for na := range *nAddreses {
		p2p.NodeAddresses[na] = struct{}{}
	}
}
