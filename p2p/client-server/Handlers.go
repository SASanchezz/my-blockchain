package clientServer

import (
	"my-blockchain/blockchain"
	"my-blockchain/p2p"
)

func HandleNewBlock(b *blockchain.Block) {
	b.Confirmations++
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
	blockchain.LocalBlockchain.Chain = append(blockchain.LocalBlockchain.Chain, b)
	if !blockchain.LocalBlockchain.IsValid() {
		panic("Invalid blockchain") //TODO handle this more gracefully
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
