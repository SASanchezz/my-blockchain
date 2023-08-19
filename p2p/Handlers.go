package p2p

import "my-blockchain/blockchain"

func HandleNewBlock(b *blockchain.Block) {
	b.Mine()
	blockchain.LocalBlockchain.Chain = append(blockchain.LocalBlockchain.Chain, b)
	if !blockchain.LocalBlockchain.IsValid() {
		panic("Invalid blockchain") //TODO handle this more gracefully
	}
}

func HandleProcessedBlock(b *blockchain.Block) {
	blockchain.LocalBlockchain.Chain = append(blockchain.LocalBlockchain.Chain, b)
}

func HandleBlockchain(bl *blockchain.Blockchain) {
	blockchain.LocalBlockchain = bl
}

func HandleNodeAddresses(nAddreses *map[string]struct{}) {
	for na := range *nAddreses {
		NodeAddresses[na] = struct{}{}
	}
}
