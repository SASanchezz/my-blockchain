package p2p

import "my-blockchain/blockchain"

func handleNewBlock(b *blockchain.Block) {
	b.Mine()
	blockchain.LocalBlockchain.Chain = append(blockchain.LocalBlockchain.Chain, b)
	if !blockchain.LocalBlockchain.IsValid() {
		panic("Invalid blockchain") //TODO handle this more gracefully
	}
}

func handleProcessedBlock(b *blockchain.Block) {
	blockchain.LocalBlockchain.Chain = append(blockchain.LocalBlockchain.Chain, b)
}

func handleNodeAddresses(nAddreses *map[string]struct{}) {
	for na := range *nAddreses {
		nodeAddresses[na] = struct{}{}
	}
}
